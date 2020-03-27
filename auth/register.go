package auth

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"tildegit.org/andinus/perseus/storage/sqlite3"
	"tildegit.org/andinus/perseus/user"
)

// Register takes in registration details and returns an error. If
// error doesn't equal nil then the registration was unsuccessful.
// uInfo should have username & password.
func Register(db *sqlite3.DB, uInfo map[string]string) error {
	u := user.User{}
	u.SetID(genID(64))
	u.SetUsername(strings.ToLower(uInfo["username"]))

	// Validate username
	re := regexp.MustCompile("^[a-z0-9]*$")
	if !re.MatchString(u.Username()) {
		return errors.New("auth/register.go: invalid username")
	}

	// Validate password
	if len(uInfo["password"]) < 8 {
		return errors.New("auth/register.go: password too short")
	}

	pass, err := hashPass(uInfo["password"])
	if err != nil {
		log.Printf("auth/register.go: %s\n",
			"hashPass func failed")
		return err
	}
	u.SetPassword(pass)

	err = u.AddUser(db)
	return err
}
