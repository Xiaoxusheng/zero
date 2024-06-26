package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"

	"time"
)

type MyCustomClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

func GetToken(uid string, key string, times int32) string {
	var mySigningKey = []byte(key)
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		uid,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(times) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)
	return ss
}
