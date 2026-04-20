package organization

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sydesk/pkg/domain"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User) error
	GetAll(ctx context.Context) ([]*domain.User, error)
}

type UserServiceImpl struct {
	Repo     domain.UserRepo
	validate *validator.Validate
}

type UserAuth struct {
	repo      domain.UserRepo
	secretKey string
}

// Instancia para servicios comunes
func NewUserService(repo domain.UserRepo) UserService {
	return &UserServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

// instancia para servicio de autenticacion
func NewAuthService(repo domain.UserRepo, secret string) domain.AuthService {
	return &UserAuth{
		repo:      repo,
		secretKey: secret,
	}
}

// funcion para la autenticacion
func (s *UserAuth) Login(email, password string) (string, error) {
	// busqueda del usuario atravez del repo
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		// error generico para evitar vulnerabilidades
		return "", errors.New("Credenciales Invalidas")
	}

	// comparacion del hash en bd
	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD_HASH), []byte(password))
	if err != nil {
		return "", errors.New("Credenciales Invalidas")
	}

	// en caso de validacion exitosa
	claims := domain.CustomClaims{
		UserId: user.ID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))), // hasta el momento el tiempo de expiracion es de  1h
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-web-demo",
		},
	}

	// Creacion y firma del token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", errors.New("error al generar el token de acceso")
	}
	return signedToken, nil
}

func (s *UserAuth) ValidateToken(tokenStr string) (*domain.CustomClaims, error) {
	// parseo del token con los claims personalizados
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

func (s *UserServiceImpl) Create(ctx context.Context, user *domain.User) error {
	if err := s.validate.StructCtx(ctx, user); err != nil {
		log.Printf("Error al procesar la solicitud: %v", err)
		return fmt.Errorf("El mensaje no cumple los parametros definidos: %w", err)
	}

	// Llamar al repositorio para guardar
	err := s.Repo.Create(ctx, user)
	if err != nil {
		log.Printf("Error en repositorio: %v", err)
		return fmt.Errorf("no se pudo crear el registro")
	}

	return nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]*domain.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err)
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return users, nil
}
