package models

import "gorm.io/gorm"

type StatusHistory struct {
	gorm.Model
	OrderID uint   `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}
