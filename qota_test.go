package qota

import (
	"gopkg.in/qamarian-lib/str.v2"
	"math/big"
	"testing"
)

// TestQota () tests all aspect of the Qota data type.
func TestQota (t *testing.T) {
	str.PrintEtr ("Test has started!", "std", "TestQota ()")

	q := New ()

	for index := big.NewInt (0); index.Cmp (big.NewInt (1000000)) != 0; index.Add (index, big.NewInt (1)) {
		value := q.Value ()

		if index.Cmp (value) != 0 {
			str.PrintEtr ("Test failed! Ref: 0", "err", "TestQota ()")
			t.FailNow ()
		}
	}

	str.PrintEtr ("Test passed!", "std", "TestQota ()")
}
