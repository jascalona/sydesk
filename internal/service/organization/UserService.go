package organization

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain"

	"github.com/go-playground/validator/v10"
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

// Instancia para servicios comunes
func NewUserService(repo domain.UserRepo) UserService {
	return &UserServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, user *domain.User) error {
	// Invocamos la validacion del struct
	if err := s.validate.StructCtx(ctx, user); err != nil {
		log.Printf("Error al procesar la solicitud: %v", err)
		return fmt.Errorf("El mensaje no cumple los parametros definidos: %w")
	}

	// hasheo del password
	length := 10
	hash, errHash := bcrypt.GenerateFromPassword([]byte(user.PASSWORD_HASH), length)
	if errHash != nil {
		log.Printf("Error al hashear el password: %v", errHash)
		return fmt.Errorf("Error interno", errHash)
	}

	// remplazo de del texto plano por el hash y asignacion de rol por defecto
	user.PASSWORD_HASH = string(hash)
	user.ROLE_ID = 6 // ROLE DEFAULT READ ONLY

	// Persistencia de datos
	err := s.Repo.Create(ctx, user)
	if err != nil {
		log.Printf("Error al procesar la solicitud: %v", err)
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
