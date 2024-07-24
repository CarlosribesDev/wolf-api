package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	gameHandler := Compose()
	router := gin.Default()

	gameRoute := router.Group("/game")
	gameRoute.GET("/:id", gameHandler.FindGameById)
	gameRoute.POST("/", gameHandler.CreateNewGame)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
