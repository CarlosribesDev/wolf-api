package repository

import "database/sql"

func InsertAndReturnID(db *sql.DB, query string, args ...interface{}) (int, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
