package models

type User struct {
	Base
	FullName string `db:"full_name" json:"fullName"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
