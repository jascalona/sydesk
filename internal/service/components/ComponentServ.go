package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type ComponentService interface {
	GetAll(ctx context.Context) ([]*components.Components, error)
	Created(ctx context.Context, components *components.Components) error
}

type ComponentServiceImpl struct {
	Repo     components.ComponentsRepo
	validate *validator.Validate
}

// instancia servicios comunes
func NewComponentService(repo components.ComponentsRepo) ComponentService {
	return &ComponentServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *ComponentServiceImpl) GetAll(ctx context.Context) ([]*components.Components, error) {
	componentes, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return componentes, nil
}

func (s *ComponentServiceImpl) Created(ctx context.Context, components *components.Components) error {
	// validacion del struc
	if err := s.validate.StructCtx(ctx, components); err != nil {
		log.Printf("Error al procesar la solicitud: %v", err.Error()) // detalles reservados para la rest_api
		return fmt.Errorf("El mensaje no cumple los parametros definidos: %w")
	}

	components.IS_ACTIVE = 0 == 0

	// persistencia en bd
	err := s.Repo.Created(ctx, components)
	if err != nil {
		log.Printf("Error al procesar la solicitud: %v", err.Error())
		return fmt.Errorf("no se pudo crear el registro")
	}
	return nil
}
