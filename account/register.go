package account

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"tildegit.org/andinus/perseus/password"
	"tildegit.org/andinus/perseus/storage"
)

// Register takes in registration details and returns an error. If
// error doesn't equal nil then the registration was unsuccessful.
func (u User) Register(db *storage.DB) error {
	var err error
	u.ID = password.RandStr(64)
	u.Username = strings.ToLower(u.Username)

	// Validate username. It must be alphanumeric and less than
	// 128 characters.
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	if !re.MatchString(u.Username) {
		return errors.New("account/register.go: invalid username")
	}
	if len(u.Username) > 128 {
		return errors.New("account/register.go: username too long")
	}

	// Validate password
	if len(u.Password) < 8 {
		return errors.New("account/register.go: password too short")
	}

	u.Hash, err = password.Hash(u.Password)
	if err != nil {
		log.Printf("account/register.go: %s\n",
			"password.Hash func failed")
		return err
	}

	err = u.addUser(db)
	if err != nil {
		log.Printf("account/register.go: %s\n",
			"addUser func failed")
	}
	return err

}
