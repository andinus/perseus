package user

import (
	"log"

	"tildegit.org/andinus/perseus/storage/sqlite3"
)

// GetID returns id from username.
func (u *User) GetID(db *sqlite3.DB) error {
	// Get password for this user from the database.
	stmt, err := db.Conn.Prepare("SELECT id FROM users WHERE username = ?")
	if err != nil {
		log.Printf("user/getid.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(u.username).Scan(&id)
	if err != nil {
		log.Printf("user/getid.go: %s\n",
			"query failed")
	}
	u.id = id

	return err
}
