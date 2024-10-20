package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("userID", claims["userID"])
		c.Set("role", claims["role"])
		c.Next()
	}
}

//func AuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Get the token from the header
//		authHeader := c.GetHeader("Authorization")
//		if authHeader == "" {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
//			c.Abort()
//			return
//		}
//
//		// Extract the token
//		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
//
//		// Parse the token
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			return []byte(os.Getenv("JWT_SECRET")), nil
//		})
//
//		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//			// Retrieve the user from the database
//			var user models.User
//			userID := uint(claims["user_id"].(float64))
//			if err := models.DB.First(&user, userID).Error; err != nil {
//				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
//				c.Abort()
//				return
//			}
//
//			// Store the user in the context
//			c.Set("user", user)
//			c.Next()
//		} else {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//			c.Abort()
//			return
//		}
//	}
//}
