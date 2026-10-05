package components

import (
	"context"
	"log"
	"sydesk/pkg/domain/components"
)

type DepartamentService interface {
	GetAll(ctx context.Context) ([]*components.Departament, error)
}

type departamentServiceImpl struct {
	repo components.InterfaceDepartament
}

func NewDepartamentService(repo components.InterfaceDepartament) DepartamentService {
	return &departamentServiceImpl{
		repo: repo,
	}
}

func (s *departamentServiceImpl) GetAll(ctx context.Context) ([]*components.Departament, error) {
	list_departaments, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los departamentos: %v", err)
		return nil, err
	}
	return list_departaments, nil
}
