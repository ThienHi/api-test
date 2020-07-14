package api

import (
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUsers will get a user from DB
func GetUsers(c echo.Context) error {
	db := db.GetDB()

	users := []models.User{}
	db.Find(&users) // select * from user
	// spew.Dump(json.Marshal(users))
	return c.JSON(http.StatusOK, users)
}

// AddUser will add a user to DB
func AddUser(c echo.Context) error {
	var db = db.GetDB()

	user := models.User{}
	error := c.Bind(&user)
	if error != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

// GetUser will get a user from db
func GetUser(c echo.Context) error {
	db := db.GetDB()
	id := c.Param("id")

	user := models.User{}
	db.Where("id = ?", id).Find(&user)

	return c.JSON(http.StatusOK, user)
}

// UpdateUser ...
func UpdateUser(c echo.Context) error {
	db := db.GetDB()

	id := c.Param("id")
	user := models.User{}

	error := db.Where("id = ?", id).First(&user).Error
	if error != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	c.Bind(&user)
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

// DeleteUser ...
func DeleteUser(c echo.Context) error {
	db := db.GetDB()

	id := c.Param("id")
	db.Where("id = ?", id).Delete(models.User{})

	return c.String(http.StatusOK, "deleted with id: "+id)
}
