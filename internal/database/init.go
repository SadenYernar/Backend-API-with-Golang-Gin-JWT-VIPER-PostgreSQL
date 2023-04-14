package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// // DB set up
func SetupDB() (*sql.DB, error) {
	connStr := "postgres://postgres:qwerty@localhost:5432/students?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	err = CreatTables(db)
	if err != nil {
		return nil, fmt.Errorf("can't table create %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}
	// query := `PRAGMA foreign_keys=1;`
	// stmt, err := db.Prepare(query)
	// if err != nil {
	// 	return nil, err
	// }
	// stmt.Exec()
	return db, nil
}
