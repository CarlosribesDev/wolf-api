package ports

import "ribi-code/wolf-api/internal/domain"

type GameRepository interface {
	Insert(newGame domain.NewGame) (domain.Game, error)
	GetById(id int) (domain.Game, error)
}

type GameService interface {
	Create(hostId int) (domain.Game, error)
	GetById(id int) (domain.Game, error)
}
