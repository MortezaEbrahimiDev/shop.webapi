package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("secret")

type Claims struct {
	UserId uint
	Roles  []string
	jwt.RegisteredClaims
}

func GenerateJWT(UserId uint, Roles []string) (string, error) {
	claims := Claims{
		UserId: UserId,
		Roles:  Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
