package models

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID
	Username string
	Email    string
	Password string
}
