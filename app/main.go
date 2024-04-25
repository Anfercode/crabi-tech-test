package main

import (
	"crabi-tech-test/app/api"
	"crabi-tech-test/src/user/application"
	"crabi-tech-test/src/user/domain"
	"crabi-tech-test/src/user/infrastructure"
	"crabi-tech-test/src/user/infrastructure/persistence"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&domain.User{})

	log.Println("Connected to database successfully.")

	router := gin.Default()

	userRepository := persistence.NewUserRepository(db)
	blacklistService := infrastructure.NewHTTPBlacklistValidator("http://44.210.144.170/check-blacklist")
	userService := application.NewUserService(userRepository, blacklistService)

	router.POST("/users", api.CreateUser(*userService))
	router.POST("/login", api.Login(*userService))
	router.GET("/users/:email", api.GetUser(*userService))
	router.Run(":8080")
}
