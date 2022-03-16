package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Sum(v string) []byte {
	m := md5.New()
	m.Write([]byte(v))
	return m.Sum(nil)
}

func Sums(v string) string {
	return hex.EncodeToString(Sum(v))
}
