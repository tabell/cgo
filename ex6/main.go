package main

import (
    "fmt"
    "github.com/tabell/cpals/utils"
    "encoding/base64"
    "strings"
    "sort"
)

func bitCount(in byte) int {
    count := 0
    for x := 0; x < 8; x++ {
        if ((in >> uint8(x)) & 0x1) == 1 {
            count++
        }
    }
    return count
}

func hamm(a []byte, b []byte) int {
    count := 0
    for x := range a {
        c := a[x]^ b[x]
        count += bitCount(c)
    }
    return count
}

func qualify() {
    a := "this is a test"
    b := "wokka wokka!!!"

    fmt.Println(a,b, hamm([]byte(a),[]byte(b)))
}

type editMap struct {
    length int
    score float64
}

func selfCorrelate(in []byte, size int) (result editMap) {
    result.score = float64(hamm(in[:size], in[size:2*size])) / float64(size)
    result.length = size
    return
}

func findKeysize(in []byte) int {
    results := make([]editMap, 40)
    for keysize := 1; keysize < 40; keysize++ {
        results[keysize] = selfCorrelate(in, keysize)
    }
    results = results[1:] /* hack because there needs to be a 0th elem */
    sort.Slice(results, func (i, j int) bool { return results[i].score < results[j].score })
    return results[0].length
}

func foldBytes(in []byte, keysize int) (blocks []byte) {
    depth := len(in) / keysize
    blocks = make([]byte, keysize)
    fmt.Printf("total=%d keysize=%d depth=%d\n", len(in), keysize, depth)
    for i := 0; i < keysize; i++ {
        slot := make([]byte, depth)
        for j := 0; j < depth; j++ {
            slot[j] = in[j*keysize+i]
        }

        bestKey := utils.FindRepeatingXorKey(slot)
        blocks[i] = byte(bestKey.Key)
        fmt.Println(bestKey.Score)
    }

    return
}

func main() {
    /* load base64 from file */
    encodedStrings := utils.StringArrayFromFile("ex6/input.txt")
    var bigStr strings.Builder

    /* build into one large base64 buffer */
    for _,s := range encodedStrings {
        bigStr.WriteString(s)
    }

    /* convert base64 to raw bytes */
    raw := make([]byte, base64.StdEncoding.DecodedLen(len(bigStr.String())))
    data,_ := base64.StdEncoding.DecodeString(bigStr.String())
    for i,b := range data {
        raw[i] = b
    }

    /* guess at key size */
    //keysize := findKeysize(raw)
    keysize := 29
        fmt.Println("found key size: ", keysize)

        key := foldBytes(raw, keysize)
        fmt.Println(string(key))

        fmt.Println(string(utils.Xor(raw,key)))

}

