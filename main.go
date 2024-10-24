package main

import (
	"github.com/gin-gonic/gin"
	"go-clean/cmd"
	"go-clean/config"
	"go-clean/db"
	"go-clean/middlewares"
	"go-clean/routes"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	db.InitDB()
	defer db.CloseDatabaseConnection(db.Data)

	if len(os.Args) > 1 {
		cmd.Commands(db.Data)
		return
	}

	server := gin.Default()

	server.Use(middlewares.CORSMiddleware())

	routes.SetupRoutes(server)

	var serve string
	if cfg.AppPort != "" {
		serve = "127.0.0.1:" + cfg.AppPort
	} else {
		serve = "127.0.0.1:" + "8080"
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
