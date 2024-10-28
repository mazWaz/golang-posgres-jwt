package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

// InitValidator initializes the validator instance and custom translations
func InitValidator() {
	validate = validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")

	err := entranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return
	}

	// Register custom translations
	err = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field())
		return field
	})
	if err != nil {
		return
	}

	err = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field())
		return field
	})
	if err != nil {
		return
	}

	err = validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} must be at least {1} characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
		return field
	})
	if err != nil {
		return
	}

	err = validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} cannot exceed {1} characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
		return field
	})
	if err != nil {
		return
	}

	err = validate.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} must be greater than or equal to {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
		return field
	})
	if err != nil {
		return
	}

	err = validate.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0} must be less than or equal to {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
		return field
	})
	if err != nil {
		return
	}
}

// Convert validation errors to custom messages
func validationErrors(err error) gin.H {
	errors := gin.H{}
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Translate(trans)
	}
	return errors
}

// ValidationMiddleware Middleware to validate query and body, allowing nil values
func ValidationMiddleware(queryObj interface{}, bodyObj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate query parameters if queryObj is provided
		if queryObj != nil {
			query := reflect.New(reflect.TypeOf(queryObj).Elem()).Interface()
			if err := c.ShouldBindQuery(query); err == nil {
				if err := validate.Struct(query); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": validationErrors(err),
					})
					c.Abort()
					return
				}
			}
		}

		// Validate JSON body without binding it in Gin
		if bodyObj != nil {
			body := reflect.New(reflect.TypeOf(bodyObj).Elem()).Interface()
			if jsonBytes, err := io.ReadAll(c.Request.Body); err == nil {
				// Check if body is empty
				if len(bytes.TrimSpace(jsonBytes)) == 0 {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Empty body",
					})
					c.Abort()
					return
				}

				// Reset the request body so that the controller can read it again
				c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

				// Unmarshal JSON for validation without consuming the body
				if err := json.Unmarshal(jsonBytes, body); err == nil {
					if err := validate.Struct(body); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"code":   http.StatusBadRequest,
							"errors": validationErrors(err),
						})
						c.Abort()
						return
					}
				} else {
					// If JSON unmarshalling fails, return a specific error
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Invalid JSON format",
					})
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
