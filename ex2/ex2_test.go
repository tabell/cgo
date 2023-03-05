package ex2

import (
    "testing" 
    "bytes"
    "cpals/ex1"
)

func TestFixedXor(t *testing.T) {
    b1 := ex1.Hex2bytes("1c0111001f010100061a024b53535009181c")
    b2 := ex1.Hex2bytes("686974207468652062756c6c277320657965")


    if !bytes.Equal(Xor(b1,b2), ex1.Hex2bytes("746865206b696420646f6e277420706c6179")) {
        t.Errorf("Xor output didn't match")
    }
}
