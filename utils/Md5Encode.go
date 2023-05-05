package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

// Md5
// 转小写
func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func GetSalt() string {
	min := 10000
	max := 99999
	randNum := rand.Intn(max-min+1) + min
	salt := Md5(string(randNum))[:5]
	return salt
}
