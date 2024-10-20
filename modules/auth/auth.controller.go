package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func Login(c *gin.Context) {
	var req ValidateLogin
	_ = c.BindJSON(&req)

	userData, errCredential := LoginWithUsernameAndPassword(req.Username, req.Password)

	if errCredential != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errCredential})
	}

	token, tokenErr := generateAuthToken(userData)

	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Could Not Generate  Token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"user":  userData,
		"token": token,
	})
}

func RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(input.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["token_type"] != "refresh" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Fetch the user from the database
	var user models.User
	userID := uint(claims["user_id"].(float64))
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Check token version
	tokenVersion := uint(claims["token_version"].(float64))
	if user.TokenVersion != tokenVersion {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been invalidated"})
		return
	}

	// Generate new tokens
	accessToken, err := generateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func Logout(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := userInterface.(models.User)
	user.TokenVersion += 1
	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not logout user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
