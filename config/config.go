package config

import (
	"crypto/md5"
	"encoding/hex"
)

const PAGINATION_DEFAULT_SIZE uint = 10

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
