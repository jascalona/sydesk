package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type SubcomponentService interface {
	GetAll(ctx context.Context) ([]*components.Subcomponents, error)
}

type SubcomponentImpl struct {
	Repo     components.SubcomponentsRepo
	validate *validator.Validate
}

// Instancia para servicios comunes
func NewSubcomponentService(repo components.SubcomponentsRepo) SubcomponentService {
	return &SubcomponentImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *SubcomponentImpl) GetAll(ctx context.Context) ([]*components.Subcomponents, error) {
	subcomponents, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return subcomponents, nil
}

// falta implementar el repo y validacion de este servicio
