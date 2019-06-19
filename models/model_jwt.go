package models

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//F2Sha256 ..
func F2Sha256(param1 string, param2 string) string {
	h := sha256.New()
	h.Write([]byte(param1 + param2))
	encryptKey := fmt.Sprintf("%x", h.Sum(nil))
	return encryptKey
}

//GetToken ..
func GetTokenJwt(key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	var mySigningKey = []byte(key)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	/* Sign the token with our secret */
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

//ValidToken ..
func ValidTokenJwt(r *http.Request, tokenAuth string) (bool, error) {
	var mySigningKey = []byte(tokenAuth)
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

	if err != nil {
		//fmt.Println("error 1", err)
		return false, err
	}

	if !token.Valid {
		//fmt.Println("error 2", tokenAuth)
		return false, errors.New("Invalid Token")
	}

	return true, nil
}
