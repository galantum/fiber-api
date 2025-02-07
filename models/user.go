package models

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"user_name" json:"user_name"`
	Email string `db:"user_email" json:"user_email"`
}
