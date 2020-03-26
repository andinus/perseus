package auth

import (
	"log"
	"strings"
	"time"

	"tildegit.org/andinus/perseus/storage/sqlite3"
	"tildegit.org/andinus/perseus/user"
)

// Register takes in registration details and returns an error. If
// error doesn't equal nil then the registration was unsuccessful.
// regInfo should have username, password & ip.
func Register(db *sqlite3.DB, regInfo map[string]string) error {
	u := user.User{}
	u.SetID(genID(64))
	u.SetUsername(strings.ToLower(regInfo["username"]))

	pass, err := hashPass(regInfo["password"])
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"hashPass func failed")
		return err
	}
	u.SetPassword(pass)

	// Acquire write lock on the database.
	db.Mu.Lock()
	defer db.Mu.Unlock()

	err = insertRecords(db, u, regInfo)
	return err
}

func insertRecords(db *sqlite3.DB, u user.User, regInfo map[string]string) error {
	// Start the transaction
	tx, err := db.Conn.Begin()
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"Failed to begin transaction")
		return err
	}

	// Insert the record into registration table
	regStmt, err := db.Conn.Prepare(`
INSERT INTO registration(id, username, reg_time, reg_ip) values(?, ?, ?, ?)`)
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"Failed to prepare statement")
		return err
	}
	defer regStmt.Close()

	_, err = regStmt.Exec(u.ID(), u.Username(), time.Now().UTC(), regInfo["ip"])
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"Failed to execute statement")
		return err
	}

	// Insert the record into users table
	usrStmt, err := db.Conn.Prepare(`
INSERT INTO users(id, username, password) values(?, ?, ?)`)
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"Failed to prepare statement")
		return err
	}
	defer usrStmt.Close()

	_, err = usrStmt.Exec(u.ID(), u.Username(), u.Password())
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"Failed to execute statement")
		return err
	}

	tx.Commit()
	return err
}
