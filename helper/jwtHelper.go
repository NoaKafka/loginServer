package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(Email string) (string, error) {
	mySigningKey := []byte("noakafka")

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
