package password

import "testing"

// TestHash tests the Hash function.
func TestHash(t *testing.T) {
	var err error
	passhash := make(map[string]string)

	// We generate random hashes with Hash, random string is
	// generate by RandStr func.
	for i := 1; i <= 8; i++ {
		p := RandStr(8)
		passhash[p], err = Hash(p)

		// Here we test if the hashPass func runs sucessfully.
		if err != nil {
			t.Errorf("Hash func failed for password: %s",
				p)
		}
	}

	// Here we are testing if the hashPass func returns correct
	// hashes. We assume that checkPass func returns correct
	// values.
	for p, h := range passhash {
		err = Check(p, h)
		if err != nil {
			t.Errorf("password: %s, hash: %s didn't match.",
				p, h)
		}
	}

}
