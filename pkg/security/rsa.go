package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// SecretKey is struct of secret key
type SecretKey struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

const (
	// rsaKeySize is constant number of bits for RSA
	rsaKeySize = 2048
)

// GenerateSecretKey is a function to generate secret key base64 based.
func GenerateSecretKey() (*SecretKey, error) {
	priKey, pubKey, err := GenerateKey64()
	if err != nil {
		return nil, err
	}

	return &SecretKey{
		PrivateKey: base64.StdEncoding.EncodeToString([]byte(priKey)),
		PublicKey:  base64.StdEncoding.EncodeToString([]byte(pubKey)),
	}, nil
}

// GenerateKey is function to generate random public & private key
func GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	pri, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return nil, nil, err
	}

	return pri, &pri.PublicKey, nil
}

// GenerateKeyBytes is function to generate bytes key
func GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	priBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return nil, nil, err
	}

	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, nil, err
	}

	return priBytes, pubBytes, nil
}

// GenerateKey64 is function to generate base64 key
func GenerateKey64() (pri, pub string, err error) {
	priBytes, pubBytes, err := GenerateKeyBytes()
	if err != nil {
		return "", "", err
	}

	priKey := base64.StdEncoding.EncodeToString(priBytes)
	pubKey := base64.StdEncoding.EncodeToString(pubBytes)

	priKeyPem := fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----\n", priKey)
	pubKeyPem := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----\n", pubKey)

	return priKeyPem, pubKeyPem, nil
}
