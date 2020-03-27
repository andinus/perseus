package auth

// genToken generates a random token string of length n. Don't forget to
// seed the random number generator otherwise it won't be random.
func genToken(n int) string {
	// Currently this is just a wrapper to genID.
	return genID(n)
}
