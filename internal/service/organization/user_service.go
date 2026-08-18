package organization

import (
	"context"
	"fmt"
	"log"
	domain "sydesk/pkg/domain/organization"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User) error
	GetAll(ctx context.Context) ([]*domain.User, error)
}

type UserServiceImpl struct {
	Repo domain.UserRepo
}

// Instancia para servicios comunes
func NewUserService(repo domain.UserRepo) UserService {
	return &UserServiceImpl{
		Repo: repo,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, user *domain.User) error {

	// validaciones de negocio si son requeridas mas adelantes

	// hasheo del password
	length := 10
	hash, errHash := bcrypt.GenerateFromPassword([]byte(user.PASSWORD_HASH), length)
	if errHash != nil {
		log.Printf("Error al hashear el password: %v", errHash)
		return fmt.Errorf("Error interno %v", errHash)
	}

	// remplazo de del texto plano por el hash y asignacion de rol por defecto
	user.PASSWORD_HASH = string(hash)
	user.ROLE_ID = 27 // ROLE DEFAULT READ ONLY

	// HAY QUE DEFINIR LA LOGICA PARA LA ACTIVACION DE LA CUENTA

	user.IS_ACTIVE = false // IS_ACTIVE FASE DEFAULT

	// Persistencia de datos
	err := s.Repo.Create(ctx, user)
	if err != nil {
		log.Printf("Error al procesar la solicitud: %v", err)
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
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
