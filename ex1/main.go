package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
)

func main() {
    in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    fmt.Println(string (utils.Hex2b64(in)))
}
