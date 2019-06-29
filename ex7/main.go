package main

import (
	"crypto/aes"
	"github.com/tabell/cpals/utils"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Hello world")
	ct := utils.BytesFromBase64File("ex7/ct.b64")
	key := []byte("YELLOW SUBMARINE")

	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Error creating aes block cipher:", err)
	}

	plain := make([]byte, len(ct))

	blocks := len(ct) / len(key)
    log.Printf("Ciphertext length=%d key length=%d blocks=%d", len(ct), len(key), blocks)

	blocksize := len(key)
	for x := 0; x < blocks; x++ {
		cipher.Decrypt(plain[x*blocksize:(x+1)*blocksize], ct[x*blocksize:(x+1)*blocksize])
	}

	log.Print("Decrypted: ", string(plain))

}
