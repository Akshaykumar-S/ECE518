// crypto/crypto.go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func validateSHA256() {
	msg := "Hello world!"
	hashExp, _ := hex.DecodeString(
		"c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a")

	sha := sha256.New()
	sha.Write([]byte(msg))

	hash := sha.Sum(nil)

	fmt.Printf("Validated %t: SHA256(%s)=%x\n",
		bytes.Compare(hash, hashExp) == 0,
		msg, hash)
}

func validateAESGCM() {
	plaintxt := "Hello world!"
	data := "ece443"

	key := make([]byte, 32) //aes key
	rand.Read(key)          //random values to fill up the array/string

	block, _ := aes.NewCipher(key) // 128 bits
	fmt.Printf("block len %d, \n", block.BlockSize())
	aesgcm, _ := cipher.NewGCM(block) // nonce 96 bits
	fmt.Printf("aes gcm %d \n", aesgcm.NonceSize())

	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce) // 96 bit random string

	ciphermac := aesgcm.Seal(nil, nonce, []byte(plaintxt), []byte(data)) // 28 bytes
	// ciphermac because seal function returns the ciphertext and message authentication code

	fmt.Printf("size %d\n", len(ciphermac))

	pbuf, err := aesgcm.Open(nil, nonce, ciphermac, []byte(data))

	fmt.Printf("Validated %t: AES-GCM(%s, data=%s, nonce=%x, key=%x)=%x\n",
		err == nil, string(pbuf), data, nonce, key, ciphermac)
}
