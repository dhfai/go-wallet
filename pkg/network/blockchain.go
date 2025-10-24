package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BlockchainExplorer struct {
	baseURL string
	client  *http.Client
}

func NewBlockchainExplorer() *BlockchainExplorer {
	baseURL := "https://blockstream.info/api"

	return &BlockchainExplorer{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type AddressInfo struct {
	Address      string       `json:"address"`
	ChainStats   AddressStats `json:"chain_stats"`
	MempoolStats AddressStats `json:"mempool_stats"`
}

type AddressStats struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}

type UTXOInfo struct {
	TxID   string `json:"txid"`
	Vout   int    `json:"vout"`
	Value  int64  `json:"value"`
	Height int    `json:"status.block_height"`
}

func (be *BlockchainExplorer) GetBalance(address string) (float64, error) {
	url := fmt.Sprintf("%s/address/%s", be.baseURL, address)

	resp, err := be.client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to query blockchain: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("blockchain API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response: %w", err)
	}

	var info AddressInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return 0, fmt.Errorf("failed to parse response: %w", err)
	}

	// Calculate balance: (total received - total spent) from confirmed + mempool
	confirmedBalance := info.ChainStats.FundedTxoSum - info.ChainStats.SpentTxoSum
	mempoolBalance := info.MempoolStats.FundedTxoSum - info.MempoolStats.SpentTxoSum
	totalSatoshis := confirmedBalance + mempoolBalance

	// Convert satoshis to BTC
	btc := float64(totalSatoshis) / 100000000.0
	return btc, nil
}

func (be *BlockchainExplorer) GetUTXOs(address string) ([]UTXOInfo, error) {
	url := fmt.Sprintf("%s/address/%s/utxo", be.baseURL, address)

	resp, err := be.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query blockchain: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("blockchain API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var utxos []UTXOInfo
	if err := json.Unmarshal(body, &utxos); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return utxos, nil
}

func (be *BlockchainExplorer) BroadcastTransaction(txHex string) (string, error) {
	url := fmt.Sprintf("%s/tx", be.baseURL)

	resp, err := be.client.Post(url, "text/plain", nil)
	if err != nil {
		return "", fmt.Errorf("failed to broadcast transaction: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("broadcast failed: %s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(body), nil
}

func (be *BlockchainExplorer) GetTransactionHistory(address string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/address/%s/txs", be.baseURL, address)

	resp, err := be.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query blockchain: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("blockchain API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var txs []map[string]interface{}
	if err := json.Unmarshal(body, &txs); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return txs, nil
}

func (be *BlockchainExplorer) VerifyAddress(address string) (bool, error) {
	url := fmt.Sprintf("%s/address/%s", be.baseURL, address)

	resp, err := be.client.Get(url)
	if err != nil {
		return false, fmt.Errorf("failed to verify address: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

func (be *BlockchainExplorer) GetRecommendedFee() (float64, error) {
	return 0.00001, nil
}
