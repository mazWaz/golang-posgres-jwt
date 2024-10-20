package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-clean/config"
)

var validate *validator.Validate

func main() {
	cfg := config.LoadConfig()
	fmt.Println("Database Connection String:", cfg.DbPassword)
}
