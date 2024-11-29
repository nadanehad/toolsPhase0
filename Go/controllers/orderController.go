package controllers

import (
	"log"
	"net/http"
	"playlist/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Helper function to handle datetime parsing and formatting
func formatDeliveryTime(dt string) (string, error) {
	// Remove 'Z' suffix if it's present (UTC indicator)
	if len(dt) > 0 && dt[len(dt)-1] == 'Z' {
		dt = dt[:len(dt)-1]
	}

	// Check if the string has at least the date and hour, and add seconds if missing
	if len(dt) == 16 { // Format like '2024-12-07T11:01'
		dt = dt + ":00" // Append missing seconds
	}

	// Parse the datetime string into the expected format
	parsedTime, err := time.Parse("2006-01-02T15:04:05", dt)
	if err != nil {
		log.Println("Datetime Parsing Error:", err)
		return "", err
	}

	// Return the formatted time in MySQL DATETIME format
	return parsedTime.Format("2006-01-02 15:04:05"), nil
}

// CreateOrder handles the creation of a new order
func CreateOrder(c *gin.Context) {
	var order models.Order

	// Bind incoming JSON to the order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("Bind Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Format delivery time if it's provided
	if order.DeliveryTime != "" {
		formattedTime, err := formatDeliveryTime(order.DeliveryTime)
		if err != nil {
			log.Println("Error formatting delivery time:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format"})
			return
		}
		order.DeliveryTime = formattedTime
	}

	// Check for user ID in the context
	userID, exists := c.Get("userID")
	if !exists {
		log.Println("Unauthorized: No user ID found")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	order.UserID = userID.(uint)

	// Create the order in the database
	if result := DB.Create(&order); result.Error != nil {
		log.Println("Database Create Error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	log.Println("Order created successfully for User ID:", userID)
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// GetUserOrders returns all orders for the logged-in user
func GetUserOrders(c *gin.Context) {
	// Get the user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Fetch the orders for the user
	var orders []models.Order
	if err := DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve orders"})
		return
	}

	// Return the orders
	c.JSON(http.StatusOK, orders)
}

// GetOrderDetails returns the details of a specific order for the logged-in user
func GetOrderDetails(c *gin.Context) {
	// Get the user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Fetch the order by ID
	orderID := c.Param("order_id")
	var order models.Order
	if err := DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Return the order details
	c.JSON(http.StatusOK, gin.H{
		"id":               order.ID,
		"pickup_location":  order.PickupLocation,
		"dropoff_location": order.DropoffLocation,
		"package_details":  order.PackageDetails,
		"delivery_time":    order.DeliveryTime,
		"status":           order.Status,
	})
}

// CancelOrder cancels an order if the status is 'Pending'
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

	if order.DeliveryTime != "" {
		formattedTime, err := formatDeliveryTime(order.DeliveryTime)
		if err != nil {
			log.Println("Error formatting delivery time:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format"})
			return
		}
		order.DeliveryTime = formattedTime
	}

	// Only allow cancellation of orders with status 'Pending'
	if order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only pending orders can be canceled"})
		return
	}

	// Update the order status to 'Canceled'
	order.Status = "Canceled"
	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel the order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}
