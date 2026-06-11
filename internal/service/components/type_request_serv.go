package components

import (
	"context"
	"fmt"
	"log"
	domain "sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type TypeRequestServ interface {
	GetAll(ctx context.Context) ([]*domain.TypeRequest, error)
}

type TypeRequestServImpl struct {
	Repo     domain.InterfaceTypeRequest
	validate *validator.Validate
}

// Instanacias para servicios comunes
func NewTypeRequestService(repo domain.InterfaceTypeRequest) TypeRequestServ {
	return &TypeRequestServImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *TypeRequestServImpl) GetAll(ctx context.Context) ([]*domain.TypeRequest, error) {

	typeRequest, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Println("Error al obtener los registros: ", err)
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return typeRequest, nil
}
