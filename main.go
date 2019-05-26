package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
    "encoding/hex"
)

func main() {
    ex1()
    ex2()
    ex3()
}

func ex1() {
    in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    fmt.Println(string (utils.Hex2b64(in)))
}

func ex2() {
    in1 := utils.Hex2bytes("1c0111001f010100061a024b53535009181c")
    in2 := utils.Hex2bytes("686974207468652062756c6c277320657965")
    fmt.Println(hex.EncodeToString(utils.Xor(in1, in2)))
}

func ex3() {
    in := utils.Hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    keys := make([]rune, 256)
    for i := range keys {
        keys[i] = rune(i)
    }
    fmt.Println(utils.BestScore(in, keys))
}
