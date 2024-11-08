package controllers

import (
	"net/http"
	"playlist/models"
	"playlist/sessions"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DB *gorm.DB

func RegisterUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&input); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func LoginUser(c *gin.Context) {
	// Use the LoginRequest struct for the login data
	var loginReq models.LoginRequest
	var storedUser models.User

	// Bind incoming JSON request to loginReq struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Find the user by email
	if err := DB.Where("email = ?", loginReq.Email).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the plain text password from the request with the stored password
	if storedUser.Password != loginReq.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a new session ID for the user
	sessionID := uuid.New().String()

	// Store the session ID in session store
	sessions.SessionStore.Lock()
	sessions.SessionStore.Sessions[sessionID] = storedUser.ID
	sessions.SessionStore.Unlock()

	// Set the session ID as a cookie in the response
	c.SetCookie("session_id", sessionID, 3600, "/", "localhost", false, true)

	// Return the login response with session token and role
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   sessionID,
		"role":    storedUser.Role,
	})
}
