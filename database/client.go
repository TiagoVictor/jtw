package database

import (
	"jwt-authentication-golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal("Could not connect to the database", dbError)
		panic("Could not connect to the database")
	}
	log.Println("Connected to the database successfully")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migrated successfully")
}
