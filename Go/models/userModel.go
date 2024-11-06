package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
<<<<<<< HEAD
	Role     string `json:"role"`
=======
>>>>>>> 26105a9d8953654e5fa465ca09aca7589d85d8f5
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&User{}, &Order{})
}
