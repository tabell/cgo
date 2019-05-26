package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
    "sort"
)

func main() {
    keys := make([]rune, 256)
    for i := range keys {
        keys[i] = rune(i)
    }

    inStrings := utils.StringArrayFromFile("ex4/input.txt")

    results := make([]utils.ScoredText, len(inStrings))

    for i,str := range inStrings {
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
