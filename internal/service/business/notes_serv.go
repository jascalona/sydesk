package business

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/business"
)

type NotesService interface {
	GetByTaskId(ctx context.Context, task_id string) ([]*business.Notes, error)
	CreatedNotes(ctx context.Context, notes *business.Notes) error
	UpdateNotes(ctx context.Context, note_id *business.Notes) error
}

type NotesServiceImpl struct {
	Repo business.InterfaceNotes
}

func NewNotesService(repo business.InterfaceNotes) NotesService {
	return &NotesServiceImpl{
		Repo: repo,
	}
}

func (s *NotesServiceImpl) GetByTaskId(ctx context.Context, task_id string) ([]*business.Notes, error) {
	if task_id == "" {
		return nil, fmt.Errorf("Error de interpretado: El ID de la tarea es requerido")
	}
	rows_notes, err := s.Repo.GetByTaskId(ctx, task_id)
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		return nil, fmt.Errorf("Error al obtener los registros: %v", err.Error())
	}
	return rows_notes, nil
}

func (s *NotesServiceImpl) UpdateNotes(ctx context.Context, note_id *business.Notes) error {
	err := s.Repo.UpdateNotes(ctx, note_id)
	if err != nil {
		log.Println("Error al actualizar el registro: ", err.Error())
		return fmt.Errorf("Error al actualizar el registro, por favor verifique la traza de la operacion")
	}
	return nil
}

func (s *NotesServiceImpl) CreatedNotes(ctx context.Context, notes *business.Notes) error {
	err := s.Repo.CreatedNotes(ctx, notes)
	if err != nil {
		log.Printf("Error al procesar la solicitud %s", err.Error())
		return fmt.Errorf("No se puede crear el registro, pro favor verifique la traza de la operacion: %v", err.Error())
	}
	return nil
}
