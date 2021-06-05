package table

import "errors"

// Possible errors for GetCoordinate() and GetInput() functions.
var (
	ErrInvalidCoordinate error = errors.New("Invalid coordinate")
	ErrInvalidInput      error = errors.New("Invalid input")
)
