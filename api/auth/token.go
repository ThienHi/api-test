package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string, admin bool) (string, error) {
	var err error

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return t, nil
}
