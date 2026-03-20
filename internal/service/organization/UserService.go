package organization

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User) error
	GetAll(ctx context.Context) ([]*domain.User, error)
}

type UserServiceImpl struct {
	Repo     domain.UserRepo
	validate *validator.Validate
}

func NewUserService(repo domain.UserRepo) UserService {
	return &UserServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
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
