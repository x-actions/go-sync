package utils

import (
	"crypto/md5"
	"fmt"
)


func Md5sum(str string) string {
	data := []byte(str)
	hash :=  md5.Sum(data)
	return fmt.Sprintf("%X", hash)
}
