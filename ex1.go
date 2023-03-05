package ex1

import (
	"encoding/base64"
	"encoding/hex"
)

func Hex2b64(str string) []byte {
	raw := Hex2bytes(str)
	encoded := base64.StdEncoding.EncodeToString(raw)
	return []byte(encoded)
}

func Hex2bytes(ins string) []byte {
	in := []byte(ins)
	raw := make([]byte, hex.DecodedLen(len(in)))
	hex.Decode(raw, in)
	return raw
}
