package business

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/business"
)

type ContactService interface {
	ContactByCustomerID(ctx context.Context, customerID int) ([]*business.Contact, error)
	Created(ctx context.Context, contact *business.Contact) error
}

type ContactServImpl struct {
	Repo business.InterfaceContact
}

func NewContactService(repo business.InterfaceContact) ContactService {
	return &ContactServImpl{
		Repo: repo,
	}
}

func (s *ContactServImpl) ContactByCustomerID(ctx context.Context, customerID int) ([]*business.Contact, error) {
	contact, err := s.Repo.ContactByCustomerID(ctx, customerID)
	if err != nil {
		log.Printf("error al obtener los registros %v")
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return contact, nil
}

func (s *ContactServImpl) Created(ctx context.Context, contact *business.Contact) error {

	// persistencia de data
	err := s.Repo.Created(ctx, contact)
	if err != nil {
		log.Printf("error al procesar la solicitud %v", err.Error())
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil

}
