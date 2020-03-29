package account

import (
	"log"
	"time"

	"tildegit.org/andinus/perseus/storage"
)

// addUser adds the user to record.
func (u *User) addUser(db *storage.DB) error {
	// Acquire write lock on the database.
	db.Mu.Lock()
	defer db.Mu.Unlock()

	// Start the transaction
	tx, err := db.Conn.Begin()
	defer tx.Rollback()
	if err != nil {
		log.Printf("account/adduser.go: %s\n",
			"failed to begin transaction")
		return err
	}

	stmt, err := db.Conn.Prepare(`
INSERT INTO accounts(id, username, hash, regTime) values(?, ?, ?, ?)`)
	if err != nil {
		log.Printf("account/adduser.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.Username, u.Hash, time.Now().UTC())
	if err != nil {
		log.Printf("account/adduser.go: %s\n",
			"failed to execute statement")
		return err
	}

	tx.Commit()
	return err
}
