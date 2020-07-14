package router

import (
	api "go-rest-api/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init will init a echo router
func Init() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())

	isLogedIn := middleware.JWT([]byte("secret"))
	// isAdmin := mdw.IsAdminMdw

	e.GET("/users", api.GetUsers, isLogedIn)

	e.GET("/user/:id", api.GetUser)

	e.POST("/login", api.Login)

	e.POST("/user", api.AddUser)
	e.PUT("/user/:id", api.UpdateUser)
	e.DELETE("/user/:id", api.DeleteUser)

	return e
}
