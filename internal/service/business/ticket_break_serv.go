package business

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/business"
)

type TicketBreakServ interface {
	GetAll(ctx context.Context) ([]*business.TicketBreak, error)
	Created(ctx context.Context, t_break *business.TicketBreak) error
	GetTicketId(ctx context.Context, ticket_id string) ([]*business.TicketBreak, error)
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

func (s *TicketBreakSrvImpl) Created(ctx context.Context, t_break *business.TicketBreak) error {
	// las validaciones de negocio deben ser agregada mas adelante

	err := s.Repo.Created(ctx, t_break)
	if err != nil {
		log.Println("Error al procesar la solicitud ", err.Error())
		return fmt.Errorf("No se pudo crear el registro, por favor verifique la traza de la operacion")
	}

	return nil
}

func (s *TicketBreakSrvImpl) GetTicketId(ctx context.Context, ticket_id string) ([]*business.TicketBreak, error) {
	if ticket_id == "" {
		return nil, fmt.Errorf("Error service: el ID del ticket es requerido")
	}

	ticket_break, err := s.Repo.GetTicketId(ctx, ticket_id)

	if err != nil {
		log.Println(ticket_id)
		log.Println("Error service: no se pudo obtener los estados del ticket: ", ticket_break)
		return nil, fmt.Errorf("no se pudo obtener los estados del ticket: %w", err)
	}

	return ticket_break, nil

}
