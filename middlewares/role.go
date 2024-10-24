package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-clean/modules/user"
	"net/http"
)

func Role(allowedRoles ...user.ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role type"})
			c.Abort()
			return
		}

		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if roleStr == string(allowedRole) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Proceed to the next middleware/handler
		c.Next()
	}
}
