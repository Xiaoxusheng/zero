package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Salt 随即盐值
func Salt(n int32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return b, nil
}

// HashPassword 加密密码
func HashPassword(password string, b []byte) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write(b)
	hashPassword := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(hashPassword)
}
