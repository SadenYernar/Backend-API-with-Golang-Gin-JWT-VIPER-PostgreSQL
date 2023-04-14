package model

import "github.com/gofrs/uuid"

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
