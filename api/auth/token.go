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
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return t, err
}

func CreateRefreshToken(username string, admin bool) (string, error) {
	var err error

	refreshtoken := jwt.New(jwt.SigningMethodHS256)

	rfclaims := refreshtoken.Claims.(jwt.MapClaims)
	rfclaims["username"] = username
	rfclaims["admin"] = true
	rfclaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshtoken.SignedString([]byte("mysecret"))

	if err != nil {
		return "", err
	}

	return rt, err
}
