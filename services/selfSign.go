package services

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/rs/zerolog/log"
	"math/big"
	"os"
	"time"
)

func GenerateKeys() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate private key")
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate serial number")
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"KSTU"},
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create certificate")
	}

	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal().Err(err).Msg("Failed to encode certificate to PEM")
	}

	if err = os.WriteFile("cert.pem", pemCert, 0644); err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msg("wrote cert.pem\n")

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshal private key")
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemKey == nil {
		log.Fatal().Err(err).Msg("Failed to encode key to PEM")
	}
	if err = os.WriteFile("key.pem", pemKey, 0600); err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msg("wrote key.pem\n")
}
