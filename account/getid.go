package account

import (
	"log"

	"tildegit.org/andinus/perseus/storage"
)

// GetID returns id from username.
func (u *User) GetID(db *storage.DB) error {
	// Acquire read lock on database.
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	// Get password for this user from the database.
	stmt, err := db.Conn.Prepare("SELECT id FROM accounts WHERE username = ?")
	if err != nil {
		log.Printf("account/getid.go: %s\n",
			"failed to prepare statement")
		return err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(u.Username).Scan(&id)
	if err != nil {
		log.Printf("account/getid.go: %s\n",
			"query failed")
	}
	u.ID = id

	return err
}
