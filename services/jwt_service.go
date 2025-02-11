package services

import (
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
)

type JsonWebTokenService struct{}

func (_ JsonWebTokenService) Parse(token string) (*jwt.RegisteredClaims, error) {
	key, _ := ioutil.ReadFile("private.key")
	claims := &jwt.RegisteredClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
