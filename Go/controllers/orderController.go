package controllers

import (
	"log"
	"net/http"
	"playlist/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	// Logging the incoming JSON data
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("Bind Error:", err) // Debug log for binding issue
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		log.Println("Unauthorized: No user ID found") // Debug log for unauthorized user
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	log.Println("User ID:", userID) // Log the User ID

	order.UserID = userID.(uint)

	if result := DB.Create(&order); result.Error != nil {
		log.Println("Database Create Error:", result.Error) // Debug log for database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	log.Println("Order created successfully for User ID:", userID) // Log successful creation

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// fetch the orders for the logged in user

func GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var orders []models.Order
	if err := DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// info of a specific order by order ID
func GetOrderDetails(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orderID := c.Param("order_id")

	var order models.Order
	if err := DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               order.ID,
		"pickup_location":  order.PickupLocation,
		"dropoff_location": order.DropoffLocation,
		"package_details":  order.PackageDetails,
		"delivery_time":    order.DeliveryTime,
		"status":           order.Status,
	})
}

// Endpoint to cancel a specific order by order ID
func CancelOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orderID := c.Param("order_id")

	var order models.Order
	if err := DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only pending orders can be canceled"})
		return
	}

	order.Status = "Canceled"
	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel the order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}
