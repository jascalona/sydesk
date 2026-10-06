package business

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sydesk/pkg/domain/business"
	"time"

	"github.com/go-playground/validator/v10"
)

type TicketRequestServ interface {
	GetAll(ctx context.Context) ([]*business.TicketRequest, error)
	Created(ctx context.Context, new_ticket *business.TicketRequest) error
	UpdateByTicketId(ctx context.Context, ticket *business.TicketRequest) error

	// interfaz dedicada a la vista
	GetVWTicketRequest(ctx context.Context) ([]*business.VWTicketRequest, error)
}

type TicketRequestServImpl struct {
	Repo     business.InterfaceTicketRequest
	validate *validator.Validate
}

func NewTicketRequestServ(repo business.InterfaceTicketRequest) TicketRequestServ {
	return &TicketRequestServImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

// METODO PARA LA VISTA
func (s *TicketRequestServImpl) GetVWTicketRequest(ctx context.Context) ([]*business.VWTicketRequest, error) {
	vw_ticket, err := s.Repo.GetVWTicketRequest(ctx)
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		return nil, fmt.Errorf("Error al obtener los registros")
	}
	return vw_ticket, nil
}

func (s *TicketRequestServImpl) GetAll(ctx context.Context) ([]*business.TicketRequest, error) {
	ticket, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		return nil, fmt.Errorf("Error al obtener los registros")
	}
	return ticket, nil
}

func (s *TicketRequestServImpl) UpdateByTicketId(ctx context.Context, ticket *business.TicketRequest) error {
	err := s.Repo.UpdateByTicketId(ctx, ticket)
	if err != nil {
		log.Println("Error al actualizar el registro: ", err.Error())
		return fmt.Errorf("Error al actualizar el registro, por favor verifique la traza de la operacion")
	}
	return nil
}

func (s *TicketRequestServImpl) Created(ctx context.Context, new_ticket *business.TicketRequest) error {
	// las validaciones de negocio deben de ser agregadas mas adelante

	// generador de codigo de ticket
	// Inicializar la semilla (esencial para que los resultados varíen en cada ejecución)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	long := 13
	const charts = "ABCD1234567890"

	var id_ticket strings.Builder
	for i := 0; i < long; i++ {
		randomIndex := rng.Intn(len(charts))
		// aniadimos el caracter a la cadena
		id_ticket.WriteByte(charts[randomIndex])
	}

	// inyectamos el valor del id generado para el ticket
	new_ticket.ID = id_ticket.String()

	err := s.Repo.Created(ctx, new_ticket)
	if err != nil {
		log.Println("Error al procesar la solicitud ", err.Error())
		return fmt.Errorf("No se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil

}
