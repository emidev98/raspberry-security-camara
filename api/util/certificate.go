package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

const (
	STORE_DIR = "store/"
	CERT      = "cert.pem"
	KEY       = "key.pem"
)

func NewCertificateIfDoesNotExist() {
	if _, err := os.Stat(STORE_DIR); os.IsNotExist(err) {
		err := os.Mkdir(STORE_DIR, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(CERT); os.IsNotExist(err) {
		generateCertificate()
	}
}

func generateCertificate() {
	// Generate a new RSA key pair.
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new self-signed X.509 certificate.
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // Expires in 10 years
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		log.Fatal(err)
	}

	// Write the certificate and key to disk.
	certFile, err := os.Create(STORE_DIR + CERT)
	if err != nil {
		log.Fatal(err)
	}
	defer certFile.Close()

	keyFile, err := os.Create(STORE_DIR + KEY)
	if err != nil {
		log.Fatal(err)
	}
	defer keyFile.Close()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		log.Fatal(err)
	}

	err = pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		log.Fatal(err)
	}
}
