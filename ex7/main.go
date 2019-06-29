package main

import (
	"github.com/tabell/cpals/utils"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ct := utils.BytesFromBase64File("ex7/ct.b64")
	key := []byte("YELLOW SUBMARINE")

	log.Print("Decrypted: ", string(utils.AesEcb(ct, key)))

}
