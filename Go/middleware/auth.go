package middleware

import (
	"net/http"
	"playlist/sessions"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve session_id from the cookies
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid session ID"})
			c.Abort()
			return
		}

		// Set the userID in the context for use in the handler functions
		c.Set("userID", userID)
		c.Next()
	}
}
