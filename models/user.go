package models

import (
	"github.com/jasurbekyuldashov/medhub_go/db"
)

type User struct {
	Base
	FullName string `db:"full_name" json:"fullName"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserGet struct {
	Base
	FullName string `db:"full_name" json:"full_name"`
	Email    string `db:"email" json:"email"`
}

type UserModel struct{}

func (m UserModel) GetAll() (users []UserGet, err error) {
	err = db.GetDB().Select(&users, "SELECT id, email, full_name, created_at, updated_at, deleted_at from users")
	return users, err
}
