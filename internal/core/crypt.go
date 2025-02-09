package core

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
)

func LoadECCPrivateKey(pemFile string) (*ecdsa.PrivateKey, error) {
	keyData, err := os.ReadFile(pemFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	return privateKey, nil
}

func ComputeFileChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func SignChecksum(privateKey *ecdsa.PrivateKey, checksum string) (string, error) {
	hash := sha256.Sum256([]byte(checksum))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	rBytes := r.Bytes()
	sBytes := s.Bytes()
	signature := hex.EncodeToString(append(rBytes, sBytes...))

	return signature, nil
}

func LoadECCPublicKey(pemFile string) (*ecdsa.PublicKey, error) {
	keyData, err := os.ReadFile(pemFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC public key: %w", err)
	}

	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid ECDSA public key")
	}

	return ecdsaPubKey, nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, checksum, signatureHex string) bool {
	hash := sha256.Sum256([]byte(checksum))

	// Decode signature from hex
	signatureBytes, err := hex.DecodeString(signatureHex)
	if err != nil || len(signatureBytes)%2 != 0 {
		fmt.Println("Invalid signature format")
		return false
	}

	r := new(big.Int).SetBytes(signatureBytes[:len(signatureBytes)/2])
	s := new(big.Int).SetBytes(signatureBytes[len(signatureBytes)/2:])

	return ecdsa.Verify(publicKey, hash[:], r, s)
}
