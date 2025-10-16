# Go Bitcoin Wallet 🪙

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)

Aplikasi Bitcoin Wallet profesional yang dibangun menggunakan Go dengan arsitektur modular, clean code, dan dokumentasi lengkap.

## 📋 Daftar Isi

- [Fitur](#-fitur)
- [Arsitektur](#-arsitektur)
- [Instalasi](#-instalasi)
- [Penggunaan](#-penggunaan)
- [Struktur Proyek](#-struktur-proyek)
- [API Documentation](#-api-documentation)
- [Keamanan](#-keamanan)
- [Kontribusi](#-kontribusi)

## ✨ Fitur

- ✅ **Pembuatan Wallet Baru** - Generate wallet dengan pasangan kunci privat/publik otomatis
- ✅ **Manajemen Multiple Wallet** - Kelola beberapa wallet dalam satu aplikasi
- ✅ **Kirim & Terima Bitcoin** - Transaksi Bitcoin dengan signature kriptografi
- ✅ **Riwayat Transaksi** - Catat dan tampilkan semua transaksi
- ✅ **Export/Import Wallet** - Backup dan restore wallet menggunakan private key
- ✅ **Balance Tracking** - Monitor saldo Bitcoin real-time
- ✅ **Secure Storage** - Penyimpanan terenkripsi dengan JSON
- ✅ **CLI Interface** - Command-line interface yang user-friendly

## 🏗️ Arsitektur

Aplikasi ini menggunakan **Clean Architecture** dengan pemisahan concern yang jelas:

```
┌─────────────────────────────────────────┐
│         CLI Interface (cmd/)            │
│  - Command Handler                      │
│  - User Interaction                     │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│      Service Layer (internal/service)   │
│  - Business Logic                       │
│  - Wallet Operations                    │
│  - Transaction Management               │
└──────────────┬──────────────────────────┘
               │
    ┌──────────┴──────────┐
    │                     │
┌───▼────────────┐  ┌────▼────────────────┐
│ Storage Layer  │  │  Crypto Utilities   │
│ (storage/)     │  │  (pkg/crypto/)      │
│ - JSON Repo    │  │  - Key Generation   │
│ - CRUD Ops     │  │  - Signing          │
└────────────────┘  └─────────────────────┘
         │
┌────────▼────────────────────────────────┐
│       Domain Models (domain/)           │
│  - Wallet                               │
│  - Transaction                          │
│  - Error Definitions                    │
└─────────────────────────────────────────┘
```

### Layer Breakdown

1. **Domain Layer** (`internal/domain/`)
   - Entitas bisnis (Wallet, Transaction)
   - Error definitions
   - Business rules

2. **Service Layer** (`internal/service/`)
   - Business logic
   - Orchestration
   - Validation

3. **Storage Layer** (`internal/storage/`)
   - Data persistence
   - Repository pattern
   - JSON file storage

4. **Crypto Utilities** (`pkg/crypto/`)
   - Key pair generation (ECDSA)
   - Bitcoin address generation
   - Transaction signing & verification

5. **CLI Interface** (`cmd/wallet/`)
   - Command parsing
   - User interaction
   - Output formatting

6. **Configuration** (`config/`)
   - Application settings
   - Storage paths

## 📦 Instalasi

### Prerequisites

- Go 1.21 atau lebih tinggi
- Git

### Langkah Instalasi

1. **Clone repository**
```bash
git clone https://github.com/dhfai/go-wallet.git
cd go-wallet
```

2. **Install dependencies**
```bash
go mod download
```

3. **Build aplikasi**
```bash
go build -o go-wallet cmd/wallet/main.go
```

4. **Jalankan aplikasi**
```bash
./go-wallet help
```

## 🚀 Penggunaan

### Membuat Wallet Baru

```bash
./go-wallet create MyFirstWallet
```

Output:
```
✓ Wallet created successfully!

=== Wallet Details ===
ID:         550e8400-e29b-41d4-a716-446655440000
Name:       MyFirstWallet
Address:    1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
Balance:    0.00000000 BTC
Created:    2025-10-17T10:30:00Z

⚠️  IMPORTANT: Please backup your private key securely!
Use 'go-wallet export <wallet-id>' to export your private key
```

### Melihat Semua Wallet

```bash
./go-wallet list
```

Output:
```
=== Wallets (2 total) ===

ID                                    Name             Address                              Balance (BTC)   Transactions
----                                  ----             ----                                 ----            ----
550e8400-e29b-41d4-a716-446655440000  MyFirstWallet    1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa  0.50000000      5
650e8400-e29b-41d4-a716-446655440001  SavingsWallet    1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2  1.25000000      3
```

### Cek Balance

```bash
./go-wallet balance 550e8400-e29b-41d4-a716-446655440000
```

### Kirim Bitcoin

```bash
./go-wallet send 550e8400-e29b-41d4-a716-446655440000 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2 0.5 0.0001 "Payment for services"
```

### Terima Bitcoin

```bash
./go-wallet receive 550e8400-e29b-41d4-a716-446655440000 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa 0.5 "Received payment"
```

### Melihat Riwayat Transaksi

```bash
# Semua transaksi
./go-wallet history 550e8400-e29b-41d4-a716-446655440000

# 10 transaksi terakhir
./go-wallet history 550e8400-e29b-41d4-a716-446655440000 10
```

### Export Private Key

```bash
./go-wallet export 550e8400-e29b-41d4-a716-446655440000
```

⚠️ **PERINGATAN**: Jangan pernah share private key Anda dengan siapapun!

### Import Wallet

```bash
./go-wallet import RestoredWallet <private-key-hex>
```

### Delete Wallet

```bash
./go-wallet delete 550e8400-e29b-41d4-a716-446655440000
```

## 📁 Struktur Proyek

```
go-wallet/
├── cmd/
│   └── wallet/
│       └── main.go                 # CLI entry point
├── internal/
│   ├── domain/
│   │   ├── wallet.go              # Domain models
│   │   └── errors.go              # Error definitions
│   ├── service/
│   │   └── wallet_service.go      # Business logic
│   └── storage/
│       └── json_repository.go     # Data persistence
├── pkg/
│   └── crypto/
│       └── bitcoin.go             # Crypto utilities
├── config/
│   └── config.go                  # Configuration
├── go.mod                         # Go module definition
└── README.md                      # Documentation
```

### Penjelasan Direktori

- **`cmd/`** - Aplikasi executable dan CLI handlers
- **`internal/`** - Private application code (tidak bisa di-import oleh package lain)
  - `domain/` - Core business entities dan rules
  - `service/` - Business logic dan use cases
  - `storage/` - Data access layer
- **`pkg/`** - Public libraries (bisa di-import oleh package lain)
  - `crypto/` - Cryptographic utilities
- **`config/`** - Configuration management

## 📚 API Documentation

### Domain Models

#### Wallet

```go
type Wallet struct {
    ID           string        // Unique identifier
    Name         string        // User-friendly name
    PrivateKey   string        // Private key (WIF format)
    PublicKey    string        // Public key (hex)
    Address      string        // Bitcoin address
    Balance      float64       // Balance in BTC
    Transactions []Transaction // Transaction history
    CreatedAt    time.Time     // Creation timestamp
    UpdatedAt    time.Time     // Last update timestamp
}
```

#### Transaction

```go
type Transaction struct {
    ID        string    // Transaction ID (hash)
    From      string    // Sender address
    To        string    // Recipient address
    Amount    float64   // Amount in BTC
    Fee       float64   // Transaction fee
    Type      string    // "send" or "receive"
    Status    string    // "pending", "confirmed", "failed"
    Timestamp time.Time // Transaction time
    Note      string    // Optional memo
}
```

### Service Methods

#### WalletService

```go
// Create new wallet
CreateWallet(name string) (*Wallet, error)

// Get wallet by ID
GetWallet(id string) (*Wallet, error)

// Get all wallets
GetAllWallets() ([]*Wallet, error)

// Send Bitcoin
SendBitcoin(fromWalletID, toAddress string, amount, fee float64, note string) (*Transaction, error)

// Receive Bitcoin
ReceiveBitcoin(toWalletID, fromAddress string, amount float64, note string) (*Transaction, error)

// Get transaction history
GetTransactionHistory(walletID string, limit int) ([]Transaction, error)

// Export private key
ExportPrivateKey(walletID string) (string, error)

// Import wallet
ImportWallet(name, privateKeyHex string) (*Wallet, error)

// Delete wallet
DeleteWallet(walletID string) error
```

## 🔒 Keamanan

### Best Practices

1. **Private Key Storage**
   - Private keys disimpan dalam file JSON dengan permission 0600
   - Default location: `~/.go-wallet/wallets.json`
   - Jangan pernah commit file ini ke version control

2. **Backup**
   - Selalu backup private key Anda
   - Simpan di lokasi yang aman dan terenkripsi
   - Gunakan multiple backup locations

3. **Network**
   - Aplikasi ini adalah wallet lokal (tidak terhubung ke network Bitcoin)
   - Untuk production, perlu integrasi dengan Bitcoin node

4. **Cryptography**
   - Menggunakan ECDSA untuk key generation
   - SHA-256 untuk hashing
   - RIPEMD-160 untuk address generation

### Security Checklist

- [ ] Backup private keys secara teratur
- [ ] Gunakan password manager untuk menyimpan private keys
- [ ] Jangan share private keys dengan siapapun
- [ ] Verifikasi address penerima sebelum mengirim
- [ ] Set file permissions yang benar pada storage file
- [ ] Update dependencies secara berkala

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/service/
```

## 🔧 Development

### Setup Development Environment

```bash
# Install development tools
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Format code
go fmt ./...

# Lint code
golangci-lint run
```

### Adding New Features

1. Update domain models jika diperlukan
2. Implement business logic di service layer
3. Update CLI handlers di cmd/wallet/main.go
4. Add tests
5. Update documentation

## 🤝 Kontribusi

Kontribusi sangat diterima! Silakan:

1. Fork repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.

## 👨‍💻 Author

**Dhfai**
- GitHub: [@dhfai](https://github.com/dhfai)

## 🙏 Acknowledgments

- Bitcoin Protocol Documentation
- Go Crypto Libraries
- Clean Architecture Principles

## 📞 Support

Jika Anda memiliki pertanyaan atau issue:

1. Check [existing issues](https://github.com/dhfai/go-wallet/issues)
2. Open new issue dengan detail yang jelas
3. Provide logs dan error messages

---

**⚠️ DISCLAIMER**: Aplikasi ini adalah untuk tujuan edukasi dan development. Untuk production use dengan Bitcoin real, diperlukan additional security measures dan integrasi dengan Bitcoin network yang proper.
