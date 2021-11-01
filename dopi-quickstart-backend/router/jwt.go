package router

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type DopiClaims struct {
	Username string `json:"username"`
	Roles    []string
	jwt.StandardClaims
}

func (claims *DopiClaims) hasRole(role string) bool {
	for _, r := range claims.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &DopiClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetClaims(token *jwt.Token) (*DopiClaims, error) {
	if claims, ok := token.Claims.(*DopiClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Token invalid")

}
