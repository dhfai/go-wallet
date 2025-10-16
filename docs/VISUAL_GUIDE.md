# Go Bitcoin Wallet - Visual Guide

## 🗂️ Project Structure

```
go-wallet/
│
├── 📁 cmd/                          # Application Entry Points
│   └── wallet/
│       └── main.go                  # CLI Application
│
├── 📁 internal/                     # Private Application Code
│   ├── domain/                      # 🎯 Domain Models
│   │   ├── wallet.go               # Wallet & Transaction entities
│   │   └── errors.go               # Custom error types
│   │
│   ├── service/                     # 💼 Business Logic
│   │   └── wallet_service.go       # Wallet operations & use cases
│   │
│   └── storage/                     # 💾 Data Persistence
│       └── json_repository.go      # JSON file storage
│
├── 📁 pkg/                          # Public Libraries
│   └── crypto/                      # 🔐 Cryptography
│       └── bitcoin.go              # Bitcoin crypto operations
│
├── 📁 config/                       # ⚙️ Configuration
│   └── config.go                   # App configuration
│
├── 📁 docs/                         # 📚 Documentation
│   ├── ARCHITECTURE.md             # Architecture details
│   └── PROJECT_SUMMARY.md          # Project summary
│
├── 📁 bin/                          # 🔨 Compiled Binaries
│   └── go-wallet                   # Executable (generated)
│
├── 📄 go.mod                        # Go module definition
├── 📄 go.sum                        # Dependency checksums
├── 📄 README.md                     # Main documentation
├── 📄 CONTRIBUTING.md               # Contribution guide
├── 📄 LICENSE                       # MIT License
├── 📄 Makefile                      # Build automation
└── 📄 .gitignore                    # Git ignore rules
```

## 🔄 Data Flow Diagram

### Creating a Wallet

```
┌─────────┐
│   CLI   │ ./go-wallet create "MyWallet"
└────┬────┘
     │
     ▼
┌─────────────────┐
│ WalletService   │ CreateWallet()
│                 │ - Validates input
│                 │ - Orchestrates creation
└────┬────────────┘
     │
     ├──────────────────┐
     │                  │
     ▼                  ▼
┌──────────┐      ┌──────────────┐
│  Crypto  │      │  Repository  │
│          │      │              │
│ Generate │      │  Save to     │
│ Keys     │      │  JSON file   │
│          │      │              │
│ Generate │      └──────────────┘
│ Address  │
└──────────┘
     │
     ▼
┌─────────────────┐
│  Return Wallet  │
│  to CLI         │
│  - ID           │
│  - Address      │
│  - Keys         │
└─────────────────┘
```

### Sending Bitcoin

```
┌─────────┐
│   CLI   │ ./go-wallet send [args]
└────┬────┘
     │
     ▼
┌──────────────────┐
│  WalletService   │ SendBitcoin()
│                  │
│  1. Get Wallet   │◄─────┐
│  2. Check Balance│      │
│  3. Create TX    │      │
│  4. Sign TX      │◄───┐ │
│  5. Update Wallet│    │ │
│  6. Save         │◄─┐ │ │
└──────────────────┘  │ │ │
                      │ │ │
              ┌───────┘ │ │
              │         │ │
              ▼         │ │
        ┌──────────┐    │ │
        │Repository│    │ │
        │          │    │ │
        │ FindByID │────┘ │
        │ Update   │──────┘
        └──────────┘
              ▲
              │
              ▼
        ┌──────────┐
        │  Crypto  │
        │          │
        │ Hash TX  │
        │ Sign TX  │
        └──────────┘
```

## 🎯 Layer Responsibilities

```
┌─────────────────────────────────────────────────────────┐
│                     CLI Layer                           │
│  • Command parsing                                      │
│  • User interaction                                     │
│  • Output formatting                                    │
└───────────────────┬─────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────────┐
│                   Service Layer                         │
│  • Business logic                                       │
│  • Use case orchestration                               │
│  • Validation                                           │
│  • Error handling                                       │
└──────────┬─────────────────────┬────────────────────────┘
           │                     │
           ▼                     ▼
    ┌──────────┐         ┌──────────────┐
    │  Crypto  │         │  Repository  │
    │  Layer   │         │    Layer     │
    │          │         │              │
    │ • Keys   │         │ • CRUD       │
    │ • Sign   │         │ • Storage    │
    │ • Verify │         │ • Queries    │
    └──────────┘         └──────┬───────┘
                                │
                                ▼
                         ┌──────────────┐
                         │  JSON File   │
                         │   Storage    │
                         └──────────────┘
                                ▲
                                │
                                ▼
                         ┌──────────────┐
                         │ Domain Layer │
                         │              │
                         │ • Wallet     │
                         │ • Transaction│
                         │ • Errors     │
                         └──────────────┘
```

## 🔐 Cryptography Flow

### Key Generation
```
Random Entropy
     │
     ▼
ECDSA Key Generation
     │
     ├──► Private Key (hex)
     │
     └──► Public Key (hex)
          │
          ▼
     SHA-256 Hash
          │
          ▼
    RIPEMD-160 Hash
          │
          ▼
   Add Version Byte (0x00)
          │
          ▼
   Double SHA-256 (checksum)
          │
          ▼
   Append Checksum
          │
          ▼
   Base58 Encode
          │
          ▼
   Bitcoin Address
```

### Transaction Signing
```
Transaction Data
(from, to, amount, timestamp)
          │
          ▼
     SHA-256 Hash
          │
          ▼
    Transaction Hash
          │
          ├──► Sign with Private Key
          │         │
          │         ▼
          │    ECDSA Signature
          │         │
          │         ▼
          │    Store in Transaction
          │
          └──► Can be verified with Public Key
```

## 💾 Storage Structure

```
~/.go-wallet/wallets.json
│
└── [
      {
        "id": "uuid-1",
        "name": "My Wallet",
        "private_key": "hex...",
        "public_key": "hex...",
        "address": "1A1z...",
        "balance": 0.5,
        "transactions": [
          {
            "id": "tx-hash",
            "from": "address1",
            "to": "address2",
            "amount": 0.1,
            "fee": 0.0001,
            "type": "send",
            "status": "confirmed",
            "timestamp": "2025-10-17T...",
            "note": "Payment"
          }
        ],
        "created_at": "2025-10-17T...",
        "updated_at": "2025-10-17T..."
      },
      { ... }
    ]
```

## 📊 Component Interaction

```
┌──────────────────────────────────────────────────────────┐
│                    User                                  │
└───────────────────┬──────────────────────────────────────┘
                    │
                    │ Commands
                    ▼
        ┌───────────────────────┐
        │   CLI Interface       │
        │   (cmd/wallet)        │
        └───────┬───────────────┘
                │
                │ Method Calls
                ▼
        ┌───────────────────────┐
        │  Wallet Service       │◄──── Config
        │  (internal/service)   │
        └────┬──────────┬───────┘
             │          │
      ┌──────┘          └──────┐
      │                        │
      ▼                        ▼
┌─────────────┐         ┌──────────────┐
│   Crypto    │         │  Repository  │
│ (pkg/crypto)│         │  (storage)   │
└─────────────┘         └──────┬───────┘
                               │
                               │ Read/Write
                               ▼
                        ┌──────────────┐
                        │  JSON File   │
                        └──────────────┘
      ▲
      │ Uses
      │
┌─────┴──────────┐
│ Domain Models  │
│ (Wallet, TX)   │
└────────────────┘
```

## 🎮 Command Flow Examples

### 1. Create Wallet
```
$ ./go-wallet create "MyWallet"
  ↓
CLI parses command
  ↓
Calls service.CreateWallet("MyWallet")
  ↓
Service generates keys (crypto layer)
  ↓
Service creates Wallet entity
  ↓
Service saves to repository
  ↓
Repository writes JSON file
  ↓
CLI displays wallet info
```

### 2. Send Bitcoin
```
$ ./go-wallet send wallet-id address 0.5 0.0001 "note"
  ↓
CLI parses arguments
  ↓
Calls service.SendBitcoin(...)
  ↓
Service loads sender wallet (repository)
  ↓
Service checks balance
  ↓
Service creates transaction hash (crypto)
  ↓
Service signs transaction (crypto)
  ↓
Service updates wallet
  ↓
Service saves wallet (repository)
  ↓
CLI displays transaction details
```

## 🔄 State Management

```
┌────────────────┐
│  Application   │
│    Start       │
└───────┬────────┘
        │
        ▼
┌────────────────┐
│  Load Config   │
└───────┬────────┘
        │
        ▼
┌────────────────────┐
│  Init Repository   │
│  (Load from file)  │
└───────┬────────────┘
        │
        ▼
┌────────────────────┐
│  Init Services     │
└───────┬────────────┘
        │
        ▼
┌────────────────────┐
│  Parse Command     │
└───────┬────────────┘
        │
        ▼
┌────────────────────┐
│  Execute Command   │
└───────┬────────────┘
        │
        ▼
┌────────────────────┐
│  Save State        │
│  (if modified)     │
└───────┬────────────┘
        │
        ▼
┌────────────────┐
│  Display       │
│  Result        │
└────────────────┘
```

## 📝 Code Example Flow

```go
// 1. User runs command
$ ./go-wallet create "MyWallet"

// 2. Main function in cmd/wallet/main.go
func main() {
    cfg := config.NewConfig()
    repo := storage.NewJSONWalletRepository(cfg.StoragePath)
    service := service.NewWalletService(repo)

    // Parse command
    if os.Args[1] == "create" {
        handleCreate(service)
    }
}

// 3. Handler calls service
func handleCreate(service *service.WalletService) {
    wallet := service.CreateWallet(name)
    // Display result
}

// 4. Service orchestrates
func (s *WalletService) CreateWallet(name string) (*Wallet, error) {
    // Generate keys
    privKey, pubKey := s.crypto.GenerateKeyPair()

    // Generate address
    address := s.crypto.GenerateAddress(pubKey)

    // Create entity
    wallet := &Wallet{...}

    // Save
    s.repo.Save(wallet)

    return wallet
}

// 5. Repository persists
func (r *JSONRepo) Save(wallet *Wallet) error {
    // Write to JSON file
    // ...
}
```

## 🎯 Key Design Decisions

1. **Clean Architecture** → Maintainability
2. **Repository Pattern** → Flexibility
3. **Dependency Injection** → Testability
4. **Interface Design** → Extensibility
5. **JSON Storage** → Simplicity
6. **CLI Interface** → Ease of Use
7. **Crypto Package** → Security
8. **Domain Models** → Business Logic Clarity

---

**This visual guide helps understand:**
- How components interact
- Data flow through the system
- Layer responsibilities
- Storage structure
- Command execution flow
