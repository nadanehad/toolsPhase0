package main

import (
	"fmt"
	"log"
	"os"
	"playlist/controllers"
	"playlist/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// MySQL connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}

	// Migrate the schema and set the DB for controllers
	models.AutoMigrate(DB)
	controllers.SetDB(DB)

	// Create a Gin router
	router := gin.Default()
	router.Use(cors.Default())

	// Define routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Start the server
	router.Run(":8080")
}
