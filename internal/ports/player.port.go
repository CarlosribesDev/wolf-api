package ports

import "ribi-code/wolf-api/internal/domain"

type PlayerService interface {
	Create() (domain.Player, error)
}

type PlayerRepository interface {
	Insert(newPlayer domain.NewPlayer) (domain.Player, error)
	GetById(id int) (domain.Player, error)
}
