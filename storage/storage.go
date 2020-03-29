package storage

import (
	"database/sql"
	"sync"
)

// DB holds the database connection, mutex & path.
type DB struct {
	Path string
	Mu   *sync.RWMutex
	Conn *sql.DB
}

// Init initializes the database.
func Init() *DB {
	db := DB{
		Mu: new(sync.RWMutex),
	}

	initDB(&db)
	return &db
}
