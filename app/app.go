package app

import (
	"log"
	"net/http"
	"ukanlearn/app/models"

	"github.com/go-redis/redis"
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

func setupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "docker",
		DB:       0,
	})
}

// Init function
func Init() {
	r := SetupRouter(setupDbOrDie())
	http.ListenAndServe(":3001", r)
}
