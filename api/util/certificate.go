package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

const (
	CERT         = "cert.pem"
	KEY          = "key.pem"
	STORE_DIR    = "store/"
	FRONTEND_DIR = "../frontend"
)

func NewCertificateIfDoesNotExist() {
	if _, err := os.Stat(STORE_DIR); os.IsNotExist(err) {
		err := os.Mkdir(STORE_DIR, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(CERT); os.IsNotExist(err) {
		err = generateCertificate()
		if err != nil {
			log.Fatal(err)
		}
		err = copyCertToFrontend()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func generateCertificate() error {
	// Generate a new RSA key pair.
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Create a new self-signed X.509 certificate.
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "emidev98"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // Expires in 10 years
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return err
	}

	// Write the certificate and key to disk.
	certFile, err := os.Create(STORE_DIR + CERT)
	if err != nil {
		return err
	}
	defer certFile.Close()

	keyFile, err := os.Create(STORE_DIR + KEY)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		return err
	}

	err = pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return err
	}

	return nil
}

func copyCertToFrontend() error {
	// remove ../frontend/store folder if exists
	if _, err := os.Stat(filepath.Join(FRONTEND_DIR, STORE_DIR)); !os.IsNotExist(err) {
		err := os.RemoveAll(filepath.Join(FRONTEND_DIR, STORE_DIR))
		if err != nil {
			return err
		}
	}
	// create the directory ../frontend/store
	err := os.MkdirAll(filepath.Join(FRONTEND_DIR, STORE_DIR), 0755)
	if err != nil {
		return err
	}
	// Copy the files ./store/cert.pem and ./store/key.pem to the ../frontend/store folder
	err = copyFile(STORE_DIR+KEY, filepath.Join(FRONTEND_DIR, STORE_DIR+KEY))
	if err != nil {
		return err
	}
	err = copyFile(STORE_DIR+CERT, filepath.Join(FRONTEND_DIR, STORE_DIR+CERT))
	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dest string) error {
	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		return err
	}
	return nil
}
