package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	// Se carga los archivos del certificado generado,
	// el cliente debe generar un certificado autofirmado
	// o que el servidor genere un certificado para el cliente
	cert, err := tls.LoadX509KeyPair("cert/gotest.pem", "cert/key.pem")
	if err != nil {
		log.Fatalf("Error loading certificateds %v ", err)
	}

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
		ServerName:         "localhost",
		ClientAuth:         tls.RequireAnyClientCert,
	}

	listener, err := net.Listen("tcp", ":6300")
	if err != nil {
		log.Fatalf("Listener errror: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Accept server error: %v", err)
		}

		tlsConn := tls.Server(conn, config)
		err = tlsConn.Handshake()

		if err != nil {
			log.Fatalf("ERROR EN CONEXION: %v", err)
		}

		buffer := make([]byte, 1024)
		_, err = tlsConn.Read(buffer)
		if err != nil {
			log.Fatalf("Error reading from buffer conn: %v", err)
		}

		fmt.Printf("Mensaje recibido: %s\n", string(buffer))
	}
}
