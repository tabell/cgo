package main

import (
    "github.com/tabell/cpals/utils"
    "fmt"
    "encoding/hex"
)

func main() {
    in := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
    fmt.Println(hex.EncodeToString(utils.Xor(in, []byte("ICE"))))
}

