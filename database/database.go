package database

import (
	Mongo "backend/database/mongo"
)

// DB represents database interactions.
// All databases can be listed here.
type DB struct {
	Mongo *Mongo.Server
}
