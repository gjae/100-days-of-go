package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"log"
	"math/big"
	"os"
	"time"
)

func prepareCert() *x509.Certificate {
	// Cargar en memopria el certificado
	certBytes, err := os.ReadFile("cert/gotest.pem")
	if err != nil {
		panic(err)
	}
	key, _ := pem.Decode(certBytes)
	serverCert, err := x509.ParseCertificate(key.Bytes)
	if err != nil {
		panic(err)
	}

	return serverCert
}

func ExistsLocalCert() bool {
	cert, err := os.Open("cert/new_pcert.pem")
	defer cert.Close()

	return err != nil
}

// En caso de no existir un certificado creado en el sistema
// se crea uno nuevo self-signed
func NewCert(certificated *x509.Certificate, pkey *pem.Block) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "System"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:     x509.KeyUsageDataEncipherment | x509.KeyUsageDigitalSignature,
		IsCA:         false,
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(pkey.Bytes)
	if err != nil {
		panic(err)
	}

	clientCertDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)

	if err != nil {
		log.Fatalf("Error certificated error: %v", err)
		panic(err)
	}

	certOut, err := os.Create("cert/new_pcert.pem")
	if errors.Is(os.ErrExist, err) {
		return
	}
	defer certOut.Close()

	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: clientCertDER})
}

func createNewCertificate(certificated *x509.Certificate, pkey *pem.Block) {
	if ExistsLocalCert() {
		NewCert(certificated, pkey)
	}
}

func NewPrivateKey() *pem.Block {
	// Generar nueva clave privada random
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Codificar clave privada en PEM
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Escribir clave privada en nuevo archivo
	privateFile, err := os.Create("cert/new_private_key.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()

	if err := pem.Encode(privateFile, privateKeyPEM); err != nil {
		panic(err)
	}

	return privateKeyPEM
}

// Si no existe una clave privada ya creada entonces
// se crea una nueva
func CurrentPrivateKey() (*pem.Block, error) {
	parse, err := os.ReadFile("cert/new_private_key.pem")
	var block *pem.Block
	if err != nil {
		block = NewPrivateKey()
	} else {
		block, _ = pem.Decode(parse)
	}

	return block, nil
}

func main() {
	privateKey, err := CurrentPrivateKey()
	if err != nil {
		panic(err)
	}

	createNewCertificate(prepareCert(), privateKey)
	cert, err := tls.LoadX509KeyPair("cert/new_pcert.pem", "cert/new_private_key.pem")

	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", ":6300", config)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello world!"))
	if err != nil {
		log.Fatalf("Send message error: %v", err)
	}
}
