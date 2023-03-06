package cpals

import (
    "testing"
    "cpals/ex2"
    "strings"
)

func TestEx6(t *testing.T) {
	/* load base64 from file */
    raw, err := BytesFromBase64File("input.txt")
    if err != nil {
        t.Errorf("%s", err)
    }

	/* guess at key size */
	keysize := FindKeysize(raw)

	key := FindKey(raw, keysize, ex2.Xor)
    if string(key) != "Terminator X: Bring the noise" {
        t.Errorf("Error finding the key")
    }

    output := string(ex2.Xor(raw, key))
    if !strings.Contains(output, "Play that funky music white boy") {
        t.Errorf("Key was correct but decoding failed")
    }
}

func TestHamm(t *testing.T) {
	a := "this is a test"
	b := "wokka wokka!!!"
    if hamm([]byte(a), []byte(b)) != 37 {
        t.Errorf("Error calculating hamming distance")
    }
}

