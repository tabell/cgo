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

func FindRepeatingXorKey(in []byte) ScoredText {
	/* todo: make this table static */
	keys := make([]rune, 256)
	for i := range keys {
		keys[i] = rune(i)
	}

	/* decode the string with each key and score the result */
	scores := make([]ScoredText, len(keys))
	for i := range keys {
		//scores[i].Text = string(Xor(in, []byte(string(scores[i].Key))))
		k := make([]byte, 1)
		scores[i].Key = byte(i)
		k[0] = scores[i].Key
		scores[i].Text = string(Xor(in, k))
		scores[i].Score = scoreText([]byte(scores[i].Text))
		//(%c) --> score = %f: key=%c decoded=%s",
		//	i, i, scores[i].Score, scores[i].Key, scores[i].Text)
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score < scores[j].Score
	})

	//log.Printf(scores[0].Text)

	return scores[0]
}

func Hex2bytes(ins string) []byte {
	in := []byte(ins)
	raw := make([]byte, hex.DecodedLen(len(in)))
	hex.Decode(raw, in)
	return raw
}

func Hex2b64(str string) []byte {
	raw := Hex2bytes(str)
	encoded := base64.StdEncoding.EncodeToString(raw)
	return []byte(encoded)
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
