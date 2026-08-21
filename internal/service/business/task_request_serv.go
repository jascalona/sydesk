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

type TaskRequestServ interface {
	CreatedRequestTask(ctx context.Context, task *business.TaskRequest) error
	GetVWTaskRequest(ctx context.Context) ([]*business.VWTaskRequest, error)
	GetVWTaskRequestByTicket(ctx context.Context, ticket_id string) ([]*business.VWTaskRequest, error)
}

type TaskRequestImpl struct {
	Repo     business.InterfaceTaskRequest
	validate *validator.Validate
}

func NewTaskRequestServ(repo business.InterfaceTaskRequest) TaskRequestServ {
	return &TaskRequestImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

// VW PARA LOS TASK
func (s *TaskRequestImpl) GetVWTaskRequest(ctx context.Context) ([]*business.VWTaskRequest, error) {
	vw_task, err := s.Repo.GetVWTaskRequest(ctx)
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		return nil, fmt.Errorf("Error al obtener los registros: %v", err.Error())
	}
	return vw_task, nil

}

func (s *TaskRequestImpl) GetVWTaskRequestByTicket(ctx context.Context, ticket_id string) ([]*business.VWTaskRequest, error) {
	if ticket_id == "" {
		return nil, fmt.Errorf("Error de interpretado: El ID de la solicitud es requerido")
	}
	rows_task, err := s.Repo.GetVWTaskRequestByTicket(ctx, ticket_id)
	if err != nil {
		log.Println("Error al obtener los registros", err.Error())
		return nil, fmt.Errorf("Error services: No se pudieron obtener los registos asociados %s", ticket_id)
	}

	return rows_task, nil
}

func (s *TaskRequestImpl) CreatedRequestTask(ctx context.Context, task *business.TaskRequest) error {

	// Generador el ID de tarea

	rtask := rand.New(rand.NewSource(time.Now().UnixNano()))

	long := 6
	const charts = "TASK1234567890"

	var id_task strings.Builder
	for i := 0; i < long; i++ {
		iteration := rtask.Intn(len(charts))
		id_task.WriteByte(charts[iteration])
	}

	// inyeccion del id_task generado
	task.ID = id_task.String()

	err := s.Repo.CreatedRequestTask(ctx, task)
	if err != nil {
		log.Println("Error al procesar la solicitud ", err.Error())
		return fmt.Errorf("No se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil

}
