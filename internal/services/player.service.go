package services

import (
	"ribi-code/wolf-api/internal/domain"
	"ribi-code/wolf-api/internal/ports"
)

type PlayerServiceImpl struct {
	repository ports.PlayerRepository
}

func (p *PlayerServiceImpl) Create() (domain.Player, error) {
	newPlayer := domain.NewPlayer{}

	playerInserted, err := p.repository.Insert(newPlayer)

	if err != nil {
		return domain.Player{}, err
	}

	return playerInserted, nil
}

func NewPlayerServiceImpl(repository ports.PlayerRepository) ports.PlayerService {
	return &PlayerServiceImpl{repository: repository}
}
