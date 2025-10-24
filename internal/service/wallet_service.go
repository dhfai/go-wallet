package service

import (
	"fmt"
	"time"

	"github.com/dhfai/go-wallet/internal/domain"
	"github.com/dhfai/go-wallet/pkg/crypto"
	"github.com/dhfai/go-wallet/pkg/network"
	"github.com/google/uuid"
)

// WalletRepository defines the interface for wallet storage operations
// WalletRepository mendefinisikan interface untuk operasi penyimpanan wallet
type WalletRepository interface {
	Save(wallet *domain.Wallet) error
	FindByID(id string) (*domain.Wallet, error)
	FindByAddress(address string) (*domain.Wallet, error)
	FindAll() ([]*domain.Wallet, error)
	Update(wallet *domain.Wallet) error
	Delete(id string) error
}

// WalletService handles all wallet business logic
// WalletService menangani semua logika bisnis wallet
type WalletService struct {
	repo   WalletRepository
	crypto *crypto.BitcoinCrypto
}

// NewWalletService creates a new WalletService instance
// NewWalletService membuat instance baru WalletService
func NewWalletService(repo WalletRepository) *WalletService {
	return &WalletService{
		repo:   repo,
		crypto: crypto.NewBitcoinCrypto(),
	}
}

// CreateWallet creates a new wallet with generated keys
// CreateWallet membuat wallet baru dengan kunci yang di-generate
func (s *WalletService) CreateWallet(name string) (*domain.Wallet, error) {
	// Validate input
	if name == "" {
		return nil, fmt.Errorf("wallet name cannot be empty")
	}

	// Generate key pair
	privateKey, publicKey, err := s.crypto.GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrKeyGeneration, err)
	}

	// Generate Native SegWit address (bc1...) - Compatible with Phantom & Exchanges
	address, err := s.crypto.GenerateSegWitAddress(publicKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrKeyGeneration, err)
	}

	// Create wallet
	wallet := &domain.Wallet{
		ID:           uuid.New().String(),
		Name:         name,
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		Address:      address,
		Balance:      0.0,
		Transactions: []domain.Transaction{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save to repository
	if err := s.repo.Save(wallet); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrStorageOperation, err)
	}

	return wallet, nil
}

// GetWallet retrieves a wallet by ID
// GetWallet mengambil wallet berdasarkan ID
func (s *WalletService) GetWallet(id string) (*domain.Wallet, error) {
	wallet, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// GetWalletByAddress retrieves a wallet by address
// GetWalletByAddress mengambil wallet berdasarkan address
func (s *WalletService) GetWalletByAddress(address string) (*domain.Wallet, error) {
	wallet, err := s.repo.FindByAddress(address)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// GetAllWallets retrieves all wallets
// GetAllWallets mengambil semua wallet
func (s *WalletService) GetAllWallets() ([]*domain.Wallet, error) {
	return s.repo.FindAll()
}

// GetBalance returns the current balance of a wallet
// GetBalance mengembalikan saldo saat ini dari wallet
func (s *WalletService) GetBalance(walletID string) (float64, error) {
	wallet, err := s.repo.FindByID(walletID)
	if err != nil {
		return 0, err
	}

	return wallet.Balance, nil
}

// SendBitcoin sends Bitcoin from one wallet to another
// SendBitcoin mengirim Bitcoin dari satu wallet ke wallet lain
func (s *WalletService) SendBitcoin(fromWalletID, toAddress string, amount, fee float64, note string) (*domain.Transaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, domain.ErrInvalidAmount
	}

	// Get sender wallet
	senderWallet, err := s.repo.FindByID(fromWalletID)
	if err != nil {
		return nil, err
	}

	// Check balance
	totalAmount := amount + fee
	if senderWallet.Balance < totalAmount {
		return nil, domain.ErrInsufficientBalance
	}

	// Create transaction hash
	txHash := s.crypto.HashTransaction(
		senderWallet.Address,
		toAddress,
		amount,
		time.Now().Unix(),
	)

	// Sign transaction
	signature, err := s.crypto.SignTransaction(txHash, senderWallet.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Create transaction
	transaction := domain.Transaction{
		ID:        txHash,
		From:      senderWallet.Address,
		To:        toAddress,
		Amount:    amount,
		Fee:       fee,
		Type:      "send",
		Status:    "confirmed",
		Timestamp: time.Now(),
		Note:      fmt.Sprintf("%s | Signature: %s", note, signature[:16]+"..."),
	}

	// Add transaction to sender wallet
	senderWallet.AddTransaction(transaction)

	// Update sender wallet
	if err := s.repo.Update(senderWallet); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrStorageOperation, err)
	}

	// Try to update receiver wallet if exists in our system
	receiverWallet, err := s.repo.FindByAddress(toAddress)
	if err == nil && receiverWallet != nil {
		receiveTx := domain.Transaction{
			ID:        txHash,
			From:      senderWallet.Address,
			To:        toAddress,
			Amount:    amount,
			Fee:       0, // Receiver doesn't pay fee
			Type:      "receive",
			Status:    "confirmed",
			Timestamp: time.Now(),
			Note:      note,
		}

		receiverWallet.AddTransaction(receiveTx)
		_ = s.repo.Update(receiverWallet) // Ignore error for receiver
	}

	return &transaction, nil
}

// ReceiveBitcoin records incoming Bitcoin to a wallet
// ReceiveBitcoin mencatat Bitcoin masuk ke wallet
func (s *WalletService) ReceiveBitcoin(toWalletID, fromAddress string, amount float64, note string) (*domain.Transaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, domain.ErrInvalidAmount
	}

	// Get receiver wallet
	receiverWallet, err := s.repo.FindByID(toWalletID)
	if err != nil {
		return nil, err
	}

	// Create transaction hash
	txHash := s.crypto.HashTransaction(
		fromAddress,
		receiverWallet.Address,
		amount,
		time.Now().Unix(),
	)

	// Create transaction
	transaction := domain.Transaction{
		ID:        txHash,
		From:      fromAddress,
		To:        receiverWallet.Address,
		Amount:    amount,
		Fee:       0,
		Type:      "receive",
		Status:    "confirmed",
		Timestamp: time.Now(),
		Note:      note,
	}

	// Add transaction to wallet
	receiverWallet.AddTransaction(transaction)

	// Update wallet
	if err := s.repo.Update(receiverWallet); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrStorageOperation, err)
	}

	return &transaction, nil
}

// GetTransactionHistory retrieves transaction history for a wallet
// GetTransactionHistory mengambil riwayat transaksi untuk wallet
func (s *WalletService) GetTransactionHistory(walletID string, limit int) ([]domain.Transaction, error) {
	wallet, err := s.repo.FindByID(walletID)
	if err != nil {
		return nil, err
	}

	return wallet.GetTransactionHistory(limit), nil
}

// DeleteWallet deletes a wallet
// DeleteWallet menghapus wallet
func (s *WalletService) DeleteWallet(walletID string) error {
	return s.repo.Delete(walletID)
}

// ExportPrivateKey exports the private key of a wallet (use with caution!)
// ExportPrivateKey mengekspor private key dari wallet (gunakan dengan hati-hati!)
func (s *WalletService) ExportPrivateKey(walletID string) (string, error) {
	wallet, err := s.repo.FindByID(walletID)
	if err != nil {
		return "", err
	}

	return wallet.PrivateKey, nil
}

// ImportWallet imports a wallet from private key
// ImportWallet mengimpor wallet dari private key
func (s *WalletService) ImportWallet(name, privateKeyHex string) (*domain.Wallet, error) {
	// Validate input
	if name == "" {
		return nil, fmt.Errorf("wallet name cannot be empty")
	}

	if privateKeyHex == "" {
		return nil, domain.ErrInvalidPrivateKey
	}

	// Try to derive public key from private key
	privKey, err := s.crypto.PrivateKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidPrivateKey, err)
	}

	// Get public key
	publicKey := fmt.Sprintf("04%064x%064x", privKey.PublicKey.X, privKey.PublicKey.Y)

	// Generate address
	address, err := s.crypto.GenerateAddress(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate address: %w", err)
	}

	// Create wallet
	wallet := &domain.Wallet{
		ID:           uuid.New().String(),
		Name:         name,
		PrivateKey:   privateKeyHex,
		PublicKey:    publicKey,
		Address:      address,
		Balance:      0.0,
		Transactions: []domain.Transaction{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save to repository
	if err := s.repo.Save(wallet); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrStorageOperation, err)
	}

	return wallet, nil
}

// SyncWallet syncs wallet balance and transactions with the Bitcoin blockchain (MAINNET ONLY)
// SyncWallet sinkronisasi saldo wallet dan transaksi dengan blockchain Bitcoin (MAINNET SAJA)
func (s *WalletService) SyncWallet(walletID string) (*domain.Wallet, error) {
	// Get wallet from repository
	wallet, err := s.repo.FindByID(walletID)
	if err != nil {
		return nil, err
	}

	// Create blockchain explorer (MAINNET ONLY)
	explorer := network.NewBlockchainExplorer()

	// Get balance from blockchain
	balance, err := explorer.GetBalance(wallet.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch balance from blockchain: %w", err)
	}

	// Update wallet balance
	wallet.Balance = balance
	wallet.UpdatedAt = time.Now()

	// Save updated wallet
	if err := s.repo.Update(wallet); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrStorageOperation, err)
	}

	return wallet, nil
}
