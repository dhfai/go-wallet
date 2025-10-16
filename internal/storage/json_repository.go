package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/dhfai/go-wallet/internal/domain"
)

// JSONWalletRepository implements WalletRepository using JSON file storage
// JSONWalletRepository mengimplementasikan WalletRepository menggunakan penyimpanan file JSON
type JSONWalletRepository struct {
	filePath string
	mu       sync.RWMutex
	wallets  map[string]*domain.Wallet
}

// NewJSONWalletRepository creates a new JSONWalletRepository
// NewJSONWalletRepository membuat instance baru JSONWalletRepository
func NewJSONWalletRepository(filePath string) (*JSONWalletRepository, error) {
	repo := &JSONWalletRepository{
		filePath: filePath,
		wallets:  make(map[string]*domain.Wallet),
	}

	// Create directory if not exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Load existing wallets
	if err := repo.load(); err != nil {
		// If file doesn't exist, that's okay
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return repo, nil
}

// Save saves a wallet to storage
// Save menyimpan wallet ke storage
func (r *JSONWalletRepository) Save(wallet *domain.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if wallet already exists
	if _, exists := r.wallets[wallet.ID]; exists {
		return domain.ErrWalletExists
	}

	r.wallets[wallet.ID] = wallet
	return r.persist()
}

// FindByID finds a wallet by ID
// FindByID mencari wallet berdasarkan ID
func (r *JSONWalletRepository) FindByID(id string) (*domain.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	wallet, exists := r.wallets[id]
	if !exists {
		return nil, domain.ErrWalletNotFound
	}

	return wallet, nil
}

// FindByAddress finds a wallet by address
// FindByAddress mencari wallet berdasarkan address
func (r *JSONWalletRepository) FindByAddress(address string) (*domain.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, wallet := range r.wallets {
		if wallet.Address == address {
			return wallet, nil
		}
	}

	return nil, domain.ErrWalletNotFound
}

// FindAll retrieves all wallets
// FindAll mengambil semua wallet
func (r *JSONWalletRepository) FindAll() ([]*domain.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	wallets := make([]*domain.Wallet, 0, len(r.wallets))
	for _, wallet := range r.wallets {
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

// Update updates an existing wallet
// Update memperbarui wallet yang sudah ada
func (r *JSONWalletRepository) Update(wallet *domain.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.wallets[wallet.ID]; !exists {
		return domain.ErrWalletNotFound
	}

	r.wallets[wallet.ID] = wallet
	return r.persist()
}

// Delete deletes a wallet by ID
// Delete menghapus wallet berdasarkan ID
func (r *JSONWalletRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.wallets[id]; !exists {
		return domain.ErrWalletNotFound
	}

	delete(r.wallets, id)
	return r.persist()
}

// persist saves wallets to JSON file
// persist menyimpan wallets ke file JSON
func (r *JSONWalletRepository) persist() error {
	// Convert map to slice for JSON
	walletSlice := make([]*domain.Wallet, 0, len(r.wallets))
	for _, wallet := range r.wallets {
		walletSlice = append(walletSlice, wallet)
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(walletSlice, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal wallets: %w", err)
	}

	// Write to file
	if err := os.WriteFile(r.filePath, data, 0600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// load loads wallets from JSON file
// load memuat wallets dari file JSON
func (r *JSONWalletRepository) load() error {
	// Read file
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	// Unmarshal JSON
	var walletSlice []*domain.Wallet
	if err := json.Unmarshal(data, &walletSlice); err != nil {
		return fmt.Errorf("failed to unmarshal wallets: %w", err)
	}

	// Convert slice to map
	r.wallets = make(map[string]*domain.Wallet)
	for _, wallet := range walletSlice {
		r.wallets[wallet.ID] = wallet
	}

	return nil
}

// GetStoragePath returns the storage file path
// GetStoragePath mengembalikan path file storage
func (r *JSONWalletRepository) GetStoragePath() string {
	return r.filePath
}
