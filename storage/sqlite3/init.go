package sqlite3

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// DB holds the database connection, mutex & path.
type DB struct {
	Path string
	Mu   *sync.RWMutex
	Conn *sql.DB
}

// initErr will log the error and close the database connection if
// necessary.
func initErr(db *DB, err error) {
	if db.Conn != nil {
		db.Conn.Close()
	}
	log.Fatalf("Initialization Error :: %s", err.Error())
}

// Init initializes a sqlite3 database.
func Init(db *DB) {
	var err error

	// We set the database path, first the environment variable
	// PERSEUS_DBPATH is checked. If it doesn't exist then use set
	// it to the default (perseus.db).
	envDBPath, exists := os.LookupEnv("PERSEUS_DBPATH")
	if !exists {
		envDBPath = "perseus.db"
	}
	db.Path = envDBPath

	db.Conn, err = sql.Open("sqlite3", db.Path)
	if err != nil {
		log.Printf("sqlite3/init.go: %s\n",
			"Failed to open database connection")
		initErr(db, err)
	}

	// Create account table, this will hold information on account
	// like id, type & other user specific information. We are
	// using id because later we may want to add username change
	// or account delete functionality. If we add user delete
	// function then we'll just have to change the username here.
	stmt, err := db.Conn.Prepare(`
CREATE TABLE IF NOT EXISTS account (
       id       TEXT PRIMARY KEY,
       type     TEXT NOT NULL DEFAULT user,
       username TEXT NOT NULL,
       password TEXT NOT NULL);`)

	if err != nil {
		log.Printf("sqlite3/init.go: %s\n",
			"Failed to prepare statement")
		initErr(db, err)
	}

	_, err = stmt.Exec()
	stmt.Close()
	if err != nil {
		log.Printf("sqlite3/init.go: %s\n",
			"Failed to execute statement")
		initErr(db, err)
	}
}
