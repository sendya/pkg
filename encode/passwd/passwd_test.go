package passwd_test

import (
	"strings"
	"testing"

	"github.com/sendya/pkg/encode/passwd"
)

func TestGenPassword(t *testing.T) {
	pwd, salt := passwd.Gen("testpwd")
	t.Logf("pwd: %s, salt: %s", pwd, salt)
}

func TestCheckPassword(t *testing.T) {
	encoded := passwd.Sum("123456", "AfoYsf")
	t.Logf("pwd: %s", encoded)

	encoded2 := passwd.Sum("123456", "AfoYsf")
	t.Logf("pwd2: %s", encoded2)

	encoded3 := passwd.Sum("123456", "uQhXOe")
	t.Logf("pwd3: %s", encoded3)

	// 一定相等
	if !strings.Contains(encoded, encoded2) {
		t.FailNow()
	}

	// 一定不等
	if strings.Contains(encoded, encoded3) {
		t.FailNow()
	}
}
