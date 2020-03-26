package storage

import (
	"sync"

	"tildegit.org/andinus/perseus/storage/sqlite3"
)

// Init initializes the database.
func Init() *sqlite3.DB {
	var db sqlite3.DB = sqlite3.DB{
		Mu: new(sync.RWMutex),
	}

	sqlite3.Init(&db)
	return &db
}
