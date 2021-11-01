package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" form:"id"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
}
