package main

import (
	"log"
	"ribi-code/wolf-api/db"
	"ribi-code/wolf-api/internal/handler"
	"ribi-code/wolf-api/internal/repository"
	"ribi-code/wolf-api/internal/services"
)

func Compose() *handler.GameHandler {

	dbConnection, err := db.InitDb(false)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	gameRepository := repository.NewGameRepository(dbConnection)
	gameService := services.NewGameServiceImpl(gameRepository)
	gameHandler := handler.NewGameHandler(gameService)

	return gameHandler
}
