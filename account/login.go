package account

import (
	"log"

	"tildegit.org/andinus/perseus/password"
	"tildegit.org/andinus/perseus/storage"
)

// Login takes in login details and returns an error. If error doesn't
// equal nil then consider login failed. It will also set the u.Token
// field.
func (u *User) Login(db *storage.DB) error {
	// Acquire read lock on the database.
	db.Mu.RLock()

	// Get password for this user from the database.
	stmt, err := db.Conn.Prepare("SELECT hash FROM accounts WHERE username = ?")
	if err != nil {
		log.Printf("account/login.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	var hash string
	err = stmt.QueryRow(u.Username).Scan(&hash)
	if err != nil {
		log.Printf("account/login.go: %s\n",
			"query failed")
		return err
	}
	u.Hash = hash

	// Check user's password.
	err = password.Check(u.Password, u.Hash)
	if err != nil {
		log.Printf("account/login.go: %s%s\n",
			"user login failed, username: ", u.Username)
		return err
	}
	db.Mu.RUnlock()

	err = u.addToken(db)
	if err != nil {
		log.Printf("account/login.go: %s\n",
			"addtoken failed")
	}
	return err
}
