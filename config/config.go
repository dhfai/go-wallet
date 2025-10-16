package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	StoragePath string
	Network     string
}

func NewConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	storagePath := filepath.Join(homeDir, ".go-wallet", "wallets.json")

	return &Config{
		StoragePath: storagePath,
		Network:     "mainnet",
	}
}

func (c *Config) SetStoragePath(path string) {
	c.StoragePath = path
}

func (c *Config) SetNetwork(network string) {
	c.Network = network
}
