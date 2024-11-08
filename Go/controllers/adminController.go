package controllers

import (
	"net/http"
	"playlist/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AssignOrder(c *gin.Context) {
	var assignInput struct {
		OrderID   uint `json:"order_id" binding:"required"`
		CourierID uint `json:"courier_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&assignInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Fetch the order to be assigned
	var order models.Order
	if result := DB.First(&order, assignInput.OrderID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Ensure the order is in a state that can be assigned 
	if order.Status != "Pending Assignment" && order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order cannot be assigned in its current status"})
		return
	}

	// Assign the order to the courier
	order.CourierID = &assignInput.CourierID
	order.Status = "Awaiting Courier Acceptance" 

	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign order"})
		return
	}

	 Add an entry in StatusHistory to log the assignment action
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
	var orders []models.Order
	if result := DB.Find(&orders); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c *gin.Context) {
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

func GetAndManageCourierOrders(c *gin.Context) {
	courierIDParam := c.Param("courier_id")
	courierID, err := strconv.Atoi(courierIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid courier ID"})
		return
	}

	var orders []models.Order
	if result := DB.Where("courier_id = ?", courierID).Find(&orders); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func ReassignOrder(c *gin.Context) {
	var input struct {
		NewCourierID uint `json:"new_courier_id" binding:"required"`
	}
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

	order.CourierID = input.NewCourierID
	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reassign courier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order reassigned successfully"})
}
