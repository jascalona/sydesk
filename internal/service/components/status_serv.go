package components

import (
	"context"
	"fmt"
	"log"
	domain "sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type StatusService interface {
	GetAll(ctx context.Context) ([]*domain.Status, error)
}

type StatusServImpl struct {
	Repo     domain.StatusRepo
	validate *validator.Validate
}

// Instancia para servicios comunes
func NewStatusService(repo domain.StatusRepo) StatusService {
	return &StatusServImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *StatusServImpl) GetAll(ctx context.Context) ([]*domain.Status, error) {
	status, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err)
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return status, nil
}
