package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"
)

type PriorityServ interface {
	GetAll(ctx context.Context) ([]*components.Priority, error)
}

type PriorityServImpl struct {
	Repo components.InterfacePriority
}

func NewPriorityServ(repo components.InterfacePriority) PriorityServ {
	return &PriorityServImpl{
		Repo: repo,
	}
}

func (s *PriorityServImpl) GetAll(ctx context.Context) ([]*components.Priority, error) {
	list_priority, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		return nil, fmt.Errorf("Error al obtener los registros")
	}
	return list_priority, nil
}
