package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"
)

type CustomProductRoleService interface {
	GetAll(ctx context.Context, role_id int) ([]*components.CustomerProductRole, error)
	Created(ctx context.Context, cpr *components.CustomerProductRole) error
}

type CustomerProductRoleImpl struct {
	Repo components.CustomerProductRoleRepo
}

func NewCustomProductRoleService(repo components.CustomerProductRoleRepo) CustomProductRoleService {
	return &CustomerProductRoleImpl{
		Repo: repo,
	}
}

func (s *CustomerProductRoleImpl) GetAll(ctx context.Context, role_id int) ([]*components.CustomerProductRole, error) {

	if role_id != 0 {
		customPR, err := s.Repo.GetAll(ctx, role_id)
		if err != nil {
			log.Printf("Error al obtener los registros: %v", err.Error())
			return nil, fmt.Errorf("error al obtener los registros")
		}

		return customPR, nil
	}
	log.Printf("el role_id no puede quedar vacio")
	return nil, fmt.Errorf("el role_id no puede quedar vacio")

}

func (s *CustomerProductRoleImpl) Created(ctx context.Context, cpr *components.CustomerProductRole) error {

	// persistencia en bd
	err := s.Repo.Created(ctx, cpr)
	if err != nil {
		log.Printf("error al crear la solicitud: %v", err.Error())
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil
}
