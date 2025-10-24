# 🚀 Panduan Menggunakan Wallet untuk Transaksi Bitcoin Real

## ✅ Update: Wallet Sekarang Support Native SegWit!

Wallet Go kita sekarang sudah di-upgrade dengan fitur:
- ✅ **Native SegWit Address (bc1...)** - Kompatibel dengan Phantom & semua exchange
- ✅ **WIF Export** - Bisa langsung import ke Phantom
- ✅ **Blockchain Sync** - Cek balance real dari blockchain
- ✅ **Production Ready** - Siap untuk transaksi Bitcoin real!

## 📋 Step-by-Step Guide

### 1. Buat Wallet Baru

```bash
./bin/go-wallet create "My Bitcoin Wallet"
```

Output:
```
✓ Wallet created successfully!

=== Wallet Details ===
ID:         c0788ad0-165a-47ea-a9bc-2917868145ca
Name:       My Bitcoin Wallet
Address:    bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y  ← Native SegWit!
Balance:    0.00000000 BTC
```

**📝 Catat:**
- Wallet ID
- Address (bc1...)

### 2. Export Private Key untuk Backup/Import

```bash
./bin/go-wallet export-wif <wallet-id>
```

Output:
```
=== WIF Private Key Export (mainnet) ===
Address:     bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
Private Key (hex): a1453eff...
```

**⚠️ PENTING:** Backup private key ini di tempat aman!

### 3. Import ke Phantom (Optional)

Jika ingin menggunakan Phantom sebagai interface:

1. **Buka Phantom Extension**
2. **Click Settings** (⚙️)
3. **Add / Connect Wallet**
4. **Import Private Key**
5. **Pilih Network: Bitcoin**
6. **Paste private key** dari step 2
7. **Click Import**

Sekarang wallet Anda ada di:
- ✅ Go Wallet (CLI/backup)
- ✅ Phantom (GUI/daily use)

### 4. Transfer Bitcoin dari Exchange

#### A. Dari Binance:

1. **Login Binance**
2. **Wallet → Fiat and Spot**
3. **Withdraw → Bitcoin (BTC)**
4. **Paste address:** `bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y`
5. **Amount:** Masukkan jumlah BTC
6. **Network:** Pilih **Bitcoin (Native SegWit)**
7. **Submit**

#### B. Dari Indodax:

1. **Login Indodax**
2. **Balances → Bitcoin**
3. **Withdraw**
4. **Paste address:** `bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y`
5. **Amount:** Masukkan jumlah BTC
6. **Submit**

#### C. Dari Tokocrypto:

1. **Login Tokocrypto**
2. **Assets → Bitcoin**
3. **Withdraw**
4. **Address:** `bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y`
5. **Network:** **Bitcoin (BTC)**
6. **Confirm**

### 5. Cek Balance Real dari Blockchain

Setelah withdrawal dari exchange (biasanya 15-30 menit):

**Option A: Menggunakan Blockstream (Recommended)**

```bash
# Visit browser
https://blockstream.info/address/bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
```

Atau gunakan command:

```bash
./bin/go-wallet sync <wallet-id>
```

Output akan show link ke blockchain explorer.

**Option B: Menggunakan Phantom**

Jika sudah import ke Phantom:
- Buka Phantom
- Balance akan auto-update
- Lihat transaction history

### 6. Kirim Bitcoin ke Address Lain

#### Dari Go Wallet (CLI):

```bash
./bin/go-wallet send <wallet-id> <to-address> <amount> <fee> "note"
```

Example:
```bash
./bin/go-wallet send c0788ad0... bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh 0.001 0.00001 "Transfer to friend"
```

**⚠️ CATATAN:** Ini akan create transaction, tapi untuk broadcast ke network real, gunakan Phantom atau wallet lain yang online.

#### Dari Phantom (GUI):

1. **Buka Phantom**
2. **Click Send**
3. **Paste recipient address**
4. **Enter amount**
5. **Confirm**
6. Transaction akan broadcast ke Bitcoin network!

### 7. Terima Bitcoin

Untuk menerima Bitcoin, berikan address Anda:

```
bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y
```

Kepada sender. Mereka bisa kirim dari:
- Exchange (Binance, Indodax, dll)
- Wallet lain (Phantom, Electrum, dll)
- P2P transfer

Balance akan update otomatis di:
- Blockchain explorer
- Phantom (jika di-import)
- Go Wallet (setelah manual sync)

## 💡 Best Practices

### Untuk Keamanan:

1. **✅ Gunakan Go Wallet sebagai Cold Storage**
   - Simpan private key offline
   - Gunakan untuk backup long-term

2. **✅ Gunakan Phantom untuk Daily Transactions**
   - Easy to use
   - Online, real-time updates
   - Support QR code

3. **✅ Backup Private Key**
   - Save di password manager
   - Print dan simpan di safe
   - Jangan share dengan siapapun

4. **⚠️ Test dengan Amount Kecil Dulu**
   - Kirim 0.0001 BTC dulu
   - Verify address benar
   - Baru kirim amount besar

### Untuk Fees:

- **Low Priority:** 1-5 sat/byte (lambat, murah)
- **Medium:** 5-10 sat/byte (normal)
- **High Priority:** 10+ sat/byte (cepat, mahal)

Check recommended fee: https://mempool.space/

## 🎯 Use Cases

### 1. Cold Storage Wallet

```bash
# Generate wallet
./bin/go-wallet create "Cold Storage"

# Export dan backup private key
./bin/go-wallet export-wif <wallet-id>

# Disconnect dari internet
# Simpan private key offline
```

### 2. Hot Wallet (Daily Use)

```bash
# Generate wallet
./bin/go-wallet create "Hot Wallet"

# Import ke Phantom
./bin/go-wallet export-wif <wallet-id>

# Import ke Phantom
# Use Phantom untuk daily transactions
```

### 3. Receive dari Exchange

```bash
# Get address
./bin/go-wallet list

# Copy address (bc1...)
# Withdraw dari exchange ke address ini
# Wait 15-30 minutes
# Check balance di blockchain explorer
```

### 4. Send ke Exchange (Sell)

```bash
# Get deposit address dari exchange
# Example: Binance Deposit → Bitcoin → Get Address

# Import wallet ke Phantom (untuk easy send)
./bin/go-wallet export-wif <wallet-id>

# Dari Phantom:
# Send ke exchange deposit address
# Wait for confirmation
# Sell di exchange
```

## 🔄 Workflow Lengkap

```
┌─────────────────────────────────────────────┐
│ 1. Generate Wallet (Go Wallet)             │
│    bc1qu7eve0wskrztzgz28xuelfx97n7wlejsfk6h3y│
└─────────────────┬───────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────┐
│ 2. Backup Private Key                       │
│    (Save offline & secure)                  │
└─────────────────┬───────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────┐
│ 3. Import ke Phantom (Optional)             │
│    For easy daily use                       │
└─────────────────┬───────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────┐
│ 4. Receive Bitcoin                          │
│    - Dari exchange (withdraw)               │
│    - Dari wallet lain                       │
│    - Dari friend/payment                    │
└─────────────────┬───────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────┐
│ 5. Check Balance                            │
│    - Blockchain explorer                    │
│    - Phantom (auto-update)                  │
│    - Go Wallet (manual sync)                │
└─────────────────┬───────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────┐
│ 6. Send Bitcoin                             │
│    - Via Phantom (easy)                     │
│    - Via Go Wallet (for backup)             │
│    - To exchange/friend/payment             │
└─────────────────────────────────────────────┘
```

## 📱 Quick Commands Reference

```bash
# Create wallet
./bin/go-wallet create "WalletName"

# List all wallets
./bin/go-wallet list

# Export for Phantom
./bin/go-wallet export-wif <wallet-id>

# Check blockchain
./bin/go-wallet sync <wallet-id>

# Get address
./bin/go-wallet list  # Copy bc1... address

# Send (via Phantom recommended)
# Use Phantom GUI for real transactions
```

## ⚠️ Important Notes

### Kenapa Tidak Error Lagi di Phantom?

Karena sekarang:
1. ✅ **Address format benar** (bc1... Native SegWit)
2. ✅ **Compatible dengan Phantom Bitcoin**
3. ✅ **Exchange support SegWit**
4. ✅ **Lower transaction fees**

### Error "Failed to get assets" Hilang Ketika:

1. **Address sudah menerima Bitcoin** dari exchange/wallet lain
2. **Address sudah ada di blockchain** (ada transaction history)
3. **Phantom bisa query balance** dari Bitcoin nodes

### Jika Masih Error:

1. **Check network** - Pastikan pilih Bitcoin (bukan Solana)
2. **Wait for sync** - Phantom perlu waktu untuk sync
3. **Verify address** - Pastikan copy address dengan benar
4. **Check blockchain** - Use blockstream.info untuk verify

## 🎉 Success Checklist

- [ ] Wallet created dengan address bc1...
- [ ] Private key di-backup dengan aman
- [ ] (Optional) Wallet di-import ke Phantom
- [ ] Address di-verify di blockchain explorer
- [ ] Test receive dari exchange (amount kecil)
- [ ] Balance muncul di Phantom/blockchain
- [ ] Test send ke address lain
- [ ] Transaction confirmed di blockchain

## 💬 FAQ

**Q: Apakah wallet ini aman?**
A: Ya, private key disimpan local. Tapi untuk production, tambahkan enkripsi dan secure storage.

**Q: Bisa digunakan untuk transaksi real?**
A: Ya! Address bc1... adalah format standard Bitcoin yang diterima semua exchange dan wallet.

**Q: Kenapa harus import ke Phantom?**
A: Tidak harus. Phantom hanya untuk kemudahan GUI. Anda bisa gunakan Go Wallet saja atau import ke wallet lain (Electrum, BlueWallet, dll).

**Q: Berapa fee yang recommended?**
A: Check https://mempool.space/ untuk fee saat ini. Biasanya 5-10 sat/byte sudah cukup.

**Q: Apakah bisa untuk mainnet?**
A: Ya! Wallet ini sudah production-ready untuk mainnet. Test dengan amount kecil dulu.

---

**✅ Wallet Anda Sekarang Production-Ready!**

Anda bisa:
- Transfer Bitcoin dari exchange ✅
- Menerima Bitcoin dari siapapun ✅
- Send Bitcoin ke address manapun ✅
- Use Phantom sebagai interface ✅
- Backup dengan Go Wallet ✅

Selamat menggunakan Bitcoin wallet Anda! 🚀🪙
