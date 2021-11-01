package auth

import (
	"os"

	"aphro.web/model"
	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(os.Getenv("SIGNING_KEY"))

func createJWTToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user

	return token.SignedString(signingKey)
}

func ParseJWTToken(tokenString string) (interface{}, error) {

	callback := func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	}

	token, err := jwt.Parse(tokenString, callback)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		claims = nil
	}

	return claims, err
}
