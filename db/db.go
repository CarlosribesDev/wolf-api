package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func InitDb(drop bool) (*sql.DB, error) {
	dbPath := os.Getenv("DATABASE_URL")

	log.Println("dbPath:", dbPath)
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		return nil, err
	}

	if drop {
		if err = dropTables(db); err != nil {
			return nil, err
		}
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	log.Println("Database and table initialized successfully")
	return db, nil
}

func createTables(db *sql.DB) error {
	tableQueries := []string{
		`CREATE TABLE IF NOT EXISTS player (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			role INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS game (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			host_id INTEGER NOT NULL,
			created_at TIMESTAMP NOT NULL,
			FOREIGN KEY (host_id) REFERENCES player(id)
		);`,
		`CREATE TABLE IF NOT EXISTS game_player (
			game_id INTEGER NOT NULL,
			player_id INTEGER NOT NULL,
			PRIMARY KEY (game_id, player_id),
			FOREIGN KEY (game_id) REFERENCES game(id),
			FOREIGN KEY (player_id) REFERENCES player(id)
		);`,
	}

	for _, query := range tableQueries {
		if _, err := db.Exec(query); err != nil {
			log.Println("Error creating table:", err)
			return err
		}
	}

	return nil
}

func dropTables(db *sql.DB) error {
	tableNames := []string{"player", "game", "game_player"}

	for _, table := range tableNames {
		query := `DROP TABLE IF EXISTS ` + table
		if _, err := db.Exec(query); err != nil {
			log.Printf("Failed to drop table %s: %v", table, err)
			return err
		}
		log.Printf("Table %s dropped successfully\n", table)
	}

	return nil
}
