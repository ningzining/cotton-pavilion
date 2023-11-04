package crypto

import (
	"crypto/md5"
	"fmt"
)

func Md5(source string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(source)))
}
