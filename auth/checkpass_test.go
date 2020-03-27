package auth

import "testing"

// TestCheckPass tests the checkPass function.
func TestCheckPass(t *testing.T) {
	var err error
	var passhash = make(map[string]string)

	// First we check with static values, these should always pass.
	passhash["ter4cQ=="] = "$2a$10$ANkaNEFFQ4zxDwTwvAUfoOCqpVIdgtPFopFOTMSrFy39WkaMAYLIC"
	passhash["G29J6A=="] = "$2a$10$1oH1PyhncIHcHJWbLt3Gv.OjClUoFoaEDaFpQ9E9atfRbIrVxsbwm"
	passhash["Z1S/kQ=="] = "$2a$10$fZ05kKmb7bh4vBLebpK1u.3bUNQ6eeX5ghT/GZaekgS.5bx4.Ru1e"
	passhash["J861dQ=="] = "$2a$10$nXb6Btn6n3AWMAUkDh9bFObvQw5V9FLKhfX.E1EzRWgVDuqIp99u2"

	// We also check with values generated with hashPass, this may
	// fail if hashPass itself fails in that case it's not
	// checkPass error so the test shouldn't fail but warning
	// should be sent. We use genID func to generate random inputs
	// for this test.
	for i := 1; i <= 4; i++ {
		p := genID(8)
		passhash[p], err = hashPass(p)
		if err != nil {
			t.Log("hashPass func failed")
		}
	}

	// We test the checkPass func by ranging over all values of
	// passhash.
	for p, h := range passhash {
		err := checkPass(p, h)
		if err != nil {
			t.Errorf("password: %s, hash: %s didn't match.",
				p, h)
		}
	}

}
