package repository

import (
	"database/sql"
	"ribi-code/wolf-api/internal/domain"
	"ribi-code/wolf-api/internal/ports"
)

type PlayerRepositoryImpl struct {
	db *sql.DB
}

func (p *PlayerRepositoryImpl) Insert(newPlayer domain.NewPlayer) (domain.Player, error) {
	query := `INSERT INTO player (role) DEFAULT VALUES`
	result, err := p.db.Exec(query)

	if err != nil {
		return domain.Player{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Player{}, err
	}

	return domain.Player{
		Id:   int(id),
		Role: domain.Roles.Unknown,
	}, nil

}

func (p *PlayerRepositoryImpl) GetById(id int) (domain.Player, error) {
	query := `SELECT * FROM game WHERE id = ?`
	row := p.db.QueryRow(query, id)

	var player domain.Player
	if err := row.Scan(&player.Id, &player.Role); err != nil {
		return domain.Player{}, err
	}

	return player, nil
}

func NewPlayerRepository(db *sql.DB) ports.PlayerRepository {
	return &PlayerRepositoryImpl{
		db: db,
	}
}
