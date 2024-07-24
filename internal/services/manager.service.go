package services

import "ribi-code/wolf-api/internal/ports"

type ManagerServiceImpl struct {
	gameService   ports.GameService
	playerService ports.PlayerService
}

func (m *ManagerServiceImpl) StartGame() {

}
