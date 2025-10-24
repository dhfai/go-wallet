package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

// GenerateSegWitAddress generates Native SegWit (bech32) address
// GenerateSegWitAddress menghasilkan alamat Native SegWit (bech32)
func (bc *BitcoinCrypto) GenerateSegWitAddress(publicKeyHex string) (string, error) {
	// Decode public key
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}

	// SHA-256 hash
	sha256Hash := sha256.Sum256(publicKeyBytes)

	// RIPEMD-160 hash
	ripemd160Hasher := ripemd160.New()
	_, err = ripemd160Hasher.Write(sha256Hash[:])
	if err != nil {
		return "", err
	}
	pubKeyHash := ripemd160Hasher.Sum(nil)

	// Encode to bech32 (witness version 0, mainnet "bc")
	address, err := bc.encodeBech32("bc", 0, pubKeyHash)
	if err != nil {
		return "", fmt.Errorf("failed to encode bech32: %w", err)
	}

	return address, nil
}

// GenerateSegWitTestnetAddress generates Native SegWit address for testnet
// GenerateSegWitTestnetAddress menghasilkan alamat Native SegWit untuk testnet
func (bc *BitcoinCrypto) GenerateSegWitTestnetAddress(publicKeyHex string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}

	sha256Hash := sha256.Sum256(publicKeyBytes)

	ripemd160Hasher := ripemd160.New()
	_, err = ripemd160Hasher.Write(sha256Hash[:])
	if err != nil {
		return "", err
	}
	pubKeyHash := ripemd160Hasher.Sum(nil)

	// Use "tb" for testnet bech32 addresses
	address, err := bc.encodeBech32("tb", 0, pubKeyHash)
	if err != nil {
		return "", fmt.Errorf("failed to encode bech32: %w", err)
	}

	return address, nil
}

// encodeBech32 encodes data to bech32 format
func (bc *BitcoinCrypto) encodeBech32(hrp string, version byte, program []byte) (string, error) {
	// Convert 8-bit to 5-bit
	converted := bc.convertBits(program, 8, 5, true)
	if converted == nil {
		return "", fmt.Errorf("failed to convert bits")
	}

	// Prepend witness version
	data := append([]byte{version}, converted...)

	// Create checksum
	checksum := bc.bech32Checksum(hrp, data)
	combined := append(data, checksum...)

	// Encode with bech32 charset
	const charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
	result := hrp + "1"
	for _, b := range combined {
		if int(b) >= len(charset) {
			return "", fmt.Errorf("invalid data")
		}
		result += string(charset[b])
	}

	return result, nil
}

// convertBits converts between bit groups
func (bc *BitcoinCrypto) convertBits(data []byte, fromBits, toBits int, pad bool) []byte {
	acc := 0
	bits := 0
	var result []byte
	maxv := (1 << toBits) - 1

	for _, value := range data {
		acc = (acc << fromBits) | int(value)
		bits += fromBits

		for bits >= toBits {
			bits -= toBits
			result = append(result, byte((acc>>bits)&maxv))
		}
	}

	if pad {
		if bits > 0 {
			result = append(result, byte((acc<<(toBits-bits))&maxv))
		}
	} else if bits >= fromBits || ((acc<<(toBits-bits))&maxv) != 0 {
		return nil
	}

	return result
}

// bech32Checksum creates bech32 checksum
func (bc *BitcoinCrypto) bech32Checksum(hrp string, data []byte) []byte {
	values := bc.bech32HrpExpand(hrp)
	values = append(values, data...)
	values = append(values, []byte{0, 0, 0, 0, 0, 0}...)

	polymod := bc.bech32Polymod(values) ^ 1
	var checksum []byte

	for i := 0; i < 6; i++ {
		checksum = append(checksum, byte((polymod>>uint(5*(5-i)))&31))
	}

	return checksum
}

// bech32HrpExpand expands HRP for checksum
func (bc *BitcoinCrypto) bech32HrpExpand(hrp string) []byte {
	var result []byte
	for _, c := range hrp {
		result = append(result, byte(c>>5))
	}
	result = append(result, 0)
	for _, c := range hrp {
		result = append(result, byte(c&31))
	}
	return result
}

// bech32Polymod calculates bech32 polymod
func (bc *BitcoinCrypto) bech32Polymod(values []byte) int {
	gen := []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
	chk := 1

	for _, value := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ int(value)

		for i := 0; i < 5; i++ {
			if (top>>uint(i))&1 == 1 {
				chk ^= gen[i]
			}
		}
	}

	return chk
}

// ConvertToWIF converts hex private key to WIF format for import
// ConvertToWIF mengkonversi private key hex ke format WIF untuk import
func (bc *BitcoinCrypto) ConvertToWIF(privateKeyHex string, compressed bool, testnet bool) (string, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid hex: %w", err)
	}

	// Version byte: 0x80 for mainnet, 0xef for testnet
	versionByte := byte(0x80)
	if testnet {
		versionByte = 0xef
	}

	payload := append([]byte{versionByte}, privateKeyBytes...)

	if compressed {
		payload = append(payload, 0x01)
	}

	// Double SHA-256 for checksum
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	checksum := secondHash[:4]

	fullPayload := append(payload, checksum...)
	wif := bc.base58Encode(fullPayload)

	return wif, nil
}

// base58Decode decodes Base58 string
func (bc *BitcoinCrypto) base58Decode(input string) []byte {
	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	reverseMap := make(map[rune]int)
	for i, c := range alphabet {
		reverseMap[c] = i
	}

	result := big.NewInt(0)
	base := big.NewInt(58)

	for _, c := range input {
		val, ok := reverseMap[c]
		if !ok {
			return nil
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(val)))
	}

	decoded := result.Bytes()

	// Add leading zeros
	for _, c := range input {
		if c != '1' {
			break
		}
		decoded = append([]byte{0}, decoded...)
	}

	return decoded
}
