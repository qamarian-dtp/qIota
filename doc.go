// This package (qota - Qamarian iOTA ) implements the iota data type (a data type which
// is like the Golang iota). When you create a new qota, it is automatically assigned a
// value of 0. However, you could grow the qota by 1, with every call to its method
// Grow ().
//
// 	someQota := qota.New () // This is how qotas should always be created.
// 	someQota.Value () // Returns 0.
//	someQota.Grow ()
// 	someQota.Value () // Returns 1.
// 	someQota.Grow ()
// 	someQota.Value () // Returns 2.
//
package qota
