package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ROLE string

const (
	SUPERADMIN ROLE = "SUPERADMIN"
	ADMIN      ROLE = "ADMIN"
	USER       ROLE = "USER"
)

func Role(allowedRoles ...ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"errors": "User not authenticated"})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"code":   http.StatusForbidden,
				"errors": "Invalid role type"})
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
			c.JSON(http.StatusForbidden, gin.H{
				"code":   http.StatusForbidden,
				"errors": "Access denied"})
			c.Abort()
			return
		}

		// Proceed to the next middleware/handler
		c.Next()
	}
}
