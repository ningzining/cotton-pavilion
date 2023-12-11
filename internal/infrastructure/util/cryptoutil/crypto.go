package cryptoutil

import (
	"crypto/md5"
	"fmt"
)

func Md5(source string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(source)))
}

func Md5Password(mobile string, password string) string {
	return Md5(fmt.Sprintf("%s%s", mobile, password))
}
