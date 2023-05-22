package Mongo

import (
	"backend/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserBeforeMarshal(t *testing.T) {
	user := &models.User{
		Name:         "test1",
		SurName:      "test1",
		Nation_Id:    "12345678910",
		Phone_Number: "+901234567890",
	}
	_ = UpdateUserBeforeMarshal(user)
	assert.False(t, user.CreatedAt.IsZero())
	assert.NotEqual(t, uuid.Nil, user.Id)
	assert.False(t, user.UpdatedAt.IsZero())
}
