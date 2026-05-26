package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	ID            string `json:"id" db:"id"`
	EMAIL         string `json:"email" db:"email"`
	IS_ACTIVE     bool   `json:"is_active" db:"is_active"`
	PASSWORD_HASH string `json:"password_hash" db:"password_hash"`

	// Recuperacion de data UX (despues agregamos el resto segun surja la necesidad)
	ALIAS   *string `json:"alias" db:"alias"`
	NAME    *string `json:"name" db:"name"`
	SURNAME *string `json:"surname" db:"surname"`
}

type CustomClaims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// interfaz de funcion
type AuthRepo interface {
	GetByEmail(email string) (*Auth, error)
}

type AuthServices interface {
	Login(email, password string) (string, error)
	ValidateToken(token string) (*CustomClaims, error)
}
