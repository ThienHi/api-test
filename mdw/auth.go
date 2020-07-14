package mdw

import (
	"go-rest-api/models"

	"github.com/labstack/echo/v4"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error) {

	req := new(models.User)
	c.Bind(req)

	if req.Username == "admin" && req.Password == "12345" {
		return true, nil
	}

	return false, nil
}
