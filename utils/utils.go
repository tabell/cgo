package utils

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
    "crypto/aes"
)

type ScoredText struct {
	Key   byte
	Text  string
	Score float64
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Xor(a []byte, b []byte) []byte {
	j := len(a)
	k := len(b)
	var out = make([]byte, max(j, k))
	for i, _ := range a {
		out[i] = a[i%j] ^ b[i%k]
	}

	return out
}

func makeStdFreq() (table [256]float64) {
	table[int(' ')] = 13
	table[int('A')] = 08.167
	table[int('B')] = 01.492
	table[int('C')] = 02.782
	table[int('D')] = 04.253
	table[int('E')] = 12.702
	table[int('F')] = 02.228
	table[int('G')] = 02.015
	table[int('H')] = 06.094
	table[int('I')] = 06.966
	table[int('J')] = 00.153
	table[int('K')] = 00.772
	table[int('L')] = 04.025
	table[int('M')] = 02.406
	table[int('N')] = 06.749
	table[int('O')] = 07.507
	table[int('P')] = 01.929
	table[int('Q')] = 00.095
	table[int('R')] = 05.987
	table[int('S')] = 06.327
	table[int('T')] = 09.056
	table[int('U')] = 02.758
	table[int('V')] = 00.978
	table[int('W')] = 02.360
	table[int('X')] = 00.150
	table[int('Y')] = 01.974
	table[int('Z')] = 00.074
	return table
}

func computeFrequencies(text []byte) (table [256]float64) {
	for i := range text {
		char := text[i]
		if char >= 97 && char <= 122 {
			char -= 32
		}
		table[int(char)]++
	}
	return
}

func testDistance() {
	in := []byte("The quick brown fox jumps over the lazy dog's back")
	fmt.Println(scoreText(in), string(in))
	in = []byte("Tomorrow, you will be released. If you are bored of brawling with thieves and want to achieve something there is a rare blue flower that grows on the eastern slopes. Pick one of these flowers. If you can carry it to the top of the mountain, you may find what you were looking for in the first place.")
	fmt.Println(scoreText(in), string(in))
	in = Hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(scoreText(in), string(in))
}

func distance(a [256]float64, b [256]float64) (distance float64) {
	for x := range a {
		d := math.Sqrt(math.Abs(a[x]*a[x] - b[x]*b[x]))
		distance += d
	}
	return
}

func FindBestKey(in []byte, cipher func([]byte, []byte)[]byte) ScoredText {
	/* todo: make this table static */
	keys := make([]rune, 256)
	for i := range keys {
		keys[i] = rune(i)
	}

	/* decode the string with each key and score the result */
	scores := make([]ScoredText, len(keys))
	for i := range keys {
		k := make([]byte, 1)
		scores[i].Key = byte(i)
		k[0] = scores[i].Key
		scores[i].Text = string(cipher(in, k))
		scores[i].Score = scoreText([]byte(scores[i].Text))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score < scores[j].Score
	})

	return scores[0]
}

func scoreText(in []byte) float64 {
	return math.Abs(1 - distance(makeStdFreq(), computeFrequencies(in)))
}

func StringArrayFromFile(filename string) (result []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func BytesFromBase64File(filename string) (raw []byte) {
	/* load base64 from file */
	encodedStrings := StringArrayFromFile(filename)
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

func FindKeysize(in []byte) int {
	results := make([]editMap, 40)
	for keysize := 2; keysize < 40; keysize++ {
		results[keysize] = selfCorrelate(in, keysize)
	}
	results = results[2:] /* hack because there needs to be a 0th elem */
	sort.Slice(results, func(i, j int) bool { return results[i].score < results[j].score })
	return results[0].length
}

func FindKey(in []byte, keysize int, cipher func([]byte, []byte)[]byte) (key []byte) {
	blockCount := len(in) / keysize
	remainder := len(in) % keysize
	blockCount -= remainder
	key = make([]byte, keysize)
	for i := 0; i < keysize; i++ {
		bucket := make([]byte, blockCount)
		for j := 0; j < blockCount; j++ {
			bucket[j] = in[j*keysize+i]
		}

		bestKey := FindBestKey(bucket, cipher)
		key[i] = byte(bestKey.Key)
	}

	return
}

func AesEcb(ct []byte, key []byte)(pt []byte) {
    if len(key) % 16 != 0 {
        log.Fatal("Key size", len(key), "invalid for AES")
    }
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Error creating aes block cipher:", err)
	}

	blocksize := len(key)
    blocks := len(ct) / len(key)
	pt = make([]byte, len(ct))

	for x := 0; x < blocks; x++ {
		cipher.Decrypt(pt[x*blocksize:(x+1)*blocksize], ct[x*blocksize:(x+1)*blocksize])
	}

    return
}
