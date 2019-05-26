package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
    "bufio"
    "os"
    "log"
    "sort"
)

func main() {
    keys := make([]rune, 256)
    for i := range keys {
        keys[i] = rune(i)
    }

    file,err := os.Open("ex4/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var testData []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        testData = append(testData, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    results := make([]utils.ScoredText, len(testData))

    for i,str := range testData {
        in := utils.Hex2bytes(str)
        result := utils.BestScore(in, keys)
        results[i] = result
    }

    sort.Slice(results, func(i, j int) bool {
        return results[i].Score < results[j].Score
    })

    result := results[0]
    fmt.Printf("key=%s score=%.3f msg=%s\n", string(result.Key), result.Score, result.Text)

}
