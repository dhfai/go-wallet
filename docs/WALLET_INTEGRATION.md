# Integrasi Wallet dengan Aplikasi Lain

## ❌ Mengapa Go Wallet Tidak Bisa Connect ke Phantom/MetaMask?

### Perbedaan Fundamental:

| Aspek | Go Bitcoin Wallet | Phantom/MetaMask |
|-------|-------------------|------------------|
| **Blockchain** | Bitcoin | Solana/Ethereum |
| **Address Format** | Base58 (1..., 3..., bc1...) | Base58 (Solana) / Hex 0x... (Ethereum) |
| **Key Algorithm** | ECDSA secp256k1 | ED25519 (Solana) / ECDSA secp256k1 (Ethereum) |
| **Interface** | CLI | Browser Extension |
| **Network** | Offline | Online (RPC nodes) |
| **Protocol** | Bitcoin Script | Smart Contracts |

## ✅ Cara Menggunakan Go Bitcoin Wallet

### 1. **Export Private Key**

```bash
./bin/go-wallet export <wallet-id>
```

Output:
```
⚠️  WARNING: KEEP THIS PRIVATE KEY SECURE!
Private Key: 5Kb8kLf9zgWQnogidDA76MzPL6TsZZY36hWXMssSzNydYXYB9KF
```

### 2. **Import ke Wallet Bitcoin**

#### A. Electrum (Recommended)
1. Download: https://electrum.org
2. Buka Electrum
3. File → New/Restore
4. Pilih "Import Bitcoin addresses or private keys"
5. Paste private key
6. Click Next

#### B. BlueWallet (Mobile)
1. Download dari App Store/Play Store
2. Tap "Add Wallet"
3. Pilih "Import Wallet"
4. Paste private key
5. Tap "Import"

#### C. Bitcoin Core (Full Node)
```bash
bitcoin-cli importprivkey "your-private-key" "label"
```

#### D. Exodus (Multi-currency)
1. Download: https://exodus.com
2. Settings → Developer → Add Custom Token
3. Import private key

## 🔄 Alternatif: Buat Wallet untuk Ethereum/Solana

Jika Anda ingin wallet yang kompatibel dengan Phantom/MetaMask, saya bisa buatkan wallet baru dengan fitur:

### Ethereum Wallet Features:
- ✅ Ethereum address generation (0x...)
- ✅ ERC-20 token support
- ✅ Smart contract interaction
- ✅ Web3 JSON-RPC integration
- ✅ WalletConnect support
- ✅ MetaMask compatible

### Solana Wallet Features:
- ✅ Solana address generation
- ✅ SPL token support
- ✅ Phantom compatible
- ✅ Transaction on Solana network
- ✅ Stake account management

## 🌐 Integrasi dengan Network

### Bitcoin Network Integration

Untuk menghubungkan wallet ke Bitcoin network:

```go
// Tambahkan di pkg/network/bitcoin.go

type BitcoinNode struct {
    rpcURL string
}

func (b *BitcoinNode) BroadcastTransaction(tx string) error {
    // Connect ke Bitcoin node
    // Broadcast transaction ke network
}

func (b *BitcoinNode) GetBalance(address string) (float64, error) {
    // Query balance dari blockchain explorer
}
```

### Blockchain Explorer API

Gunakan API untuk cek balance dan broadcast:

**Blockchain.info API:**
```bash
# Get balance
curl https://blockchain.info/q/addressbalance/<address>

# Get transactions
curl https://blockchain.info/rawaddr/<address>
```

**BlockCypher API:**
```bash
# Get address info
curl https://api.blockcypher.com/v1/btc/main/addrs/<address>
```

## 🔌 WalletConnect Integration

Untuk mendukung WalletConnect (seperti Phantom/MetaMask):

### 1. Tambah Dependencies
```bash
go get github.com/ethereum/go-ethereum
go get github.com/gagliardetto/solana-go
```

### 2. Implement WalletConnect Bridge
```go
type WalletConnectBridge struct {
    sessions map[string]*Session
}

func (w *WalletConnectBridge) Connect(uri string) error {
    // Parse WalletConnect URI
    // Establish WebSocket connection
    // Handle pairing
}

func (w *WalletConnectBridge) SignTransaction(tx Transaction) error {
    // Sign transaction
    // Send back signature
}
```

## 📱 Web Interface untuk Wallet

Buat web interface agar bisa diakses browser:

```go
// cmd/server/main.go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // API endpoints
    r.GET("/wallets", listWallets)
    r.POST("/wallets", createWallet)
    r.POST("/transactions", sendTransaction)

    r.Run(":8080")
}
```

## 🔗 Browser Extension

Untuk membuat browser extension seperti MetaMask:

### Structure:
```
extension/
├── manifest.json
├── popup.html
├── popup.js
├── background.js
└── content.js
```

### manifest.json:
```json
{
  "manifest_version": 3,
  "name": "Go Wallet",
  "version": "1.0",
  "permissions": ["storage", "tabs"],
  "action": {
    "default_popup": "popup.html"
  },
  "background": {
    "service_worker": "background.js"
  }
}
```

## 🎯 Recommendation

### Untuk Pengguna Bitcoin:
1. ✅ **Gunakan Go Wallet untuk development/testing**
2. ✅ **Export private key ke Electrum/BlueWallet untuk actual use**
3. ✅ **Jangan gunakan untuk mainnet tanpa proper security**

### Untuk Pengguna Ethereum/Solana:
1. ❌ **Wallet ini tidak kompatibel**
2. ✅ **Buat wallet baru dengan support Ethereum/Solana**
3. ✅ **Atau gunakan MetaMask/Phantom directly**

## 🚀 Next Steps

### Jika ingin Go Wallet support Ethereum:

**Saya bisa buatkan:**
1. Ethereum wallet generator
2. Web3 integration
3. MetaMask compatibility
4. Smart contract interaction
5. ERC-20 token support

**Atau tambahkan:**
- REST API server
- WebSocket untuk real-time updates
- React/Vue frontend
- Browser extension

### Jika ingin Go Wallet support Solana:

**Saya bisa buatkan:**
1. Solana wallet generator
2. Phantom compatibility
3. SPL token support
4. Transaction on Solana network
5. Staking features

## ⚠️ Important Notes

1. **Go Wallet ini adalah Bitcoin wallet** - tidak bisa connect ke Phantom/MetaMask
2. **Phantom adalah Solana wallet** - hanya support Solana blockchain
3. **MetaMask adalah Ethereum wallet** - hanya support Ethereum & EVM chains
4. **Untuk menggunakan Bitcoin wallet**, gunakan Bitcoin-specific wallet apps

## 💡 Quick Fix

**Jika Anda ingin test wallet functionality:**

```bash
# 1. Create wallet
./bin/go-wallet create "TestWallet"

# 2. Note the address
# Address: 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

# 3. Get testnet Bitcoin from faucet
# Visit: https://testnet-faucet.mempool.co/
# Enter your address

# 4. Check balance (after few minutes)
./bin/go-wallet balance <wallet-id>

# 5. Send to another address
./bin/go-wallet send <wallet-id> <to-address> <amount> <fee> "test"
```

---

**Kesimpulan:**
- Go Wallet = Bitcoin only
- Phantom = Solana only
- MetaMask = Ethereum only

Mereka tidak bisa saling connect karena blockchain yang berbeda!
