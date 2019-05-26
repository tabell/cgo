package main

import (
    "fmt"
    "encoding/base64"
    "encoding/hex"
    "unicode"
    "math"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    ex1()
    ex2()
    ex3()
}

func ex1() {
    in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    fmt.Println(string (hex2b64(in)))
}

func ex2() {
    in1 := hex2bytes("1c0111001f010100061a024b53535009181c")
    in2 := hex2bytes("686974207468652062756c6c277320657965")
    var result = make([]byte, max(len(in1), len(in2)))
    for i,_ := range in1 {
        result[i] = in1[i] ^ in2[i]
    }
    fmt.Println(hex.EncodeToString(result))

}

func makeStdFreq() (table [256]float64) {
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

func distance(a [256]float64, b [256]float64) (distance float64) {
    for x := range a {
        distance += math.Sqrt(math.Abs(a[x]*a[x] - b[x]*b[x]))
    }
    return
}


func ex3() {
    //in := hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    in := []byte("The quick brown fox jumps over the lazy dog's back")
    fmt.Println(score(in), string(in))
    in = []byte("Tomorrow, you will be released. If you are bored of brawling with thieves and want to achieve something there is a rare blue flower that grows on the eastern slopes. Pick one of these flowers. If you can carry it to the top of the mountain, you may find what you were looking for in the first place.")
    fmt.Println(score(in), string(in))
    in = hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    fmt.Println(score(in), string(in))
}

func hex2bytes(ins string) []byte {
    in := []byte(ins)
    raw := make([]byte, hex.DecodedLen(len(in)))
    hex.Decode(raw, in)
    return raw
}

func hex2b64(str string) []byte {
    raw := hex2bytes(str)
    encoded := base64.StdEncoding.EncodeToString(raw)
    return []byte (encoded)
}

func score(in []byte) float64 {
    return math.Log(distance(makeStdFreq(), computeFrequencies(in)))
}
