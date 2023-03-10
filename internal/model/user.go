package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"ID"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate string    `json:"birthDate"`
}

// type FIO struct {
// 	Name       string
// 	Surname    string
// 	Patronymic string
// }

func NewUser(name string, login string, password string, phone string, birthDate string) *User {
	return &User{
		ID:        uuid.New(),
		Name:      name,
		Login:     login,
		Password:  password,
		Phone:     phone,
		BirthDate: birthDate,
	}
}
