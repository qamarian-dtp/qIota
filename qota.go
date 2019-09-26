package qota

// Always create new Quotas with this function.
func New () (*Qota) {
	return &Qota {0}
}

// Always create data of this type with New ().
type Qota struct {
	value uint8
}

// Value () returns the current value of the qota.
func (q *Qota) Value () (uint8) {
	return q.value
}

// Grow () returns true if the qota was successfully grown. Otherwise, it returns false.
func (q *Qota) Grow () (bool) {
	if q.value >= 255 {
		return false
	}
	q.value ++
	return true
}
