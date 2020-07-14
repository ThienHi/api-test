package api

import (
	"go-rest-api/api/auth"
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db := db.GetDB()

	users := []models.User{}
	user := models.User{}

	c.Bind(&user)
	db.Find(&users)

	var err error
	var token string

	for _, v := range users {

		if v.Username != user.Username {
			if v.Username != "admin" {
				return c.String(http.StatusUnauthorized, "Wrong Username")
			}
		}

		if v.Password != user.Password {
			if v.Password != "12345" {
				return c.String(http.StatusUnauthorized, "Wrong PassWord")
			}
		}

		if v.Username == user.Username && v.Password == user.Password {
			if v.Username == "admin" && v.Password == "12345" {
				token, err = auth.CreateToken(user.Username, true)
				if err != nil {
					return err
				}
				return c.JSON(http.StatusOK, map[string]string{
					"token": token, "Admin": "Admin",
				})
			}

			token, err := auth.CreateToken(user.Username, false)

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, map[string]string{
				"token": token, "Admin": "Not Admin",
			})
		}
	}
	return c.String(http.StatusUnauthorized, "nil")
}
