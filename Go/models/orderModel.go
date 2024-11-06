package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          uint   `json:"user_id"`
	PickupLocation  string `json:"pickup_location" binding:"required"`
	DropoffLocation string `json:"dropoff_location" binding:"required"`
	PackageDetails  string `json:"package_details"`
	DeliveryTime    string `json:"delivery_time"`
	Status          string `json:"status" gorm:"default:'Pending'"`
}
