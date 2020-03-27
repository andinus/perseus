package auth

import "testing"

// TestHashPass tests the checkPass function.
func TestHashPass(t *testing.T) {
	var err error
	passhash := make(map[string]string)

	// We generate random hashes with hashPass, random string is
	// generate by genID func.
	for i := 1; i <= 8; i++ {
		p := genID(8)
		passhash[p], err = hashPass(p)

		// Here we test if the hashPass func runs sucessfully.
		if err != nil {
			t.Errorf("hashPass func failed for password: %s",
				p)
		}
	}

	// Here we are testing if the hashPass func returns correct
	// hashes. We assume that checkPass func returns correct
	// values.
	for p, h := range passhash {
		err = checkPass(p, h)
		if err != nil {
			t.Errorf("password: %s, hash: %s didn't match.",
				p, h)
		}
	}

}
