package utils

import (
	"fmt"
	"reflect"
)

func Exclude(obj interface{}, keys []string) (interface{}, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("exclude: expected a struct, got %s", val.Kind())
	}

	newObj := reflect.New(val.Type()).Elem() // Create a new struct instance

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Name

		// Check if the field should be excluded
		if contains(keys, fieldName) {
			continue
		}

		// Set the value in the new struct
		newObj.Field(i).Set(val.Field(i))
	}

	return newObj.Interface(), nil
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
