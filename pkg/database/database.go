package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(filePath string) (*sql.DB, error) {
	db, err := sql.Open(SQLite3Str, filePath)
	return db, err
}