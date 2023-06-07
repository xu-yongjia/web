//Package util ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 14:37:47
 * @LastEditors: congz
 * @LastEditTime: 2020-08-09 19:28:00
 */
package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("CgtUWk4VVIMvbO3MdRbo7kR1uDTlHbHneh1xn7veqfEsbUZnaCcCixokmgDd3f2")

// Claims ...
type Claims struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(username, password string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2400 * time.Hour)

	claims := Claims{
		username,
		password,
		authority,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "canteenpre",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
