package token

import (
	"errors"
	"log"

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
