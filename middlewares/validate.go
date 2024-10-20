package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

// Initialize the validator instance
func init() {
	validate = validator.New()
	// Add Custom Validation

	//validate.RegisterValidation("", func)
}

// ValidationMiddleware function
func ValidationMiddleware(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the JSON body to the object
		//if err := c.ShouldBindJSON(obj); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		//	c.Abort()
		//	return
		//}

		// Validate the struct
		if err := validate.Struct(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "error": err.Error()})
			c.Abort()
			return
		}

		// Continue to the next handler if validation passes
		c.Next()
	}
}
