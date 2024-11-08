package middleware

import (
	"net/http"
	"playlist/models"
	"playlist/sessions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the session ID from the cookie
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the session ID exists in the session store
		sessions.SessionStore.RLock()
		userID, exists := sessions.SessionStore.Sessions[sessionID]
		sessions.SessionStore.RUnlock()

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Query the database to get the user and their role
		var user models.User
		if err := DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Set the user ID and role in the context for role-based access control
		c.Set("userID", user.ID)
		c.Set("role", user.Role)

		// Continue to the next middleware or handler
		c.Next()
	}
}
