package main

import (
	"fmt"
	"log"
	"os"
	"playlist/controllers"
	"playlist/middleware"
	"playlist/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	controllers.DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to connect to database:", errDB)
	}

	// Run migrations
	controllers.DB.AutoMigrate(&models.User{}, &models.Order{})
}

func main() {
	initDB()

	router := gin.Default()

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.GET("/courier/:courier_id/orders", controllers.GetOrdersByCourierID)
	router.PUT("/order/:order_id/status", controllers.UpdateOrderStatus)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/create-order", controllers.CreateOrder)
		protected.GET("/orders", controllers.GetUserOrders)
		protected.GET("/order/:order_id", controllers.GetOrderDetails)
		protected.POST("/order/:order_id/cancel", controllers.CancelOrder)
	}

	router.Run(":8080")
}
