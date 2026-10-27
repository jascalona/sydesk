package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"
)

type SLAService interface {
	GetAll(ctx context.Context) ([]*components.SLA, error)
}

type SLAServiceImpl struct {
	Repo components.InterfaceSLA
}

func NewSLAService(repo components.InterfaceSLA) SLAService {
	return &SLAServiceImpl{
		Repo: repo,
	}
}

func (s *SLAServiceImpl) GetAll(ctx context.Context) ([]*components.SLA, error) {
	slas, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros: ", err)
		return nil, fmt.Errorf("Error al obtener los registros: %w", err)
	}
	return slas, nil
}
