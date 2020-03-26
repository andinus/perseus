package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// hashPass takes a string as input and returns the hash of the
// password.
func hashPass(password string) (string, error) {
	// 10 is the default cost.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
