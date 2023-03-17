package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"

	"io"
)


func main(){
	fmt.Println("Encryption program")

	text := []byte("My super secred code stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	
	fmt.Println(len(key))
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or galois/counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)

	nonce := make([]byte, gcm.NonceSize())
	// if any error generating new GCM
	// handle them
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}



	fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	err = os.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 077)

	if err != nil {

		panic(err)
	}
}
