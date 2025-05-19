package database

import "errors"

var (
	ErrNotConnectedDB = errors.New("Could'n connected to the database")
)
