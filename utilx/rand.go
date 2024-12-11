package utilx

import (
	"crypto/rand"
	"math/big"
)

const (
	charsetStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// RandStr 生成指定长度的随机字符串
func RandStr(length int) (string, error) {
	charsetLength := big.NewInt(int64(len(charsetStr)))

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		randomString[i] = charsetStr[randomIndex.Int64()]
	}
	return string(randomString), nil
}
