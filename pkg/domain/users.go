package domain

import (
	"context"
)

type User struct {
	ID             string `json:"id" db:"id"`
	ALIAS          string `json:"alias" db:"alias"`
	NAME           string `json:"name" db:"name"`
	SURNAME        string `json:"surname" db:"surname"`
	EMAIL          string `json:"email" db:"email"`
	PHONE          string `json:"phone" db:"phone"`
	PASSWORD_HASH  string `json:"password_hash" db:"password_hash"`
	DEPARTAMENT_ID int    `json:"department_id" db:"department_id"`
	ROLE_ID        int    `json:"role_id" db:"role_id"`
	IS_ACTIVE      bool   `json:"is_active" db:"is_active"`
	CREATED_AT     string `json:"created_at" db:"created_at"`
}

type UserValidation struct {
	ID             string `json:id validate:"required,max=15"`
	NAME           string `json:validate:"required"`
	SURNAME        string `json:surname:"required"`
	EMAIL          string `json:email:"required,email"`
	PHONE          string `json:phone:"required, min=11,max=11"`
	DEPARTAMENT_ID int    `json:department_id:"required"`
	ROLE_ID        int    `json:role_id:"required"`
	IS_ACTIVE      bool   `json:is_active:"required"`
}

type UserRepo interface {
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, users *User) error
}
