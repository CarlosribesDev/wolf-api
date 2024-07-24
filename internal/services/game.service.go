package services

import (
	"ribi-code/wolf-api/internal/domain"
	"ribi-code/wolf-api/internal/ports"
	"time"
)

type GameServiceImpl struct {
	repository ports.GameRepository
}

func (g *GameServiceImpl) Create(hostId int) (domain.Game, error) {
	newGame := domain.NewGame{
		HostId:    hostId,
		CreatedAt: time.Now().UTC(),
	}

	gameInserted, err := g.repository.Insert(newGame)

	if err != nil {
		return domain.Game{}, err
	}

	return gameInserted, nil
}

func (g *GameServiceImpl) GetById(hostId int) (domain.Game, error) {
	game, err := g.repository.GetById(hostId)

	if err != nil {
		return domain.Game{}, err
	}

	return game, nil
}

func NewGameServiceImpl(repository ports.GameRepository) ports.GameService {
	return &GameServiceImpl{
		repository: repository,
	}
}
