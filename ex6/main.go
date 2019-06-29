package main

import (
	"fmt"
	"github.com/tabell/cpals/utils"
	"log"
)

func findKey(in []byte, keysize int) (key []byte) {
	blockCount := len(in) / keysize
	remainder := len(in) % keysize
	blockCount -= remainder
	key = make([]byte, keysize)
	for i := 0; i < keysize; i++ {
		bucket := make([]byte, blockCount)
		for j := 0; j < blockCount; j++ {
			bucket[j] = in[j*keysize+i]
		}

		bestKey := utils.FindBestKey(bucket, utils.Xor)
		key[i] = byte(bestKey.Key)
	}

	return
}

func main() {
	/* load base64 from file */
	log.SetFlags(log.LstdFlags | log.Lshortfile)

    raw := utils.BytesFromBase64File("ex6/input.txt")

	/* guess at key size */
	keysize := utils.FindKeysize(raw)
	log.Println("found key size: ", keysize)

	key := findKey(raw, keysize)
	log.Println("Key is: ", string(key))

	fmt.Println(string(utils.Xor(raw, key)))
}
