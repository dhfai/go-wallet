package domain

import "errors"

var (
	ErrWalletNotFound = errors.New("wallet not found")

	ErrInsufficientBalance = errors.New("insufficient balance")

	ErrInvalidAddress = errors.New("invalid Bitcoin address")

	ErrInvalidAmount = errors.New("invalid amount")

	ErrWalletExists = errors.New("wallet already exists")

	ErrInvalidPrivateKey = errors.New("invalid private key")

	ErrKeyGeneration = errors.New("failed to generate keys")

	ErrStorageOperation = errors.New("storage operation failed")
)
