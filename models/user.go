package models

import (
	"errors"
	"github.com/jasurbekyuldashov/medhub_go/db"
	"github.com/jasurbekyuldashov/medhub_go/forms"

	"golang.org/x/crypto/bcrypt"
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

var authModel = new(AuthModel)

func (m UserModel) GetAll() (users []UserGet, err error) {
	err = db.GetDB().Select(&users, "SELECT id, email, full_name, created_at, updated_at, deleted_at from users")
	return users, err
}

// Login ...
func (m UserModel) Login(form forms.LoginForm) (user User, token Token, err error) {

	err = db.GetDB().Get(&user, "SELECT id, email, password, full_name, updated_at, created_at FROM public.users WHERE email=LOWER($1) LIMIT 1", form.Email)

	if err != nil {
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

// Register ...
func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	getDb := db.GetDB()

	var count = 0

	//Check if the user exists in database
	err1 := getDb.Get(&count, "SELECT count(id) FROM public.users WHERE email=LOWER($1) LIMIT 1", form.Email)
	if err1 != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	if count > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	//Create the user and return back the user ID
	err = getDb.QueryRow("INSERT INTO public.users(email, password, name) VALUES($1, $2, $3) RETURNING id", form.Email, string(hashedPassword), form.Name).Scan(&user.ID)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	//
	//user.Name = form.Name
	//user.Email = form.Email

	return user, err
}

// One ...
func (m UserModel) One(userID int64) (user User, err error) {
	err = db.GetDB().Get(&user, "SELECT id, email, name FROM public.user WHERE id=$1 LIMIT 1", userID)
	return user, err
}
