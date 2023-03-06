package ex3

import (
	"math"
	"sort"
)

type ScoredText struct {
	Key   byte
	Text  string
	Score float64
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

func distance(a [256]float64, b [256]float64) (distance float64) {
	for x := range a {
		d := math.Sqrt(math.Abs(a[x]*a[x] - b[x]*b[x]))
		distance += d
	}
	return
}

func scoreText(in []byte) float64 {
	return math.Abs(1 - distance(makeStdFreq(), computeFrequencies(in)))
}

