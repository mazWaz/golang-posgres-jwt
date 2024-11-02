package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

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

type Validator struct {
	Param interface{}
	Query interface{}
	Body  interface{}
}

// InitValidator initializes the validator and sets up custom translations
func InitValidator() {
	validate = validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	var err error
	trans, _ = uni.GetTranslator("en")

	err = entranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return
	}

	// Register custom translations for various tags
	registerCustomTranslations("required", "{0} is a required field")
	registerCustomTranslations("email", "{0} must be a valid email address")
	registerCustomTranslations("min", "{0} must be at least {1} characters")
	registerCustomTranslations("max", "{0} cannot exceed {1} characters")
	registerCustomTranslations("gte", "{0} must be greater than or equal to {1}")
	registerCustomTranslations("lte", "{0} must be less than or equal to {1}")
}

func registerCustomTranslations(tag, message string) {
	err := validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
		return field
	})
	if err != nil {
		return
	}
}

func validationErrors(err error) gin.H {
	errors := gin.H{}
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Translate(trans)
	}
	return errors

}

func getStructFields(obj interface{}) map[string]struct{} {
	fields := make(map[string]struct{})
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		jsonTag := typ.Field(i).Tag.Get("json")
		if jsonTag != "" {
			// Handle tags like `json:"username,omitempty"`
			tagParts := strings.Split(jsonTag, ",")
			fieldName := tagParts[0]
			if fieldName != "-" && fieldName != "" {
				fields[fieldName] = struct{}{}
			}
		} else {
			// If no JSON tag, use the field name in lowercase
			fieldName := strings.ToLower(typ.Field(i).Name)
			fields[fieldName] = struct{}{}
		}
	}

	return fields
}

func ValidationMiddleware(v Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v.Param != nil {
			param := reflect.New(reflect.TypeOf(v.Param).Elem()).Interface()
			if err := c.ShouldBindUri(param); err == nil {
				// Langsung validasi struct tanpa parsing JSON
				if err := validate.Struct(param); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": validationErrors(err),
					})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":   http.StatusBadRequest,
					"errors": "Invalid URI parameters",
				})
				c.Abort()
				return
			}
		}

		if v.Query != nil {
			query := reflect.New(reflect.TypeOf(v.Query).Elem()).Interface()
			if err := c.ShouldBindQuery(query); err == nil {
				// Convert query parameters to JSON format
				jsonBytes, err := json.Marshal(c.Request.URL.Query())
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Unable to parse query parameters",
					})
					c.Abort()
					return
				}

				// Parse JSON query data to payloadMap
				var payloadMap map[string]interface{}
				if err := json.Unmarshal(jsonBytes, &payloadMap); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Invalid query format",
					})
					c.Abort()
					return
				}

				// Get the known fields from the struct
				knownFields := getStructFields(v.Query)

				var unknownFields []string
				for key := range payloadMap {
					if _, exists := knownFields[key]; !exists {
						unknownFields = append(unknownFields, key)
					}
				}

				// Validate query structure
				if err := validate.Struct(query); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": validationErrors(err),
					})
					c.Abort()
					return
				}

				// Check for unknown fields
				if len(unknownFields) > 0 {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": fmt.Sprintf("Unknown or invalid query parameter(s): %s", strings.Join(unknownFields, ", ")),
					})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":   http.StatusBadRequest,
					"errors": "Invalid query parameters",
				})
				c.Abort()
				return
			}
		}

		if v.Body != nil {
			body := reflect.New(reflect.TypeOf(v.Body).Elem()).Interface()
			if jsonBytes, err := io.ReadAll(c.Request.Body); err == nil {
				if len(bytes.TrimSpace(jsonBytes)) == 0 {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Empty body",
					})
					c.Abort()
					return
				}

				c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

				var payloadMap map[string]interface{}
				if err := json.Unmarshal(jsonBytes, &payloadMap); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Invalid JSON format",
					})
					c.Abort()
					return
				}

				// Get the known fields from the struct
				knownFields := getStructFields(v.Body)

				var unknownFields []string
				for key := range payloadMap {
					if _, exists := knownFields[key]; !exists {
						unknownFields = append(unknownFields, key)
					}
				}

				if err := json.Unmarshal(jsonBytes, body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": "Invalid JSON structure",
					})
					c.Abort()
					return
				}

				if err := validate.Struct(body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": validationErrors(err),
					})
					c.Abort()
					return
				}

				if len(unknownFields) > 0 {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":   http.StatusBadRequest,
						"errors": fmt.Sprintf("Unknown or invalid JSON field(s): %s", strings.Join(unknownFields, ", ")),
					})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":   http.StatusBadRequest,
					"errors": "Unable to read request body",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
