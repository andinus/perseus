package auth

import (
	"log"
	"strings"

	"tildegit.org/andinus/perseus/storage/sqlite3"
	"tildegit.org/andinus/perseus/user"
)

// Register takes in registration details and returns an error. If
// error doesn't equal nil then the registration was unsuccessful.
// regInfo should have username & password.
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

	err = u.AddUser(db)
	return err
}
