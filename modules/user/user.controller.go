package user

//import (
//	"github.com/gin-gonic/gin"
//	"golang.org/x/crypto/bcrypt"
//	"net/http"
//)
//
//func GetProfile(c *gin.Context) {
//	user, _ := c.Get("user")
//	c.JSON(http.StatusOK, user)
//}
//
//func GetUserByID(c *gin.Context) {
//	var user models.User
//	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		return
//	}
//	c.JSON(http.StatusOK, user)
//}
//
//func GetAllUsers(c *gin.Context) {
//	var users []models.User
//	models.DB.Find(&users)
//	c.JSON(http.StatusOK, users)
//}
//
//func UpdateUser(c *gin.Context) {
//	var user models.User
//	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		return
//	}
//
//	var input validations.UpdateUserInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	if input.Password != "" {
//		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
//		input.Password = string(hashedPassword)
//	}
//
//	models.DB.Model(&user).Updates(input)
//	c.JSON(http.StatusOK, user)
//}
//
//func DeleteUser(c *gin.Context) {
//	var user models.User
//	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		return
//	}
//
//	models.DB.Delete(&user)
//	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
//}
