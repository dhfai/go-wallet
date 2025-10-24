# 🔴 MAINNET CONFIGURATION - Production Bitcoin Wallet

## ⚠️ CRITICAL SAFETY INFORMATION

**THIS WALLET IS NOW CONFIGURED FOR REAL BITCOIN MAINNET**

- **NO TESTING MODE** - All transactions use real Bitcoin
- **REAL MONEY** - Any Bitcoin sent here is real value
- **IRREVERSIBLE** - Bitcoin transactions cannot be undone
- **SECURE YOUR KEYS** - Never share your private keys or WIF exports

---

## ✅ Mainnet Features Enabled

### 1. **Native SegWit Addresses (bc1...)**
Your wallet generates modern Native SegWit addresses that start with `bc1...`

**Benefits:**
- ✅ Lower transaction fees
- ✅ Compatible with all modern exchanges (Binance, Coinbase, Kraken, etc.)
- ✅ Compatible with hardware wallets (Ledger, Trezor)
- ✅ Compatible with software wallets (Phantom, Electrum, BlueWallet)

### 2. **Real Blockchain Integration**
Connected to Blockstream.info API for mainnet data:
- ✅ Real-time balance checking
- ✅ Transaction history viewing
- ✅ UTXO management
- ✅ Address verification

**Endpoint:** `https://blockstream.info/api` (MAINNET ONLY)

### 3. **WIF Export for Wallet Import**
Export private keys in WIF (Wallet Import Format) to import into other wallets:
- ✅ Phantom wallet (Bitcoin network)
- ✅ Electrum
- ✅ BlueWallet
- ✅ Any BIP38/WIF compatible wallet

---

## 🚀 Usage Guide

### Creating a Wallet (Mainnet)

```bash
./bin/go-wallet create "My Bitcoin Wallet"
```

**Output:**
```
✅ Wallet created successfully!

Wallet Details:
  ID:         c0788ad0-165a-47ea-a9bc-2917868145ca
  Name:       My Bitcoin Wallet
  Address:    bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
  Private Key: a1b2c3d4e5f6... (KEEP THIS SECRET!)
  Balance:    0.00000000 BTC

⚠️  IMPORTANT SECURITY NOTES:
1. NEVER share your Private Key with anyone
2. NEVER store it in plain text on cloud services
3. Consider using hardware wallet for large amounts
4. This is a MAINNET address - it can receive REAL Bitcoin
```

### Checking Real Balance

```bash
./bin/go-wallet sync <wallet-id>
```

**What it does:**
1. Connects to Bitcoin mainnet via Blockstream API
2. Queries your address balance
3. Updates local wallet data
4. Shows confirmed + unconfirmed balance

**Output:**
```
🔄 Syncing wallet with Bitcoin blockchain (MAINNET)...
⚠️  WARNING: This is REAL Bitcoin mainnet - not test network

✅ Wallet synced successfully!

📊 Wallet Details:
   Name:     My Bitcoin Wallet
   Address:  bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
   Balance:  0.00000000 BTC
   Updated:  2025-10-24 14:40:54

🔍 View on blockchain: https://blockstream.info/address/bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
```

### Exporting to Phantom Wallet

```bash
./bin/go-wallet export-wif <wallet-id>
```

**Steps to import into Phantom:**
1. Open Phantom wallet
2. Click Settings → Import Private Key
3. Select "Bitcoin" network
4. Paste the WIF key shown
5. Your wallet is now accessible in Phantom!

---

## 🔐 Security Best Practices

### 1. **Private Key Storage**
```json
// Wallets are stored in: data/wallets.json
{
  "id": "...",
  "private_key": "EXTREMELY_SENSITIVE_DATA",
  "address": "bc1q..."
}
```

**Recommendations:**
- ❌ DO NOT commit `data/wallets.json` to Git
- ❌ DO NOT store in Dropbox/Google Drive
- ✅ DO encrypt your hard drive
- ✅ DO use hardware wallet for large amounts
- ✅ DO backup offline (encrypted USB/paper wallet)

### 2. **For Production Use**

If you're building a production service:

```go
// Use environment variables for sensitive data
privateKey := os.Getenv("BTC_PRIVATE_KEY")

// Use HSM (Hardware Security Module) for key storage
// Consider services like AWS KMS, Azure Key Vault

// Implement rate limiting on API calls
// Implement multi-signature for large amounts
// Use cold storage for majority of funds
```

### 3. **Receiving Bitcoin from Exchanges**

**Step-by-step guide:**

1. **Create wallet**
   ```bash
   ./bin/go-wallet create "Exchange Withdrawal"
   ```

2. **Copy your address** (starts with `bc1...`)
   ```
   Address: bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
   ```

3. **On exchange (e.g., Binance, Coinbase):**
   - Go to Withdraw → Bitcoin
   - Select Network: **Bitcoin (BTC)** [NOT Bitcoin Lightning Network]
   - Paste your `bc1...` address
   - Enter amount
   - Confirm withdrawal

4. **Wait for confirmations:**
   - 1 confirmation: ~10 minutes (visible in mempool)
   - 3 confirmations: ~30 minutes (recommended)
   - 6 confirmations: ~60 minutes (fully secure)

5. **Check balance:**
   ```bash
   ./bin/go-wallet sync <wallet-id>
   ```

---

## 🌐 Blockchain Explorers

View your transactions on these explorers:

### Blockstream (Recommended)
```
https://blockstream.info/address/YOUR_ADDRESS
```

### Blockchain.com
```
https://www.blockchain.com/explorer/addresses/btc/YOUR_ADDRESS
```

### Mempool.space
```
https://mempool.space/address/YOUR_ADDRESS
```

---

## 📊 Transaction Status

### Understanding Confirmations

| Confirmations | Status | Time | Safety |
|--------------|--------|------|--------|
| 0 | Pending (Mempool) | 0 min | ⚠️ Not safe |
| 1 | First confirmation | ~10 min | ⚠️ Low risk purchases |
| 3 | Standard safety | ~30 min | ✅ Most purchases |
| 6 | Full security | ~60 min | ✅ Large amounts |

### Transaction Fees

**Current implementation uses default fee:**
- Fee: 0.00001 BTC (1,000 satoshis)
- For urgent transactions, increase fee
- Check current fees: https://mempool.space/

---

## 🛠️ Available Commands

| Command | Description | Example |
|---------|-------------|---------|
| `create` | Create new mainnet wallet | `./bin/go-wallet create "My Wallet"` |
| `list` | List all wallets | `./bin/go-wallet list` |
| `sync` | Sync balance from blockchain | `./bin/go-wallet sync <id>` |
| `balance` | Show wallet balance | `./bin/go-wallet balance <id>` |
| `receive` | Show receive address | `./bin/go-wallet receive <id>` |
| `export` | Export wallet details | `./bin/go-wallet export <id>` |
| `export-wif` | Export WIF for import | `./bin/go-wallet export-wif <id>` |
| `history` | Show transaction history | `./bin/go-wallet history <id>` |
| `delete` | Delete wallet | `./bin/go-wallet delete <id>` |

---

## ⚠️ Important Warnings

### 1. **This is NOT a Testnet Wallet**
- ❌ No testnet mode available
- ❌ Cannot use testnet faucets
- ✅ All addresses are mainnet (bc1...)
- ✅ All transactions use real Bitcoin

### 2. **Transaction Broadcasting Not Yet Implemented**
Currently, you can:
- ✅ Generate addresses
- ✅ Check balances
- ✅ Export private keys
- ❌ Send Bitcoin (coming soon)

To send Bitcoin for now:
1. Export your private key (WIF)
2. Import into Phantom, Electrum, or BlueWallet
3. Send from there

### 3. **API Rate Limits**
Blockstream API has rate limits:
- Don't spam sync command
- Wait 1-2 seconds between requests
- For production, consider running your own Bitcoin node

---

## 🔄 Integration with Phantom Wallet

### Why Phantom Works Now

**Phantom recently added Bitcoin support** (previously Solana-only)

**Compatible features:**
- ✅ Native SegWit addresses (bc1...)
- ✅ WIF private key import
- ✅ Mainnet transaction signing
- ✅ Hardware wallet integration

**How to import:**
1. Export WIF from this wallet
2. Open Phantom → Settings
3. Import Private Key → Select Bitcoin network
4. Paste WIF key
5. Done! Your Bitcoin is accessible in Phantom

---

## 📞 Support & Resources

### Documentation
- [Main README](../README.md)
- [Architecture Guide](./ARCHITECTURE.md)
- [Real Transaction Guide](./REAL_TRANSACTION_GUIDE.md)
- [Phantom Integration](./PHANTOM_INTEGRATION.md)

### Bitcoin Resources
- Bitcoin Whitepaper: https://bitcoin.org/bitcoin.pdf
- Bitcoin Core: https://bitcoin.org/en/download
- Learn Bitcoin: https://learn.saylor.org/course/view.php?id=468

### API Documentation
- Blockstream API: https://github.com/Blockstream/esplora/blob/master/API.md
- Bitcoin RPC: https://developer.bitcoin.org/reference/rpc/

---

## 🎯 Roadmap

**Current Status:** ✅ Ready for receiving Bitcoin

**Coming Soon:**
- [ ] Transaction broadcasting (send Bitcoin)
- [ ] Fee estimation API integration
- [ ] Multi-signature support
- [ ] Hardware wallet integration
- [ ] Transaction history from blockchain
- [ ] QR code generation for addresses

---

## 📝 License & Disclaimer

**MIT License** - Use at your own risk

**DISCLAIMER:**
This software is provided "as is" without warranty of any kind. The authors are not responsible for any loss of funds. Always test with small amounts first. Bitcoin transactions are irreversible.

**For production use:**
- Conduct security audit
- Use hardware wallets
- Implement multi-signature
- Use cold storage for majority of funds
- Never store large amounts in hot wallets

---

**Last Updated:** 2025-10-24
**Wallet Version:** 1.0.0-mainnet
**Bitcoin Network:** MAINNET ONLY
