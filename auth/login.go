package auth

import (
	"log"

	"tildegit.org/andinus/perseus/storage/sqlite3"
	"tildegit.org/andinus/perseus/user"
)

// Login takes in login details and returns an error. If error doesn't
// equal nil then consider login failed.
func Login(db *sqlite3.DB, uInfo map[string]string) error {
	// Acquire read lock on the database.
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	u := user.User{}
	u.SetUsername(uInfo["username"])

	// Get password for this user from the database.
	stmt, err := db.Conn.Prepare("SELECT password FROM users WHERE username = ?")
	if err != nil {
		log.Printf("auth/login.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	var pass string
	err = stmt.QueryRow(u.Username()).Scan(&pass)
	if err != nil {
		log.Printf("auth/login.go: %s\n",
			"query failed")
		return err
	}
	u.SetPassword(pass)

	// Check user's password.
	err = checkPass(uInfo["password"], u.Password())
	if err != nil {
		log.Printf("auth/login.go: %s%s\n",
			"user login failed, username: ", u.Username())
	}

	return err
}
