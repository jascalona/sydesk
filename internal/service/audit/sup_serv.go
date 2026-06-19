package audit

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"
)

type SupService interface {
	GetAll(ctx context.Context) ([]*audit.SupServ, error)
}

type SupServiceImpl struct {
	Repo audit.InterfaceSup
}

func NewSupService(repo audit.InterfaceSup) SupService {
	return &SupServiceImpl{
		Repo: repo,
	}
}

func (s *SupServiceImpl) GetAll(ctx context.Context) ([]*audit.SupServ, error) {

	subproduct, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("error al obtener los registros %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}

	return subproduct, nil

}
