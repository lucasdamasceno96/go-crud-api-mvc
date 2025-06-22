package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/config/logger"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/controller/routes"
)

func main() {
	logger.Info("Starting the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	v1 := router.Group("/api/v1")
	routes.InitRoutes(v1)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
