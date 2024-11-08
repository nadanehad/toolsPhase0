package controllers

import (
	"net/http"
	"playlist/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AcceptOrDeclineOrder(c *gin.Context) {
	var action struct {
		Accept bool `json:"accept"`
	}
	orderIDParam := c.Param("order_id")
	courierID, exists := c.Get("courierID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&action); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var order models.Order
	if result := DB.Where("id = ? AND courier_id = ?", orderIDParam, courierID).First(&order); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or not assigned to this courier"})
		return
	}

	if action.Accept {
		order.Status = "In Progress"
	} else {
		order.Status = "Pending Assignment"
		order.CourierID = nil // Reset courier assignment
	}

	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order acceptance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order acceptance status updated"})
}

func GetOrdersByCourierID(c *gin.Context) {
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

func UpdateOrderStatus(c *gin.Context) {
	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}
	orderIDParam := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var order models.Order
	if result := DB.First(&order, orderID); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
		}
		return
	}

	order.Status = statusUpdate.Status
	if err := DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	// Log the status change in StatusHistory
	statusHistory := models.StatusHistory{
		OrderID: uint(orderID),
		Status:  statusUpdate.Status,
	}
	if err := DB.Create(&statusHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record status history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated and recorded successfully"})
}
