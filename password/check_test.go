package password

import "testing"

// TestCheck tests the Check function.
func TestCheck(t *testing.T) {
	var err error
	passhash := make(map[string]string)

	// First we check with static values, these should always pass.
	passhash["ter4cQ=="] = "$2a$10$ANkaNEFFQ4zxDwTwvAUfoOCqpVIdgtPFopFOTMSrFy39WkaMAYLIC"
	passhash["G29J6A=="] = "$2a$10$1oH1PyhncIHcHJWbLt3Gv.OjClUoFoaEDaFpQ9E9atfRbIrVxsbwm"
	passhash["Z1S/kQ=="] = "$2a$10$fZ05kKmb7bh4vBLebpK1u.3bUNQ6eeX5ghT/GZaekgS.5bx4.Ru1e"
	passhash["J861dQ=="] = "$2a$10$nXb6Btn6n3AWMAUkDh9bFObvQw5V9FLKhfX.E1EzRWgVDuqIp99u2"

	// We also check with values generated with Hash, this may
	// fail if Hash itself fails in that case it's not Check error
	// so the test shouldn't fail but warning should be sent. We
	// use genID func to generate random inputs for this test.
	for i := 1; i <= 4; i++ {
		p := RandStr(8)
		passhash[p], err = Hash(p)
		if err != nil {
			t.Log("hashPass func failed")
		}
	}

	// We test the Check func by ranging over all values of
	// passhash. We assume that Hash func returns correct hashes.
	for p, h := range passhash {
		err = Check(p, h)
		if err != nil {
			t.Errorf("password: %s, hash: %s didn't match.",
				p, h)
		}
	}

}
