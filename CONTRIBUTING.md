# Contributing to Go Bitcoin Wallet

Terima kasih atas minat Anda untuk berkontribusi! 🎉

## Cara Berkontribusi

### Melaporkan Bug

1. Pastikan bug belum pernah dilaporkan dengan mencari di [Issues](https://github.com/dhfai/go-wallet/issues)
2. Buat issue baru dengan:
   - Deskripsi jelas tentang masalahnya
   - Steps to reproduce
   - Expected vs actual behavior
   - Go version dan OS yang digunakan
   - Log errors jika ada

### Mengusulkan Fitur Baru

1. Buat issue dengan label "feature request"
2. Jelaskan use case dan manfaatnya
3. Diskusikan implementasi dengan maintainer

### Pull Request Process

1. **Fork** repository
2. **Clone** fork Anda:
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-wallet.git
   ```

3. **Create branch** untuk fitur/bugfix:
   ```bash
   git checkout -b feature/nama-fitur
   ```

4. **Implement** perubahan Anda:
   - Follow coding standards (lihat di bawah)
   - Tambahkan tests jika perlu
   - Update dokumentasi

5. **Test** perubahan Anda:
   ```bash
   make test
   make lint
   ```

6. **Commit** dengan pesan yang jelas:
   ```bash
   git commit -m "feat: add new wallet export format"
   ```

7. **Push** ke fork Anda:
   ```bash
   git push origin feature/nama-fitur
   ```

8. **Create Pull Request** dari fork Anda

## Coding Standards

### Go Style Guide

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `go fmt` untuk formatting
- Use `golangci-lint` untuk linting
- Maximum line length: 120 characters

### Commit Messages

Format: `<type>: <description>`

Types:
- `feat`: Fitur baru
- `fix`: Bug fix
- `docs`: Dokumentasi
- `style`: Formatting, semicolons, dll
- `refactor`: Code refactoring
- `test`: Menambah tests
- `chore`: Maintenance tasks

Examples:
```
feat: add multi-signature wallet support
fix: resolve balance calculation bug
docs: update README with new examples
```

### Code Structure

- Gunakan **Clean Architecture**
- Separate concerns (domain, service, storage)
- Dependency injection pattern
- Interface-based design

### Testing

- Write unit tests untuk business logic
- Integration tests untuk end-to-end flows
- Minimum coverage: 70%
- Use table-driven tests

Example:
```go
func TestCreateWallet(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
    }{
        {"valid name", "MyWallet", false},
        {"empty name", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Documentation

- Semua public functions harus ada godoc comments
- Gunakan bahasa Indonesia dan Inggris
- Include examples untuk complex functions

Example:
```go
// CreateWallet creates a new Bitcoin wallet with generated keys
// CreateWallet membuat wallet Bitcoin baru dengan kunci yang di-generate
//
// Parameters:
//   - name: User-friendly name for the wallet
//
// Returns:
//   - *Wallet: Created wallet instance
//   - error: Error if creation fails
//
// Example:
//   wallet, err := service.CreateWallet("MyWallet")
//   if err != nil {
//       log.Fatal(err)
//   }
func (s *WalletService) CreateWallet(name string) (*Wallet, error) {
    // implementation
}
```

## Development Setup

1. **Install Go 1.21+**
2. **Install tools**:
   ```bash
   go install golang.org/x/tools/cmd/goimports@latest
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

3. **Clone & setup**:
   ```bash
   git clone https://github.com/dhfai/go-wallet.git
   cd go-wallet
   make deps
   ```

4. **Run tests**:
   ```bash
   make test
   ```

## Project Structure

```
go-wallet/
├── cmd/          - Applications
├── internal/     - Private code
│   ├── domain/   - Business entities
│   ├── service/  - Business logic
│   └── storage/  - Data persistence
├── pkg/          - Public libraries
└── config/       - Configuration
```

## Questions?

Jangan ragu untuk bertanya di:
- [GitHub Discussions](https://github.com/dhfai/go-wallet/discussions)
- [Issues](https://github.com/dhfai/go-wallet/issues)

## Code of Conduct

- Be respectful
- Be collaborative
- Be professional
- Focus on the code, not the person

---

Terima kasih atas kontribusi Anda! 🙏
