package str_test

import (
	"testing"

	"github.com/sendya/pkg/rand/str"
)

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := str.New(6)
		t.Log(s)
	}
}
