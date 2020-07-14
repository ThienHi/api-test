package db

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Init creates a connection to postgres database and
// migrates any new models
func Init() {
	config := config.GetConfig()
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_NAME, config.DB_PASSWORD)
	db, _ = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	db.AutoMigrate(&models.User{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB ...
func CloseDB() {
	db.Close()
}
