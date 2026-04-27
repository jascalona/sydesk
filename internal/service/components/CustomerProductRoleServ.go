package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type CustomProductRoleService interface {
	GetAll(ctx context.Context) ([]*components.CustomerProductRole, error)
}

type CustomerProductRoleImpl struct {
	Repo     components.CustomerProductRoleRepo
	validate *validator.Validate
}

func NewCustomProductRoleService(repo components.CustomerProductRoleRepo) CustomProductRoleService {
	return &CustomerProductRoleImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *CustomerProductRoleImpl) GetAll(ctx context.Context) ([]*components.CustomerProductRole, error) {
	customPR, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}

	return customPR, nil
}
