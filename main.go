package main

import (
	"github.com/go-playground/validator/v10"
	"go-clean/cmd"
	"go-clean/config"
	"go-clean/db"
	"go-clean/middlewares"
	"go-clean/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var validate *validator.Validate

func main() {
	config.LoadConfig()

	db.InitDB()
	defer db.CloseDatabaseConnection(db.Data)

	if len(os.Args) > 1 {
		cmd.Commands(db.Data)
		return
	}

	middlewares.InitValidator()
	server := gin.Default()

	server.Use(middlewares.CORSMiddleware())

	routes.SetupRoutes(server)

	var serve string
	if config.Data.AppPort != "" {
		serve = "127.0.0.1:" + config.Data.AppPort
	} else {
		serve = "127.0.0.1:" + "8080"
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
