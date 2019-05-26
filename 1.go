package main

import (
    "fmt"
    "encoding/base64"
    "encoding/hex"
)

func main() {
    ex1()
}

func ex1() {
    in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    fmt.Println(hex2b64([]byte(in)))
}

func hex2b64(str []byte) string {
    raw := make([]byte, hex.DecodedLen(len(str)))
    hex.Decode(raw, str)
    encoded := base64.StdEncoding.EncodeToString(raw)
    return encoded
}
