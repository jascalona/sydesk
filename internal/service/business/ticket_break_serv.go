package business

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/business"
)

type TicketBreakServ interface {
	GetAll(ctx context.Context) ([]*business.TicketBreak, error)
}

type TicketBreakSrvImpl struct {
	Repo business.InterfaceBreak
}

func NewTicketBreakServ(repo business.InterfaceBreak) TicketBreakServ {
	return &TicketBreakSrvImpl{
		Repo: repo,
	}
}

func (s *TicketBreakSrvImpl) GetAll(ctx context.Context) ([]*business.TicketBreak, error) {
	ticket, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		return nil, fmt.Errorf("Error al obtener los registros")
	}
	return ticket, nil
}
