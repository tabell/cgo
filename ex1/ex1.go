package ex1

import (
	"encoding/base64"
	"encoding/hex"
)

func Hex2b64(hex string) string {
    return string(BytesToBase64(Hex2bytes(hex)))
}

func BytesToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}


func Hex2bytes(ins string) []byte {
	in := []byte(ins)
	raw := make([]byte, hex.DecodedLen(len(in)))
	hex.Decode(raw, in)
	return raw
}
