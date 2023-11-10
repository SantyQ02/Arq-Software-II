package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Active    bool      `json:"active,omitempty"`
}

type UserResponses []UserResponse
