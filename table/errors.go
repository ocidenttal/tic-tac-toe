// possible errors that
// can occur are declared
// here
package table

import "errors"

// Errors declaration
var (
    InvalidCoordinate error = errors.New("Invalid coordinate")
    InvalidInput error = errors.New("Invalid input")
)

