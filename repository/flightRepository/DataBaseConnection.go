package flightRepository

import (
	"fmt"
	"ticket/models"
	"ticket/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDbConnection() *gorm.DB {

	dbURI := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.ENV("DB_USERNAME"), utils.ENV("DB_PASSWORD"), utils.ENV("DB_PORT"), utils.ENV("DB_DATABASE"))
	// Connect to the database

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Flight{})
	if err != nil {
		panic(err)
	}
	return db
}
