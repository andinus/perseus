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
		// Create users table, this will hold information on
		// account like id, type & other user specific
		// information. We are using id because later we may
		// want to add username change or account delete
		// functionality. username here is not unique because
		// if user deletes account then we'll change it to
		// "ghost" or something. This doesn't mean usernames
		// shouldn't be unique, registration table requires
		// them to be unique so it'll fail if they aren't
		// unique.
		`CREATE TABLE IF NOT EXISTS users (
       id       TEXT PRIMARY KEY,
       type     TEXT NOT NULL DEFAULT notadmin,
       username TEXT NOT NULL,
       password TEXT NOT NULL);`,

		// Create registration table, this will hold user
		// account details like registration time, ip &
		// similar details. This is the only place that will
		// relate the username to id even after deletion.
		// usernames must be unique in this table.
		`CREATE TABLE IF NOT EXISTS registration (
       id       TEXT PRIMARY KEY,
       username TEXT NOT NULL UNIQUE,
       reg_time TEXT NOT NULL,
       reg_ip   TEXT NOT NULL);`,
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
