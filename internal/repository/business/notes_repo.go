package business

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/business"
)

type NotesRepo struct {
	DB *sql.DB
}

func NewNotesRepo(db *sql.DB) business.InterfaceNotes {
	return &NotesRepo{DB: db}
}

func (r *NotesRepo) GetByTaskId(ctx context.Context, task_id string) ([]*business.Notes, error) {
	query := `
		SELECT 
			id,
			task_id,
			topic,
			description,
			created_by,
			created_at
		FROM notes
		WHERE task_id = $1`

	rows, err := r.DB.QueryContext(ctx, query, task_id)
	if err != nil {
		log.Print("Error al correr el query: ", err.Error())
		return nil, err
	}

	var notes_by_task []*business.Notes
	for rows.Next() {
		list := &business.Notes{}
		if err := rows.Scan(
			&list.ID,
			&list.TASK_ID,
			&list.TOPIC,
			&list.DESCRIPTION,
			&list.CREATED_BY,
			&list.CREATED_AT,
		); err != nil {
			log.Print("Error al escanear los registros: ", err.Error())
			return nil, err
		}
		notes_by_task = append(notes_by_task, list)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notes_by_task, nil
}

func (r *NotesRepo) CreatedNotes(ctx context.Context, notes *business.Notes) error {

	query := `INSERT INTO notes (
				task_id,
				topic,
				description,
				created_by)VALUES($1,$2,$3,$4)`

	_, err := r.DB.QueryContext(ctx, query,
		notes.TASK_ID,
		notes.TOPIC,
		notes.DESCRIPTION,
		notes.CREATED_BY)

	if err != nil {
		return err
	}
	return nil

}
