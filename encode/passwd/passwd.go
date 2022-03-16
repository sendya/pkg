package passwd

import (
	"fmt"

	"github.com/sendya/pkg/encode/md5"
	"github.com/sendya/pkg/rand/str"
)

const PASSWORD_KEY_PREFIX = "$$"
const PASSWORD_SALT_SIZE = 12

// Sum password by salt ... return encrypted password
func Sum(password, salt string) string {
	return fmt.Sprintf("%s%s", PASSWORD_KEY_PREFIX, md5.Sums(fmt.Sprintf("%s:%s", salt, password)))
}

// Gen password ... return encrypted password and salt key
func Gen(password string) (string, string) {
	salt := str.New(PASSWORD_SALT_SIZE)
	return Sum(password, salt), salt
}
