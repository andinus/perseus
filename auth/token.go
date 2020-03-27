package auth

import (
	"errors"
	"log"
	"time"

	"tildegit.org/andinus/perseus/storage/sqlite3"
	"tildegit.org/andinus/perseus/user"
)

// ValToken will validate the token and returns an error. If error
// doesn't equal nil then consider token invalid.
func ValToken(db *sqlite3.DB, uInfo map[string]string) error {
	// Acquire read lock on the database.
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	u := user.User{}
	u.SetUsername(uInfo["username"])

	// Set user id from username.
	err := u.GetID(db)
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"failed to get id from username")
		return err
	}

	// Check if user's token is valid.
	stmt, err := db.Conn.Prepare("SELECT token FROM access WHERE id = ?")
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	var token string
	err = stmt.QueryRow(u.ID()).Scan(&token)
	if err != nil {
		log.Printf("auth/token.go: %s\n",
			"query failed")
		return err
	}

	if token != uInfo["token"] {
		err = errors.New("token mismatch")
	}

	return err
}

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
INSERT INTO access(id, username, genTime) values(?, ?, ?)`)
	if err != nil {
		log.Printf("auth/tokenr.go: %s\n",
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
