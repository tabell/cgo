package main

import (
    "fmt"
    "encoding/base64"
    "encoding/hex"
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

func computeFrequencies(in []byte) (table [256]rune) {
    for i := range in {
        table[int(in[i])]++
    }
    return
}

func ex3() {
    in := hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    fmt.Println("score ", score(in))
    fmt.Println("freq ", computeFrequencies(in))
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

func score(in []byte) int {
    return 0
}
