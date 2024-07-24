package ports

import "ribi-code/wolf-api/internal/domain"

type ManagerService interface {
	InitGame() (domain.Game, error)
}
