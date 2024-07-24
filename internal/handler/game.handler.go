package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ribi-code/wolf-api/internal/ports"
	"strconv"
)

type GameHandler struct {
	gameService ports.GameService
}

type NewGamePayload struct {
	HostId int `json:"hostId"`
}

func NewGameHandler(gameService ports.GameService) *GameHandler {
	return &GameHandler{
		gameService: gameService,
	}
}

func (g *GameHandler) HostGame(c *gin.Context) {

}

func (g *GameHandler) FindGameById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	game, err := g.gameService.GetById(id)

	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (g *GameHandler) CreateNewGame(ctx *gin.Context) {
	var newGame NewGamePayload
	if err := ctx.ShouldBindJSON(&newGame); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	gameCreated, err := g.gameService.Create(newGame.HostId)

	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gameCreated)
}
