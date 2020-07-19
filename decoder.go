package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
)

func DecodeJwt(tokenString string) (string, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	headers, err := json.MarshalIndent(token.Header, "", "\t")
	if err != nil {
		return "", err
	}
	payload, err := json.MarshalIndent(token.Claims, "", "\t")
	if err != nil {
		return "", err
	}
	return string(headers) + "\n" + string(payload), err
}
