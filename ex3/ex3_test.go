package ex3

import (
    "testing"
    "cpals/ex1"
    "cpals/ex2"
)

func TestFindBestKey(t *testing.T) {
    in := ex1.Hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    result := FindBestKey(in, ex2.Xor)
    if result.Key != 'X' {
        t.Errorf("Key was wrong")
    }
    if result.Text != "Cooking MC's like a pound of bacon" {
        t.Errorf("Key was right but decoding failed")
    }
}

func TestScoreText(t *testing.T) {

	in := []byte("The quick brown fox jumps over the lazy dog's back")
    if int(scoreText(in)) != 106 {
        t.Errorf("Score calculated incorrectly for %s", string(in))
    }
	in = []byte("Tomorrow, you will be released. If you are bored of brawling with thieves and want to achieve something there is a rare blue flower that grows on the eastern slopes. Pick one of these flowers. If you can carry it to the top of the mountain, you may find what you were looking for in the first place.")
    if int(scoreText(in)) != 272 {
        t.Errorf("Score calculated incorrectly for %s", string(in))
    }
	in = ex1.Hex2bytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    if int(scoreText(in)) != 145 {
        t.Errorf("Score calculated incorrectly for %s", string(in))
    }
}

