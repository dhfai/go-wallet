package domain

import (
	"time"
)

type Wallet struct {
	ID           string        `json:"id"`           // Unique identifier for the wallet
	Name         string        `json:"name"`         // User-friendly name for the wallet
	PrivateKey   string        `json:"private_key"`  // Private key in WIF format
	PublicKey    string        `json:"public_key"`   // Public key in hex format
	Address      string        `json:"address"`      // Bitcoin address
	Balance      float64       `json:"balance"`      // Current balance in BTC
	Transactions []Transaction `json:"transactions"` // Transaction history
	CreatedAt    time.Time     `json:"created_at"`   // Wallet creation timestamp
	UpdatedAt    time.Time     `json:"updated_at"`   // Last update timestamp
}

type Transaction struct {
	ID        string    `json:"id"`        // Transaction ID (hash)
	From      string    `json:"from"`      // Sender address
	To        string    `json:"to"`        // Recipient address
	Amount    float64   `json:"amount"`    // Amount in BTC
	Fee       float64   `json:"fee"`       // Transaction fee in BTC
	Type      string    `json:"type"`      // Type: "send" or "receive"
	Status    string    `json:"status"`    // Status: "pending", "confirmed", "failed"
	Timestamp time.Time `json:"timestamp"` // Transaction timestamp
	Note      string    `json:"note"`      // Optional note/memo
}

type Key struct {
	PrivateKey string `json:"private_key"` // Private key in WIF format
	PublicKey  string `json:"public_key"`  // Public key in hex format
	Address    string `json:"address"`     // Bitcoin address
}

type WalletInfo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (w *Wallet) ToInfo() WalletInfo {
	return WalletInfo{
		ID:        w.ID,
		Name:      w.Name,
		Address:   w.Address,
		Balance:   w.Balance,
		CreatedAt: w.CreatedAt,
	}
}

func (w *Wallet) AddTransaction(tx Transaction) {
	w.Transactions = append(w.Transactions, tx)
	w.UpdatedAt = time.Now()

	if tx.Type == "receive" {
		w.Balance += tx.Amount
	} else if tx.Type == "send" {
		w.Balance -= (tx.Amount + tx.Fee)
	}
}

func (w *Wallet) GetTransactionHistory(limit int) []Transaction {
	if limit <= 0 || limit > len(w.Transactions) {
		return w.Transactions
	}

	start := len(w.Transactions) - limit
	return w.Transactions[start:]
}
