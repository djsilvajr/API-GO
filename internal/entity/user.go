package entity

import (
	"github.com/devfullcycle/goexpert/9-APIS/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       entity.ID  `json:"id"`
	name     string     `json:"name"`
	email    string     `json:"email"`
	password string     `json:"-"`
}

func newUser(name, email, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		id:       entity.NewID(),
		name:     name,
		email:    email,
		password: string(hash),
	}, nil
}


func (u *User) validatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))

	return err == nil
}