// crypto/crypto.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	//"time"
)

func findPassword() {
	nonce := make([]byte, 12)
	data := "jwang34@iit.edu"

	ciphermac, _ := hex.DecodeString(
		"7d4eb640daf43844bd605bb1c2a66046fd33cba7d2ba35828d25056c953834d93b9a04c54fa86147f62a")

	fmt.Printf("\n\n\nfinding password for nonce=%x, data=%s, ciphermac=%x...\n",
		nonce, data, ciphermac)

	for i := 0; i < 10000; i++ {
		password := fmt.Sprintf("%04d", i)
		// Modify code below to check that if sha256(password) is a correct
		// key to decrypt the ciphertext with MAC in ciphermac, which is
		// obtained via AES-GCM with a 0 nouce and the data being jwang34@iit.edu
		//
		// Once found, show the correct password as well as the plaintext.

		// Code to find the sha256 of password
		sha := sha256.New()
		sha.Write([]byte(password))
		hash := sha.Sum(nil)

		// Advanced Encryption Standard - Galois / Counter Mode
		block, _ := aes.NewCipher(hash)
		aesgcm, _ := cipher.NewGCM(block)
		pbuf, err := aesgcm.Open(nil, nonce, ciphermac, []byte(data))
		if err == nil {
			fmt.Printf(" The correct password is %s\n", password)
			fmt.Printf(" The original Message is %s\n", pbuf)
			break
		}
	}
}

func timetocomputeSHA256() time.Duration {
	startTime := time.Now()
	message := make([]byte, 16)
	rand.Read(message)

	sha := sha256.New()
	sha.Write([]byte(message))
	sha.Sum(nil)

	return time.Since(startTime)
}

func initializeAES() (time.Duration, time.Duration) {
	plaintext := make([]byte, 1024)
	rand.Read(plaintext)

	key := make([]byte, 32) //aes key
	rand.Read(key)

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce)

	startTimeE := time.Now()
	ciphermac := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	durationE := time.Since(startTimeE)

	startTimeD := time.Now()
	aesgcm.Open(nil, nonce, ciphermac, nil)
	durationD := time.Since(startTimeD)

	return durationE, durationD
}

func performanceEvaluation() {
	var itr int64 = 1000000
	var SHA256sum int64 = 0
	var Encryptsum int64 = 0
	var Decryptsum int64 = 0
	for i := 0; i < 1000000; i++ {
		SHA256sum = SHA256sum + timetocomputeSHA256().Nanoseconds()
		durationE, durationD := initializeAES()
		Encryptsum = Encryptsum + durationE.Nanoseconds()
		Decryptsum = Decryptsum + durationD.Nanoseconds()
	}

	SHA256time := SHA256sum / itr
	Encrypttime := Encryptsum / itr
	Decrypttime := Decryptsum / itr
	fmt.Printf("Time taken to compute the SHA256 hash of a 16-byte message :  %d nanoseconds\n", SHA256time)
	fmt.Printf("Time taken to encrypt a 1M-byte message (1024*1024 bytes) with AES-GCM using 256-bit AES key and no additional data : %d nanoseconds\n", Encrypttime)
	fmt.Printf("Time taken to decrypt the message above with AES-GCM using 256-bit AES key and no additional data : %d nanoseconds\n", Decrypttime)
}

func main() {
	validateSHA256()
	validateAESGCM()
	findPassword()
	performanceEvaluation()

}
