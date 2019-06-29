package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tabell/cpals/utils"
	"log"
	"sort"
	"strings"
)

func bitCount(in byte) int {
	count := 0
	for x := 0; x < 8; x++ {
		if ((in >> uint8(x)) & 0x1) == 1 {
			count++
		}
	}
	return count
}

func hamm(a []byte, b []byte) int {
	count := 0
	for x := range a {
		c := a[x] ^ b[x]
		count += bitCount(c)
	}
	return count
}

func qualify() {
	a := "this is a test"
	b := "wokka wokka!!!"

	log.Println(a, b, hamm([]byte(a), []byte(b)))
}

type editMap struct {
	length int
	score  float64
}

func selfCorrelate(in []byte, size int) (result editMap) {
	blocks := len(in) / size
	score := 0.0
	for b := 0; b < blocks-1; b++ {
		s1 := size * b
		s2 := size * (b + 1)
		s3 := size * (b + 2)
		score += float64(hamm(in[s1:s2], in[s2:s3])) / float64(size)
	}
	result.score = score / float64(blocks)
	result.length = size
	return
}

func findKeysize(in []byte) int {
	results := make([]editMap, 40)
	for keysize := 2; keysize < 40; keysize++ {
		results[keysize] = selfCorrelate(in, keysize)
	}
	results = results[2:] /* hack because there needs to be a 0th elem */
	sort.Slice(results, func(i, j int) bool { return results[i].score < results[j].score })
	return results[0].length
}

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

		bestKey := utils.FindRepeatingXorKey(bucket)
		key[i] = byte(bestKey.Key)
	}

	return
}

func main() {
	/* load base64 from file */
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	encodedStrings := utils.StringArrayFromFile("ex6/input.txt")
	var bigStr strings.Builder

	/* build into one large base64 buffer */
	for _, s := range encodedStrings {
		bigStr.WriteString(s)
	}

	/* allocate ..? */
	raw := make([]byte, base64.StdEncoding.DecodedLen(len(bigStr.String())))

	/* convert base64 to raw bytes */
	data, err := base64.StdEncoding.DecodeString(bigStr.String())
	if err != nil {
		log.Fatal("Couldn't decode base64:", err)
	}
	for i, b := range data {
		raw[i] = b
	}

	/* guess at key size */
	keysize := findKeysize(raw)
	log.Println("found key size: ", keysize)

	key := findKey(raw, keysize)
	log.Println("Key is: ", string(key))

	fmt.Println(string(utils.Xor(raw, key)))
}
