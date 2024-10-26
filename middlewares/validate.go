package middlewares

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidationMiddleware(queryType interface{}, bodyType interface{}) gin.HandlerFunc {

	fmt.Println(queryType)
	return func(c *gin.Context) {
		// Validate query parameters
		if queryType != nil {
			queryObj := reflect.New(reflect.TypeOf(queryType).Elem()).Interface()
			if err := validate.Struct(queryObj); err != nil {
				var errors []string
				for _, err := range err.(validator.ValidationErrors) {
					errors = append(errors, fmt.Sprintf(" '%s' '%s'", err.Field(), err.Tag()))
				}
				c.JSON(http.StatusBadRequest, gin.H{
					"statusCode": http.StatusBadRequest,
					"error":      errors,
				})
				c.Abort()
				return
			}
		}

		// Validate request body
		if bodyType != nil {
			bodyObj := reflect.New(reflect.TypeOf(bodyType).Elem()).Interface()
			fmt.Println(bodyType)
			if err := validate.Struct(bodyObj); err != nil {
				var errors []string
				for _, err := range err.(validator.ValidationErrors) {
					errors = append(errors, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
				}
				c.JSON(http.StatusBadRequest, gin.H{
					"statusCode": http.StatusBadRequest,
					"error":      errors,
				})
				c.Abort()
				return
			}
		}

		// Continue to the next handler if validation passes
		c.Next()
	}
}
