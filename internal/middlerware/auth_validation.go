package middlerware

import (
	"errors"
	"fmt"
	"log"
	domain "sydesk/pkg/domain/organization"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserAuth struct {
	repo      domain.AuthRepo
	secretKey string
}

// Instancia para la autenticacion
func NewAuthService(repo domain.AuthRepo, secret string) domain.AuthServices {
	return &UserAuth{
		repo:      repo,
		secretKey: secret,
	}
}

// Funcion para validar la autenticacion
func (s *UserAuth) Login(email, password string) (string, error) {

	// validar si el usuario existe
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		// error generico para vulnerabilidades
		log.Println("Error search user", err.Error())
		return "", errors.New("Credenciales Invalidas")
	}

	// validacion del password_hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD_HASH)[:], []byte(password))
	if err != nil {
		log.Println("Error al decifrar credenciales", err.Error())
		return "", errors.New("Credenciales Invalidas")
	}

	// nueva validacion la la logica de registros (usuarios no verificados)
	if user.IS_ACTIVE != true {
		log.Println("Lo sentimos este usuario no ha sido verificado", err)
		return "", fmt.Errorf("Usuario no ha sido verificado")
	}

	// en caso de validacion exitosa
	claims := domain.CustomClaims{
		UserId: user.ID,
		Email:  email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(5))), // hasta el momento el tiempo de expiracion es de  1h
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-web-demo",
		},
	}

	// creacion y firma del token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Println("Error signing token", err.Error())
		return "", errors.New("Error al generar el token de acceso")
	}

	return signedToken, nil
}

func (s *UserAuth) ValidateToken(tokenStr string) (*domain.CustomClaims, error) {
	// parseo del token
	token, err := jwt.ParseWithClaims(tokenStr, &domain.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		// validacion del metodo de firma antes de retornar la llave
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metodo de firma inesperado")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// extraccion de clamis en caso de que el token sea valido
	if clamis, ok := token.Claims.(*domain.CustomClaims); ok && token.Valid {
		return clamis, nil
	}
	return nil, errors.New("Token Invalido")
}
