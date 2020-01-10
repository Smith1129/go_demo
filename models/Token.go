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

 func ParseToken(tokenString string) (*UserClaims,error){
	 key := []byte(SignKey)
	 token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		 return key, nil
	 })
	 //log.Fatalln(err,"err")
	 if err != nil {
		 return nil,TokenFail
	 }
	 //fmt.Printf("%s",token)
	 //fmt.Println(reflect.TypeOf(token))
	 //a,err := json.Marshal(token)
	 //fmt.Println(string(a))
	 //fmt.Printf("%s",token)
	 if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		 return claims, nil
	 }
	 return nil, TokenFail
 }

 //func RefreshToken(tokenString string) (string,error){
	// key := []byte(SignKey)
	// jwt.TimeFunc = func() time.Time {
	//	 return time.Unix(0, 0)
	// }
	// token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	 return key, nil
	// })
	// if err != nil {
	//	 return "", err
	// }
	// if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
	//	 jwt.TimeFunc = time.Now
	//	 claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
	//	 return CreateToken(*claims)
	// }
	// return "", TokenFail
 //}
