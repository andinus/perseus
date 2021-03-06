// Password package contains functions related to passwords.
package password

import "golang.org/x/crypto/bcrypt"

// Check takes a string and hash as input and returns an error. If
// the error is not nil then the consider the password wrong. We're
// returning error instead of a bool so that we can print failed
// logins to log and logging shouldn't happen here.
func Check(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
