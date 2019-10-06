package qota

import (
	"math/big"
)

// Always create new Quotas with this function.
func New () (*Qota) {
	return &Qota {big.NewInt (0)}
}

// Always create data of this type with New ().
type Qota struct {
	value *big.Int
}

// Value () returns true if the qota was successfully grown. Otherwise, it returns false.
func (q *Qota) Value () (*big.Int) {
	q.value.Add (q.value, big.NewInt (1))

	output := big.NewInt (-1)
	return output.Add (output, q.value)
}
