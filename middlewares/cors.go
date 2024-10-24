package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		//requestedWith := c.GetHeader("X-Requested-With")
		//if requestedWith != "XMLHttpRequest" {
		//	// Deny the request if the header is missing or not as expected
		//	c.JSON(http.StatusForbidden, gin.H{
		//		"code":  http.StatusForbidden,
		//		"error": "Forbidden",
		//	})
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}
