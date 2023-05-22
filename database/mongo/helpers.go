package Mongo

import (
	"backend/models"
	"time"

	"github.com/google/uuid"
)

// UpdateUserBeforeMarshal initialize User values before created
func UpdateUserBeforeMarshal(u *models.User) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	if u.Id == uuid.Nil {
		u.Id = uuid.New()
	}
	u.UpdatedAt = time.Now()
	return nil
}
