package ex1

import "testing" 

func TestHex2bytes(t *testing.T) {
    in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    out := Hex2b64(in)
    if out != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
        t.Errorf("Base64 output was not correct")
    }
}
