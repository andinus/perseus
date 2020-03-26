package sqlite3

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
