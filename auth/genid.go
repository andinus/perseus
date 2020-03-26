package auth

import (
	"crypto/rand"
	"encoding/base64"
)

// genID generates a random id string of length n. Don't forget to
// seed the random number generator otherwise it won't be random.
func genID(n int) string {
	b := make([]byte, n/2)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
