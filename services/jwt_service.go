package services

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type JsonWebTokenService struct{}

func (_ JsonWebTokenService) Parse(token string) (*jwt.RegisteredClaims, error) {
	key, err := os.ReadFile("private.key")

	if err != nil {
		return nil, err
	}

	claims := &jwt.RegisteredClaims{}

	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
