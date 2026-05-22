package components

import (
	"context"
	"log"
	"sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type EnviromentServ interface {
	GetAll(ctx context.Context) ([]*components.Enviroment, error)
}

type EnviromentServImpl struct {
	Repo     components.EnviromentRepo
	validate *validator.Validate
}

// Instancia para servicios comunes
func NewEnviromentService(repo components.EnviromentRepo) EnviromentServ {
	return &EnviromentServImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *EnviromentServImpl) GetAll(ctx context.Context) ([]*components.Enviroment, error) {
	enviroment, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, err
	}
	return enviroment, nil
}
