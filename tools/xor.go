package main

import (
    "os"
    "fmt"
    "cpals/ex2"
    "strings"
)

func main() {
    key := os.Args[1]
    plaintext := os.Args[2:]
    fmt.Println(string(ex2.Xor([]byte(strings.Join(plaintext, " ")),[]byte(key))))
}
