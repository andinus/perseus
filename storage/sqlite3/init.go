package sqlite3

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

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
	// it to the default (perseus.db). Note that this is LookupEnv
	// so if the user has set PERSEUS_DBPATH="" then it'll return
	// true for exists as it should because technically user has
	// set the env var, the sql.Open statement will fail though.
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

	sqlstmt := []string{
		// Users can login with multiple devices and so
		// multiple tokens will be created. This shouldn't be
		// used for login, logins should be verified with
		// users table only.
		`CREATE TABLE IF NOT EXISTS access (
       id       TEXT NOT NULL,
       token    TEXT NOT NULL,
       genTime TEXT NOT NULL);`,

		`CREATE TABLE IF NOT EXISTS users (
       id       TEXT PRIMARY KEY,
       type     TEXT NOT NULL DEFAULT user,
       username VARCHAR(128) NOT NULL UNIQUE,
       password TEXT NOT NULL,
       regTime  TEXT NOT NULL);`,
	}

	// We range over statements and execute them one by one, this
	// is during initialization so it doesn't matter if it takes
	// few more ms. This way we know which statement caused the
	// program to fail.
	for _, s := range sqlstmt {
		stmt, err := db.Conn.Prepare(s)

		if err != nil {
			log.Printf("sqlite3/init.go: %s\n",
				"Failed to prepare statement")
			log.Println(s)
			initErr(db, err)
		}

		_, err = stmt.Exec()
		stmt.Close()
		if err != nil {
			log.Printf("sqlite3/init.go: %s\n",
				"Failed to execute statement")
			log.Println(s)
			initErr(db, err)
		}
	}
}
