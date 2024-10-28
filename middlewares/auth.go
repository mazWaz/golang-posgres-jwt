package middlewares

import (
	"fmt"
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"errors": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Get JWT secret from environment variable
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":   http.StatusInternalServerError,
				"errors": "JWT secret not configured",
			})
			c.Abort()
			return
		}

		// Parse the token with validation of the signing method
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token's signing method is as expected
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"errors": "Failed to parse token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// Validate the token and extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Set("role", claims["role"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"errors": "Invalid token claims",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
