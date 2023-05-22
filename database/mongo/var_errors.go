package Mongo

import "errors"

var (
	ErrFind               = errors.New("failed to find data")
	ErrCursorAll          = errors.New("an error occurred while retrieving data from the database")
	ErrNotFoundResult     = errors.New("no data found")
	ErrUpdateOne          = errors.New("failed to update data")
	ErrCarAlreadyReserved = errors.New("the car is already reserved")
	ErrNoActiveOffices    = errors.New("no active offices available")
)
