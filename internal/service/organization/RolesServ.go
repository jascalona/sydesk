package organization

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain"

	"github.com/go-playground/validator/v10"
)

type RolesServices interface {
	GetAll(ctx context.Context) ([]*domain.Roles, error)
	Created(ctx context.Context, rol *domain.Roles) error
}

type RolesServiceImpl struct {
	Repo     domain.RolesRepo
	validate *validator.Validate
}

func NewRoleService(repo domain.RolesRepo) RolesServices {
	return &RolesServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *RolesServiceImpl) GetAll(ctx context.Context) ([]*domain.Roles, error) {

	roles, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return roles, nil
}

func (s *RolesServiceImpl) Created(ctx context.Context, rol *domain.Roles) error {
	// Invocamos la validacion del struct
	if err := s.validate.Struct(rol); err != nil {
		log.Printf("Error al procesar la solicitud %v:", err.Error())
		return fmt.Errorf("El mensaje no cumple con los parametros definidos: %v", err)
	}

	err := s.Repo.Created(ctx, rol)
	if err != nil {
		log.Printf("no se puede crear el registro", err.Error())
		return fmt.Errorf("No se pudo crear el registro, revisar los logs para mayor detalle")
	}
	return nil
}
