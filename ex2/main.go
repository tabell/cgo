package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
    "encoding/hex"
)

func main() {
    in1 := utils.Hex2bytes("1c0111001f010100061a024b53535009181c")
    in2 := utils.Hex2bytes("686974207468652062756c6c277320657965")
    fmt.Println(hex.EncodeToString(utils.Xor(in1, in2)))
}
