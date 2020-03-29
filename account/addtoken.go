package account

import (
	"log"
	"time"

	"tildegit.org/andinus/perseus/password"
	"tildegit.org/andinus/perseus/storage"
)

// addToken will generate a random token, add it to database and
// return the token.
func (u *User) addToken(db *storage.DB) error {
	u.Token = password.RandStr(64)

	// Set user id from username.
	err := u.GetID(db)
	if err != nil {
		log.Printf("account/addtoken.go: %s\n",
			"failed to get id from username")
		return err
	}

	// Acquire write lock on the database.
	db.Mu.Lock()
	defer db.Mu.Unlock()

	// Start the transaction
	tx, err := db.Conn.Begin()
	defer tx.Rollback()
	if err != nil {
		log.Printf("account/addtoken.go: %s\n",
			"failed to begin transaction")
		return err
	}

	stmt, err := db.Conn.Prepare(`
INSERT INTO access(id, token, genTime) values(?, ?, ?)`)
	if err != nil {
		log.Printf("account/addtoken.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.Token, time.Now().UTC())
	if err != nil {
		log.Printf("account/addtoken.go: %s\n",
			"failed to execute statement")
		return err
	}

	tx.Commit()
	return err

}
