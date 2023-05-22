package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	// Unexported fields
	Id        uuid.UUID  `json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
