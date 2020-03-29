package password

import (
	"crypto/rand"
	"encoding/base64"
)

// RandStr will return a random base64 encoded string of length n.
func RandStr(n int) string {
	b := make([]byte, n/2)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
