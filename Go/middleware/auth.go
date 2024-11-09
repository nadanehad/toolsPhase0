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
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		sessions.SessionStore.RLock()
		userID, exists := sessions.SessionStore.Sessions[sessionID]
		sessions.SessionStore.RUnlock()

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var user models.User
		if err := DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("userID", user.ID)
		c.Set("role", user.Role)

		c.Next()
	}
}
