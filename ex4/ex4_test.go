package cpals

import (
    "sort"
    "testing"
    "cpals/ex1"
    "cpals/ex2"
    "cpals/ex3"
)

func TestFindBestKey2(t *testing.T) {

    err, inStrings := StringArrayFromFile("input.txt")
    if err != nil {
        t.Errorf("Error loading string array from file input.txt")
    }

    results := make([]ex3.ScoredText, len(inStrings))

    for i,str := range inStrings {
        in := ex1.Hex2bytes(str)
        result := ex3.FindBestKey(in, ex2.Xor)
        results[i] = result
    }

    sort.Slice(results, func(i, j int) bool {
        return results[i].Score < results[j].Score
    })
    winningResult := results[0]

    if winningResult.Key != '5' {
        t.Errorf("Key was wrong")
    }
    if winningResult.Text != "Now that the party is jumping\n" {
        t.Errorf("Key was right but decoding failed")
    }
}
