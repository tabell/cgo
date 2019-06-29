package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
)

func main() {
    in := utils.Hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    result := utils.FindBestKey(in, utils.Xor)
    fmt.Printf("key=%s score = %.3f msg=%s\n",
        string(result.Key), result.Score, result.Text)
}
