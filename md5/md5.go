package md5

import (
	"crypto/md5"
	"fmt"
)

// GetMd5 is helper func for get md5 hash
func GetMd5(in []byte) string {
	return fmt.Sprintf("%x", md5.Sum(in))
}
