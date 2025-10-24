package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

type BitcoinCrypto struct{}

func NewBitcoinCrypto() *BitcoinCrypto {
	return &BitcoinCrypto{}
}

func (bc *BitcoinCrypto) GenerateKeyPair() (privateKey, publicKey string, err error) {

	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %w", err)
	}

	privateKeyBytes := privKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	pubKeyBytes := elliptic.Marshal(privKey.PublicKey.Curve, privKey.PublicKey.X, privKey.PublicKey.Y)
	publicKeyHex := hex.EncodeToString(pubKeyBytes)

	return privateKeyHex, publicKeyHex, nil
}

func (bc *BitcoinCrypto) GenerateAddress(publicKeyHex string) (string, error) {

	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid public key hex: %w", err)
	}

	sha256Hash := sha256.Sum256(publicKeyBytes)

	ripemd160Hasher := ripemd160.New()
	_, err = ripemd160Hasher.Write(sha256Hash[:])
	if err != nil {
		return "", fmt.Errorf("ripemd160 hash failed: %w", err)
	}
	publicKeyHash := ripemd160Hasher.Sum(nil)

	versionedPayload := append([]byte{0x00}, publicKeyHash...)

	firstHash := sha256.Sum256(versionedPayload)
	secondHash := sha256.Sum256(firstHash[:])
	checksum := secondHash[:4]

	fullPayload := append(versionedPayload, checksum...)

	address := bc.base58Encode(fullPayload)

	return address, nil
}

func (bc *BitcoinCrypto) PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key hex: %w", err)
	}

	privKey := new(ecdsa.PrivateKey)
	privKey.PublicKey.Curve = elliptic.P256()
	privKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privKey.PublicKey.X, privKey.PublicKey.Y = privKey.PublicKey.Curve.ScalarBaseMult(privateKeyBytes)

	return privKey, nil
}

func (bc *BitcoinCrypto) SignTransaction(txHash string, privateKeyHex string) (string, error) {
	privKey, err := bc.PrivateKeyFromHex(privateKeyHex)
	if err != nil {
		return "", err
	}

	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return "", fmt.Errorf("invalid transaction hash: %w", err)
	}

	r, s, err := ecdsa.Sign(rand.Reader, privKey, txHashBytes)
	if err != nil {
		return "", fmt.Errorf("signing failed: %w", err)
	}

	signature := append(r.Bytes(), s.Bytes()...)
	signatureHex := hex.EncodeToString(signature)

	return signatureHex, nil
}

func (bc *BitcoinCrypto) VerifySignature(txHash, signatureHex, publicKeyHex string) (bool, error) {

	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return false, fmt.Errorf("invalid public key hex: %w", err)
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), publicKeyBytes)
	if x == nil {
		return false, fmt.Errorf("invalid public key")
	}

	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	signatureBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, fmt.Errorf("invalid signature hex: %w", err)
	}

	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])

	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return false, fmt.Errorf("invalid transaction hash: %w", err)
	}

	valid := ecdsa.Verify(pubKey, txHashBytes, r, s)

	return valid, nil
}

func (bc *BitcoinCrypto) base58Encode(input []byte) string {
	const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz" // Bitcoin Base58 alphabets

	x := new(big.Int).SetBytes(input)

	var result []byte
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	for _, b := range input {
		if b != 0 {
			break
		}
		result = append(result, base58Alphabet[0])
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func (bc *BitcoinCrypto) HashTransaction(from, to string, amount float64, timestamp int64) string {
	data := fmt.Sprintf("%s%s%f%d", from, to, amount, timestamp)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
