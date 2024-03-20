package models

// Book struct to describe book object.
type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	PassWord  string `json:"pass_word"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
