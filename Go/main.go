package main

import (
	"fmt"
	"log"
	"net/http"
	"time" 
	"os"
	"playlist/controllers"
	"playlist/middleware"
	"playlist/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
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

	DB, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to connect to database:", errDB)
	}

	DB.AutoMigrate(&models.User{}, &models.Order{}, &models.StatusHistory{})
	return DB
}

func main() {
	DB := initDB()

	controllers.DB = DB

	router := gin.Default()

	router.Use(cors.New(cors.Config{
	AllowOrigins: []string{
		"https://frontend-nadanehad-dev.apps.rm2.thpm.p1.openshiftapps.com", // Frontend URL
		
	},
	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Include OPTIONS method
    AllowHeaders: []string{
        "Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With",
    },
    ExposeHeaders: []string{
        "Content-Length", "Access-Control-Allow-Origin",
    },
    AllowCredentials: true, // Allow cookies
    MaxAge: 12 * time.Hour, // Cache the CORS preflight response
}))
	router.OPTIONS("/*path", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Status(http.StatusOK)
    })



	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(DB))
	{
		protected.POST("/create-order", controllers.CreateOrder)
		protected.GET("/orders", controllers.GetUserOrders)
		protected.GET("/order/:order_id", controllers.GetOrderDetails)
		protected.POST("/order/:order_id/cancel", controllers.CancelOrder)
	}

	courier := router.Group("/courier")
	courier.Use(middleware.AuthMiddleware(DB), CourierOnlyMiddleware())
	{
		courier.GET("/:courier_id/orders", controllers.GetOrdersByCourierID)
		courier.POST("/order/:order_id/accept", controllers.AcceptOrDeclineOrder)
		courier.PUT("/order/:order_id/status", controllers.UpdateOrderStatus)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(DB), AdminOnlyMiddleware())
	{
		admin.POST("/assign-order", controllers.AssignOrder)
		admin.GET("/orders", controllers.GetAllOrders)
		admin.PUT("/order/:order_id", controllers.UpdateOrder)
		admin.DELETE("/order/:order_id", controllers.DeleteOrder)
		admin.GET("/assigned-orders", controllers.GetAwaitingCourierAcceptanceOrders)
		admin.PUT("/reassign-orders", controllers.ReassignOrders)
	}

	router.Run(":8080")
}

func CourierOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "Courier" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Couriers only"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "Admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
