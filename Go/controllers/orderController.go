package controllers

import (
	"net/http"
	"playlist/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	// Bind and validate JSON input
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve the userID from the context set by AuthMiddleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Set the UserID in the order model
	order.UserID = userID.(uint)

	// Save the order in the database
	if result := DB.Create(&order); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

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

	// Get the order ID from the URL parameter
	orderID := c.Param("order_id")

	// Find the order in the database
	var order models.Order
	if err := DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Return the order details as JSON
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
	// Retrieve the user ID from context (set by AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the order ID from the URL parameter
	orderID := c.Param("order_id")

	// Find the order in the database
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
