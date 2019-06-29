package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
    "sort"
)

func main() {

    inStrings := utils.StringArrayFromFile("ex4/input.txt")

    results := make([]utils.ScoredText, len(inStrings))

    for i,str := range inStrings {
        in := utils.Hex2bytes(str)
        result := utils.FindRepeatingXorKey(in)
        results[i] = result
    }

    sort.Slice(results, func(i, j int) bool {
        return results[i].Score < results[j].Score
    })
    winningResult := results[0]

    fmt.Printf("key=%c score=%.3f msg=%s\n", rune(winningResult.Key), winningResult.Score, winningResult.Text)
}
