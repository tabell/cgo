package main

import (
	"fmt"
	"github.com/tabell/cpals/utils"
	"log"
    "sort"
)


func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

    inStrings := utils.StringArrayFromFile("ex8/8.txt")

    results := make([]utils.ScoredText, len(inStrings))

    for i,str := range inStrings {
        in := utils.Hex2bytes(str)
        result := utils.FindBestKey(in, utils.AesEcb)
        results[i] = result
    }

    sort.Slice(results, func(i, j int) bool {
        return results[i].Score < results[j].Score
    })
    winningResult := results[0]

    fmt.Printf("key=%c score=%.3f msg=%s\n", winningResult.Key, winningResult.Score, winningResult.Text)

}
