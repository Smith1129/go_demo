package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	TokenFail error = errors.New("token is unuseable")
	SignKey string = "zzz_godemo"
)

type UserClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

 func CreateToken(claims UserClaims) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(SignKey)
	return token.SignedString(key)
 }
