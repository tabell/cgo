package main

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/tabell/cpals/utils"
	"log"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Hello world")
	ct := bytesFromBase64File("ex7/ct.b64")
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

func bytesFromBase64File(filename string) (raw []byte) {
	/* load base64 from file */
	encodedStrings := utils.StringArrayFromFile(filename)
	var bigStr strings.Builder

	/* build into one large base64 buffer */
	for _, s := range encodedStrings {
		bigStr.WriteString(s)
	}

	/* allocate ..? */
	raw = make([]byte, base64.StdEncoding.DecodedLen(len(bigStr.String())))

	/* convert base64 to raw bytes */
	data, err := base64.StdEncoding.DecodeString(bigStr.String())
	if err != nil {
		log.Fatal("Couldn't decode base64:", err)
	}
	for i, b := range data {
		raw[i] = b
	}
	return
}
