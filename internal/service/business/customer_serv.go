package business

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/business"
)

type CustomerServ interface {
	GetAll(ctx context.Context) ([]*business.Customer, error)
	Created(ctx context.Context, customer *business.Customer) error
}

type customerServImpl struct {
	Repo business.CustomerRepo
}

func NewCustomerServ(repo business.CustomerRepo) CustomerServ {
	return &customerServImpl{
		Repo: repo,
	}
}

func (s *customerServImpl) GetAll(ctx context.Context) ([]*business.Customer, error) {
	custom, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return custom, nil
}

func (s *customerServImpl) Created(ctx context.Context, customer *business.Customer) error {
	// Validaciones de negocio se son requeridas mas adelantes

	// persistencia de datos
	err := s.Repo.Created(ctx, customer)
	if err != nil {
		log.Printf("Error al crear la solicitud: %v", err)
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil
}
