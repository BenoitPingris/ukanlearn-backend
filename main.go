package main

import (
	"log"
	"ukanlearn/app"
	"ukanlearn/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDbOrDie() *gorm.DB {
	dsn := "host=localhost user=ukanlearn password=docker dbname=ukanlearn port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}

func main() {
	app.Init(setupDbOrDie())
}
