package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/google/uuid"
	"software.sslmate.com/src/go-pkcs12"
)

func main() {
	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Generate serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(err)
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 180),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Generate certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}

	// Encode private key and certificate to PEM format
	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	privPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes})
	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})

	// Write private key and certificate to files
	privFile, err := os.Create(uuid.NewString() + ".key")
	if err != nil {
		panic(err)
	}
	defer privFile.Close()
	if _, err := privFile.Write(privPem); err != nil {
		panic(err)
	}

	certFile, err := os.Create(uuid.NewString() + ".crt")
	if err != nil {
		panic(err)
	}
	defer certFile.Close()
	if _, err := certFile.Write(certPem); err != nil {
		panic(err)
	}

	// Generate .pfx file
	pfxFile, err := os.Create(uuid.NewString() + ".pfx")
	if err != nil {
		panic(err)
	}
	defer pfxFile.Close()

	pfxBytes, err := generatePFX(priv, certBytes, "password")
	if err != nil {
		panic(err)
	}

	if _, err := pfxFile.Write(pfxBytes); err != nil {
		panic(err)
	}

	fmt.Println("Written A1 certificate files:")
	fmt.Println("Private Key:", privFile.Name())
	fmt.Println("Certificate:", certFile.Name())
	fmt.Println("PFX File:", pfxFile.Name())
}

func generatePFX(priv *rsa.PrivateKey, certBytes []byte, passphrase string) ([]byte, error) {
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	pfx, err := pkcs12.Encode(rand.Reader, priv, cert, nil, passphrase)
	if err != nil {
		return nil, err
	}

	return pfx, nil
}
