package main

import (
	"fmt"
	"github.com/tabell/cpals/utils"
	"log"
)

func main() {
	/* load base64 from file */
	log.SetFlags(log.LstdFlags | log.Lshortfile)

    raw := utils.BytesFromBase64File("ex6/input.txt")

	/* guess at key size */
	keysize := utils.FindKeysize(raw)
	log.Println("found key size: ", keysize)

	key := utils.FindKey(raw, keysize, utils.Xor)
	log.Println("Key is: ", string(key))

	fmt.Println(string(utils.Xor(raw, key)))
}
