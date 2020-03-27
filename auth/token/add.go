package token

import (
	"log"
	"time"

	"tildegit.org/andinus/perseus/storage/sqlite3"
	"tildegit.org/andinus/perseus/user"
)

// AddToken will generate a random token, add it to database and
// return the token.
func AddToken(db *sqlite3.DB, uInfo map[string]string) (token string, err error) {
	// Acquire write lock on the database.
	db.Mu.Lock()
	defer db.Mu.Unlock()

	token = genToken(64)

	u := user.User{}
	u.SetUsername(uInfo["username"])

	// Set user id from username.
	err = u.GetID(db)
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"failed to get id from username")
		return
	}

	// Start the transaction
	tx, err := db.Conn.Begin()
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"failed to begin transaction")
		return
	}

	stmt, err := db.Conn.Prepare(`
INSERT INTO access(id, token, genTime) values(?, ?, ?)`)
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"failed to prepare statement")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID(), u.Username(), time.Now().UTC())
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"failed to execute statement")
		return
	}

	tx.Commit()
	return

}
