package tools

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Sha8(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs[:4])
}

func Skey(name string, num int) string {
	return fmt.Sprintf("%s_%d", Sha8(name), num)
}

const (
	D0 string = "D0"
	// origin
	D1 string = "D1"
	D2 string = "D2"
	D3 string = "D3"
)
