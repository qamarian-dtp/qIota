package qota

import (
	"gopkg.in/qamarian-lib/str.v2"
	"testing"
)

// TestQota test all aspect of the Qota data type.
func TestQota (t *testing.T) {
	q := New ()

	for index := 0; index < 512; index ++ {
		// Testing if Qota starts with its expected value, and if it also grows
		// as it is supposed to. { ...
		if (uint8 (index) != q.Value ()) && index < 256 {
			str.PrintEtr ("Test failed! Problem with qota's value.", "err",
				"TestQota ()")
			t.FailNow ()
		}
		// ... }

		// Testing if Qota did not overgrow or undergrow, after reaching its
		// limit. { ...
		if index > 255 && q.Value () != 255 {
			str.PrintEtr ("Test failed! Qota overgrew or undergrew, after" +
				" reaching it limit.", "err", "TestQota ()")
			t.FailNow ()
		}
		// ... }

		grown := q.Grow ()

		// Testing if Qota is not claimed to have grown, after reching its
		// limit. { ...
		if index > 254 && grown == true {
			str.PrintEtr ("Test failed! Qota claimed to have grown to a " +
				"value over 255.", "err", "TestQota ()")
			t.FailNow ()
		}
		// ... }
	}

	str.PrintEtr ("Test passed!", "std", "TestQota ()")
}
