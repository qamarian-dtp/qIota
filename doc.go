// Package qota (Qamarian iota) implements the iota data type -- a data type which works
// like the Golang "iota".
//
// 	i := qota.New () // This is how qotas should always be created.
// 	i.Value () // Returns 0.
// 	i.Value () // Returns 1.
// 	i.Value () // Returns 2.
//
package qota
