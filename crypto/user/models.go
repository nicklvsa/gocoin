package user

import "github.com/gofrs/uuid"

type User struct {
	UserID uuid.UUID
	Email  string `json:"email"`
}
