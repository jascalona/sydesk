package audit

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"
)

type ItemServ interface {
	GetAll(ctx context.Context) ([]*audit.ItemActivities, error)
	Created(ctx context.Context, item *audit.ItemActivities) error
}

type ItemImpl struct {
	Repo audit.InterfaceItem
}

func NewItemService(repo audit.InterfaceItem) ItemServ {
	return &ItemImpl{
		Repo: repo,
	}
}

func (s *ItemImpl) GetAll(ctx context.Context) ([]*audit.ItemActivities, error) {

	items, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("error al obtener los registros %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return items, nil
}

func (s *ItemImpl) Created(ctx context.Context, item *audit.ItemActivities) error {
	// validaciones de negocios si son necesarias

	err := s.Repo.Created(ctx, item)
	if err != nil {
		log.Printf("Error al crear la registro: %v", err.Error())
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil
}
