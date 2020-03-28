package user

import (
	"log"
	"time"

	"tildegit.org/andinus/perseus/storage/sqlite3"
)

// AddUser adds the user to record.
func (u User) AddUser(db *sqlite3.DB) error {
	// Acquire write lock on the database.
	db.Mu.Lock()
	defer db.Mu.Unlock()

	// Start the transaction
	tx, err := db.Conn.Begin()
	if err != nil {
		log.Printf("user/adduser.go: %s\n",
			"failed to begin transaction")
		return err
	}

	usrStmt, err := db.Conn.Prepare(`
INSERT INTO users(id, username, password, regTime) values(?, ?, ?, ?)`)
	if err != nil {
		log.Printf("user/adduser.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer usrStmt.Close()

	_, err = usrStmt.Exec(u.ID, u.Username, u.Password, time.Now().UTC())
	if err != nil {
		log.Printf("user/adduser.go: %s\n",
			"failed to execute statement")
		return err
	}

	tx.Commit()
	return err
}
