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
	DEPARTAMENT_ID int    `json:"departament_id" db:"departament_id"`
	ROLE_ID        int    `json:"role_id" db:"role_id"`
	IS_ACTIVE      bool   `json:"is_active" db:"is_active"`
	CREATED_AT     string `json:"created_at" db:"created_at"`
}

type ValidationUser struct {
	ID             string `json:"id" binding:"required,max=15,min=6"`
	ALIAS          string `json:"alias" binding:"required,max=100,min=2"`
	NAME           string `json:"name" binding:"required,max=100,min=2"`
	SURNAME        string `json:"surname" binding:"required,max=100,min=2"`
	EMAIL          string `json:"email" binding:"required,email"`
	PHONE          string `json:"phone" binding:"required,min=11,max=11"`
	PASSWORD_HASH  string `json:"password_hash" binding:"required"`
	DEPARTAMENT_ID int    `json:"departament_id" binding:"required"`
	ROLE_ID        int    `json:"role_id" binding:"required"`
	//IS_ACTIVE      bool   `json:"is_active" binding:"required"`
}

type UserRepo interface {
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, users *User) error
}
