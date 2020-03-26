package auth

import (
	"testing"
)

// TestCheckPass tests the checkPass function.
func TestCheckPass(t *testing.T) {
	var passhash = make(map[string]string)
	passhash["password"] = "$2a$10$hyV9vtsYXX88wz1rmA1x0.tcdkyvd6QsmV6gLOcR5wtYBE2GaSqT."

	for p, h := range passhash {
		err := checkPass(p, h)
		if err != nil {
			t.Errorf("password: %s, hash: %s didn't match.",
				p, h)
		}
	}
}
