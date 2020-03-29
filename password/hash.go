package password

import "golang.org/x/crypto/bcrypt"

// Hash takes a string as input and returns the hash of the
// password.
func Hash(password string) (string, error) {
	// 10 is the default cost.
	out, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(out), err
}
