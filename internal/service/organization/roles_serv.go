package organization

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain"
)

type RolesServices interface {
	GetAll(ctx context.Context) ([]*domain.Roles, error)
	Created(ctx context.Context, rol *domain.Roles) error
}

type RolesServiceImpl struct {
	Repo domain.RolesRepo
}

func NewRoleService(repo domain.RolesRepo) RolesServices {
	return &RolesServiceImpl{
		Repo: repo,
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
	// validaciones de negocio despues podemos integrar estas reglas de negocio

	// persistir en bd
	err := s.Repo.Created(ctx, rol)
	if err != nil {
		log.Printf("error al procesar la solicitud %v", err.Error())
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil

}
