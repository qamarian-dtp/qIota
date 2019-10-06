// Package (qota - Qamarian iota) implements the iota data type (a data type which works
// like the Golang "iota").
//
// 	someQota := qota.New () // This is how qotas should always be created.
// 	someQota.Value () // Returns 0.
// 	someQota.Value () // Returns 1.
// 	someQota.Value () // Returns 2.
//
package qota
