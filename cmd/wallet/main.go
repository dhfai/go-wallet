package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/dhfai/go-wallet/config"
	"github.com/dhfai/go-wallet/internal/service"
	"github.com/dhfai/go-wallet/internal/storage"
)

func main() {
	cfg := config.NewConfig()

	repo, err := storage.NewJSONWalletRepository(cfg.StoragePath)
	if err != nil {
		fmt.Printf("Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	walletService := service.NewWalletService(repo)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "create":
		handleCreate(walletService)
	case "list":
		handleList(walletService)
	case "balance":
		handleBalance(walletService)
	case "send":
		handleSend(walletService)
	case "receive":
		handleReceive(walletService)
	case "history":
		handleHistory(walletService)
	case "export":
		handleExport(walletService)
	case "import":
		handleImport(walletService)
	case "delete":
		handleDelete(walletService)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Go Bitcoin Wallet - Professional Bitcoin Wallet Management")
	fmt.Println("\nUsage:")
	fmt.Println("  go-wallet <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  create <name>                          Create a new wallet")
	fmt.Println("  list                                   List all wallets")
	fmt.Println("  balance <wallet-id>                    Get wallet balance")
	fmt.Println("  send <from-id> <to-address> <amount> <fee> [note]  Send Bitcoin")
	fmt.Println("  receive <to-id> <from-address> <amount> [note]     Receive Bitcoin")
	fmt.Println("  history <wallet-id> [limit]            Get transaction history")
	fmt.Println("  export <wallet-id>                     Export private key (USE WITH CAUTION!)")
	fmt.Println("  import <name> <private-key>            Import wallet from private key")
	fmt.Println("  delete <wallet-id>                     Delete wallet")
	fmt.Println("  help                                   Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  go-wallet create MyWallet")
	fmt.Println("  go-wallet list")
	fmt.Println("  go-wallet balance abc-123")
	fmt.Println("  go-wallet send abc-123 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa 0.5 0.0001 \"Payment for services\"")
}

func handleCreate(service *service.WalletService) {
	if len(os.Args) < 3 {
		fmt.Println("Error: wallet name is required")
		fmt.Println("Usage: go-wallet create <name>")
		os.Exit(1)
	}

	name := os.Args[2]

	wallet, err := service.CreateWallet(name)
	if err != nil {
		fmt.Printf("Error creating wallet: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Wallet created successfully!")
	fmt.Println("\n=== Wallet Details ===")
	fmt.Printf("ID:         %s\n", wallet.ID)
	fmt.Printf("Name:       %s\n", wallet.Name)
	fmt.Printf("Address:    %s\n", wallet.Address)
	fmt.Printf("Balance:    %.8f BTC\n", wallet.Balance)
	fmt.Printf("Created:    %s\n", wallet.CreatedAt.Format(time.RFC3339))
	fmt.Println("\n⚠️  IMPORTANT: Please backup your private key securely!")
	fmt.Println("Use 'go-wallet export <wallet-id>' to export your private key")
}

func handleList(service *service.WalletService) {
	wallets, err := service.GetAllWallets()
	if err != nil {
		fmt.Printf("Error listing wallets: %v\n", err)
		os.Exit(1)
	}

	if len(wallets) == 0 {
		fmt.Println("No wallets found. Create one with: go-wallet create <name>")
		return
	}

	fmt.Printf("\n=== Wallets (%d total) ===\n\n", len(wallets))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tAddress\tBalance (BTC)\tTransactions")
	fmt.Fprintln(w, "----\t----\t----\t----\t----")

	for _, wallet := range wallets {
		fmt.Fprintf(w, "%s\t%s\t%s\t%.8f\t%d\n",
			wallet.ID,
			wallet.Name,
			wallet.Address,
			wallet.Balance,
			len(wallet.Transactions),
		)
	}

	w.Flush()
}

func handleBalance(service *service.WalletService) {
	if len(os.Args) < 3 {
		fmt.Println("Error: wallet ID is required")
		fmt.Println("Usage: go-wallet balance <wallet-id>")
		os.Exit(1)
	}

	walletID := os.Args[2]

	balance, err := service.GetBalance(walletID)
	if err != nil {
		fmt.Printf("Error getting balance: %v\n", err)
		os.Exit(1)
	}

	wallet, _ := service.GetWallet(walletID)

	fmt.Printf("\n=== Wallet Balance ===\n")
	fmt.Printf("Wallet:  %s (%s)\n", wallet.Name, wallet.ID)
	fmt.Printf("Address: %s\n", wallet.Address)
	fmt.Printf("Balance: %.8f BTC\n", balance)
}

func handleSend(service *service.WalletService) {
	if len(os.Args) < 6 {
		fmt.Println("Error: insufficient arguments")
		fmt.Println("Usage: go-wallet send <from-id> <to-address> <amount> <fee> [note]")
		os.Exit(1)
	}

	fromID := os.Args[2]
	toAddress := os.Args[3]
	amount, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Printf("Error: invalid amount: %v\n", err)
		os.Exit(1)
	}

	fee, err := strconv.ParseFloat(os.Args[5], 64)
	if err != nil {
		fmt.Printf("Error: invalid fee: %v\n", err)
		os.Exit(1)
	}

	note := ""
	if len(os.Args) > 6 {
		note = os.Args[6]
	}

	tx, err := service.SendBitcoin(fromID, toAddress, amount, fee, note)
	if err != nil {
		fmt.Printf("Error sending Bitcoin: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Transaction sent successfully!")
	fmt.Println("\n=== Transaction Details ===")
	fmt.Printf("TX ID:      %s\n", tx.ID)
	fmt.Printf("From:       %s\n", tx.From)
	fmt.Printf("To:         %s\n", tx.To)
	fmt.Printf("Amount:     %.8f BTC\n", tx.Amount)
	fmt.Printf("Fee:        %.8f BTC\n", tx.Fee)
	fmt.Printf("Status:     %s\n", tx.Status)
	fmt.Printf("Time:       %s\n", tx.Timestamp.Format(time.RFC3339))
	if tx.Note != "" {
		fmt.Printf("Note:       %s\n", tx.Note)
	}
}

func handleReceive(service *service.WalletService) {
	if len(os.Args) < 5 {
		fmt.Println("Error: insufficient arguments")
		fmt.Println("Usage: go-wallet receive <to-id> <from-address> <amount> [note]")
		os.Exit(1)
	}

	toID := os.Args[2]
	fromAddress := os.Args[3]
	amount, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Printf("Error: invalid amount: %v\n", err)
		os.Exit(1)
	}

	note := ""
	if len(os.Args) > 5 {
		note = os.Args[5]
	}

	tx, err := service.ReceiveBitcoin(toID, fromAddress, amount, note)
	if err != nil {
		fmt.Printf("Error receiving Bitcoin: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Bitcoin received successfully!")
	fmt.Println("\n=== Transaction Details ===")
	fmt.Printf("TX ID:      %s\n", tx.ID)
	fmt.Printf("From:       %s\n", tx.From)
	fmt.Printf("To:         %s\n", tx.To)
	fmt.Printf("Amount:     %.8f BTC\n", tx.Amount)
	fmt.Printf("Status:     %s\n", tx.Status)
	fmt.Printf("Time:       %s\n", tx.Timestamp.Format(time.RFC3339))
	if tx.Note != "" {
		fmt.Printf("Note:       %s\n", tx.Note)
	}
}

func handleHistory(service *service.WalletService) {
	if len(os.Args) < 3 {
		fmt.Println("Error: wallet ID is required")
		fmt.Println("Usage: go-wallet history <wallet-id> [limit]")
		os.Exit(1)
	}

	walletID := os.Args[2]
	limit := 0

	if len(os.Args) > 3 {
		var err error
		limit, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Error: invalid limit: %v\n", err)
			os.Exit(1)
		}
	}

	transactions, err := service.GetTransactionHistory(walletID, limit)
	if err != nil {
		fmt.Printf("Error getting transaction history: %v\n", err)
		os.Exit(1)
	}

	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Printf("\n=== Transaction History (%d transactions) ===\n\n", len(transactions))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Time\tType\tAmount (BTC)\tFrom/To\tStatus")
	fmt.Fprintln(w, "----\t----\t----\t----\t----")

	for _, tx := range transactions {
		address := tx.To
		if tx.Type == "receive" {
			address = tx.From
		}

		fmt.Fprintf(w, "%s\t%s\t%.8f\t%s...\t%s\n",
			tx.Timestamp.Format("2006-01-02 15:04"),
			tx.Type,
			tx.Amount,
			address[:10],
			tx.Status,
		)
	}

	w.Flush()
}

func handleExport(service *service.WalletService) {
	if len(os.Args) < 3 {
		fmt.Println("Error: wallet ID is required")
		fmt.Println("Usage: go-wallet export <wallet-id>")
		os.Exit(1)
	}

	walletID := os.Args[2]

	privateKey, err := service.ExportPrivateKey(walletID)
	if err != nil {
		fmt.Printf("Error exporting private key: %v\n", err)
		os.Exit(1)
	}

	wallet, _ := service.GetWallet(walletID)

	fmt.Println("\n⚠️  WARNING: KEEP THIS PRIVATE KEY SECURE!")
	fmt.Println("Anyone with this key can access your funds.")
	fmt.Println("\n=== Private Key Export ===")
	fmt.Printf("Wallet:      %s (%s)\n", wallet.Name, wallet.ID)
	fmt.Printf("Address:     %s\n", wallet.Address)
	fmt.Printf("Private Key: %s\n", privateKey)
	fmt.Println("\n⚠️  Do NOT share this key with anyone!")
}

func handleImport(service *service.WalletService) {
	if len(os.Args) < 4 {
		fmt.Println("Error: insufficient arguments")
		fmt.Println("Usage: go-wallet import <name> <private-key>")
		os.Exit(1)
	}

	name := os.Args[2]
	privateKey := os.Args[3]

	wallet, err := service.ImportWallet(name, privateKey)
	if err != nil {
		fmt.Printf("Error importing wallet: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Wallet imported successfully!")
	fmt.Println("\n=== Wallet Details ===")
	fmt.Printf("ID:         %s\n", wallet.ID)
	fmt.Printf("Name:       %s\n", wallet.Name)
	fmt.Printf("Address:    %s\n", wallet.Address)
	fmt.Printf("Balance:    %.8f BTC\n", wallet.Balance)
}

func handleDelete(service *service.WalletService) {
	if len(os.Args) < 3 {
		fmt.Println("Error: wallet ID is required")
		fmt.Println("Usage: go-wallet delete <wallet-id>")
		os.Exit(1)
	}

	walletID := os.Args[2]

	wallet, err := service.GetWallet(walletID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = service.DeleteWallet(walletID)
	if err != nil {
		fmt.Printf("Error deleting wallet: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Wallet deleted successfully!\n")
	fmt.Printf("Deleted: %s (%s)\n", wallet.Name, wallet.Address)
	fmt.Println("\n⚠️  Make sure you have backed up the private key if needed!")
}
