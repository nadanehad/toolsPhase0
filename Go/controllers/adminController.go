package controllers

import (
	"fmt"
	"net/http"
	"playlist/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AssignOrder(c *gin.Context) {
	// Check if the user's role is "admin"
	role, _ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var assignInput struct {
		OrderID   uint `json:"order_id" binding:"required"`
		CourierID uint `json:"courier_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&assignInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	var order models.Order
	if result := DB.First(&order, assignInput.OrderID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.Status != "Pending Assignment" && order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order cannot be assigned in its current status"})
		return
	}

	order.CourierID = assignInput.CourierID
	order.Status = "Awaiting Courier Acceptance"

	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign order"})
		return
	}

	statusHistory := models.StatusHistory{
		OrderID: order.ID,
		Status:  "Assigned to Courier",
	}
	if err := DB.Create(&statusHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record assignment in status history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order assigned successfully", "order": order})
}

func GetAllOrders(c *gin.Context) {
	// Check if the user's role is "admin"
	role, _ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var orders []models.Order
	if result := DB.Find(&orders); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c *gin.Context) {
	// Check if the user's role is "admin"
	role, _ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Order
	orderIDParam := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var order models.Order
	if err := DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.PickupLocation = input.PickupLocation
	order.DropoffLocation = input.DropoffLocation
	order.PackageDetails = input.PackageDetails
	order.DeliveryTime = input.DeliveryTime
	order.DeliveryFee = input.DeliveryFee
	order.Status = input.Status
	order.CourierID = input.CourierID

	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(c *gin.Context) {
	// Check if the user's role is "admin"
	role, _ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	orderIDParam := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := DB.Delete(&models.Order{}, orderID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func GetAwaitingCourierAcceptanceOrders(c *gin.Context) {
	// Check if the user's role is "admin"
	role, _ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Retrieve all orders with the "Awaiting Courier Acceptance" status
	var orders []models.Order
	if result := DB.Where("status = ?", "Awaiting Courier Acceptance").Find(&orders); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	// Return the orders with "Awaiting Courier Acceptance" status
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// ReassignOrders function

func ReassignOrders(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Ensure field names match the frontend payload
	var requestBody struct {
		OrderID      uint `json:"order_id" binding:"required"`
		NewCourierID uint `json:"new_courier_id" binding:"required"` // Correct the field name to match frontend
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		fmt.Println("Error binding JSON:", err) // Log the error for debugging
		return
	}

	var order models.Order
	// Ensure the order exists in the DB
	if result := DB.First(&order, requestBody.OrderID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// If the order was declined, reset CourierID and change the status to "Pending"
	if order.Status == "Awaiting Courier Acceptance" || order.Status == "Declined" {
		order.Status = "Pending" // Change the status back to Pending or any suitable state
		order.CourierID = 0      // Reset CourierID since the order was declined
	}

	// Perform reassign
	order.CourierID = requestBody.NewCourierID
	order.Status = "Awaiting Courier Acceptance" // Change status to "Awaiting Courier Acceptance"

	// Save updated order
	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reassign order"})
		return
	}

	// Create status history for reassignment
	statusHistory := models.StatusHistory{
		OrderID: order.ID,
		Status:  "Reassigned to new courier",
	}
	if err := DB.Create(&statusHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record reassignment history"})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{"message": "Order reassigned successfully", "order": order})
}
