package repository

import (
	"database/sql"
	"errors"
	"ribi-code/wolf-api/internal/domain"
	"ribi-code/wolf-api/internal/ports"
)

type GameRepositorySqlite struct {
	db *sql.DB
}

func (r *GameRepositorySqlite) Insert(newGame domain.NewGame) (domain.Game, error) {
	query := `INSERT INTO game (host_id, created_at) VALUES (?, ?)`
	result, err := r.db.Exec(query, newGame.HostId, newGame.CreatedAt)
	if err != nil {
		return domain.Game{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Game{}, err
	}

	return domain.Game{
		Id:        int(id),
		HostId:    newGame.HostId,
		CreatedAt: newGame.CreatedAt,
	}, nil
}

func (r *GameRepositorySqlite) GetById(id int) (domain.Game, error) {
	query := `SELECT id, host_id, created_at FROM game WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var game domain.Game
	err := row.Scan(&game.Id, &game.HostId, &game.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Game{}, nil
		}
		return domain.Game{}, err
	}
	return game, nil
}

func NewGameRepository(db *sql.DB) ports.GameRepository {
	return &GameRepositorySqlite{db: db}
}
