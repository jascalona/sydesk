package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type CustomRoleService interface {
	GetAll(ctx context.Context) ([]*components.CustomerRole, error)
}

type CustomRoleServiceImpl struct {
	Repo     components.CustomRoleRepo
	validate *validator.Validate
}

// Instancia para del repo para los servicios comunes
func NewCustomRoleService(repo components.CustomRoleRepo) CustomRoleService {
	return &CustomRoleServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *CustomRoleServiceImpl) GetAll(ctx context.Context) ([]*components.CustomerRole, error) {
	customRole, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return customRole, nil
}
