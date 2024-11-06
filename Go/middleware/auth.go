package middleware

import (
	"net/http"
	"playlist/sessions" // Use the correct path to the sessions package

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		// Set the user ID in the context
		c.Set("userID", userID)
		c.Next()
	}
}
