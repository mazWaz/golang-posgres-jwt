package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidationMiddleware validates the query parameters and request body against the provided struct types
func ValidationMiddleware(queryType interface{}, bodyType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate query parameters
		if queryType != nil {
			queryObj := reflect.New(reflect.TypeOf(queryType).Elem()).Interface()
			if err := c.ShouldBindQuery(queryObj); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"statusCode": http.StatusBadRequest,
					"error":      "Invalid query parameters: " + err.Error(),
				})
				c.Abort()
				return
			}
			if err := validate.Struct(queryObj); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"statusCode": http.StatusBadRequest,
					"error":      "Query validation failed: " + err.Error(),
				})
				c.Abort()
				return
			}
			// Optionally pass the validated query object to the next handler
			c.Set("validatedQuery", queryObj)
		}

		// Validate request body
		if bodyType != nil {
			bodyObj := reflect.New(reflect.TypeOf(bodyType).Elem()).Interface()
			if err := c.ShouldBindJSON(bodyObj); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"statusCode": http.StatusBadRequest,
					"error":      "Invalid request body: " + err.Error(),
				})
				c.Abort()
				return
			}
			if err := validate.Struct(bodyObj); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"statusCode": http.StatusBadRequest,
					"error":      "Body validation failed: " + err.Error(),
				})
				c.Abort()
				return
			}
			// Optionally pass the validated body object to the next handler
			c.Set("validatedBody", bodyObj)
		}

		// Continue to the next handler if validation passes
		c.Next()
	}
}
