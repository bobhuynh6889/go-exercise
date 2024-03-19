package models

import (
	"time"
)

// Book struct to describe book object.
type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	PassWord  string    `json:"pass_word"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
