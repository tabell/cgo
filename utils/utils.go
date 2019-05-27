package utils

import (
    "fmt"
    "encoding/base64"
    "encoding/hex"
    "unicode"
    "math"
    "sort"
    "os"
    "log"
    "bufio"
)

type ScoredText struct {
    Key rune
    Text string
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
    var out = make([]byte, max(j,k))
    for i,_ := range a {
        out[i] = a[i % j] ^ b[i % k]
    }

    return out
}

func makeStdFreq() (table [256]float64) {
    table[int(' ')] =  0.05
    table[int('A')] =  0.08167
    table[int('B')] =  0.01492
    table[int('C')] =  0.02782
    table[int('D')] =  0.04253
    table[int('E')] =  0.12702
    table[int('F')] =  0.02228
    table[int('G')] =  0.02015
    table[int('H')] =  0.06094
    table[int('I')] =  0.06966
    table[int('J')] =  0.00153
    table[int('K')] =  0.00772
    table[int('L')] =  0.04025
    table[int('M')] =  0.02406
    table[int('N')] =  0.06749
    table[int('O')] =  0.07507
    table[int('P')] =  0.01929
    table[int('Q')] =  0.00095
    table[int('R')] =  0.05987
    table[int('S')] =  0.06327
    table[int('T')] =  0.09056
    table[int('U')] =  0.02758
    table[int('V')] =  0.00978
    table[int('W')] =  0.02360
    table[int('X')] =  0.00150
    table[int('Y')] =  0.01974
    table[int('Z')] =  0.00074
    return table
}


func computeFrequencies(text []byte) (table [256]float64) {
    for i := range text {
        text[i] = byte(unicode.ToUpper(rune(text[i])))
    }
    count := 0
    for i := range text {
        count++
        table[int(text[i])]++
    }
    for i := range text {
        table[int(text[i])] /= float64(count)
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
        distance += math.Sqrt(math.Abs(a[x]*a[x] - b[x]*b[x]))
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
        scores[i].Key = keys[i]
        scores[i].Text = string(Xor(in, []byte(string(scores[i].Key))))
        scores[i].Score = scoreText([]byte(scores[i].Text))
    }
    sort.Slice(scores, func(i, j int) bool {
        return scores[i].Score < scores[j].Score
    })

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
    return []byte (encoded)
}

func scoreText(in []byte) float64 {
    return math.Abs(1 - distance(makeStdFreq(), computeFrequencies(in)))
}

func StringArrayFromFile(filename string) (result []string) {
    file,err := os.Open(filename)
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

