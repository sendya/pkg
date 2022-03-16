package str

import (
	"math/rand"
	"time"
)

// init
var (
	seededRand  = rand.New(rand.NewSource(time.Now().UnixNano()))
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	timeLayout  = "20060102150405"
)

// New ... string randon custom len.
func New(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[seededRand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewTime() string {
	t := time.Now()
	return t.Format(timeLayout) + string(rune(t.Nanosecond()/int(time.Millisecond)))
}
