package business

import (
	"context"
	"database/sql"
	"log"
	"sydesk/pkg/domain/business"
)

type TaskRequestRepo struct {
	DB *sql.DB
}

func NewTaskRequestRepo(db *sql.DB) business.InterfaceTaskRequest {
	return &TaskRequestRepo{DB: db}
}

func (r *TaskRequestRepo) CreatedRequestTask(ctx context.Context, task *business.TaskRequest) error {

	query := `
		INSERT INTO task_request(
			id,
			ticket_id,
			priority_id,
			status_id,
			created_by,
			assigned_departament,
			assigned_user,
			topic,
			description,
			expired_in,
			closed_in)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := r.DB.ExecContext(ctx, query,
		task.ID,
		task.TICKET_ID,
		task.PRIORITY,
		task.STATUS_ID,
		task.CREATED_BY,
		task.ASSIGNED_DPT,
		task.ASSIGNED_USER,
		task.TOPIC,
		task.DESCRIPTION,
		task.EXPIRED_IN,
		task.CLOSED_IN,
	)

	if err != nil {
		log.Printf("Error al correr el query: %v", err.Error())
		return err
	}

	return nil

}

// METODO PARA CONSUMO DE VW
func (r *TaskRequestRepo) GetVWTaskRequest(ctx context.Context) ([]*business.VWTaskRequest, error) {
	query := `
		SELECT 
			id,
			ticket_id,
			created_by,
			topic,
			description,
			created_at,
			expired_in,
			closed_in,
			priority_name,
			status_name,
			assigned_departament_name,
			assigned_user_name
		FROM vw_task_request_details ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("Error al correr el query")
		return nil, err
	}

	defer rows.Close()

	vw_task := make([]*business.VWTaskRequest, 0)

	for rows.Next() {
		lt := &business.VWTaskRequest{}
		err := rows.Scan(
			&lt.ID,
			&lt.TICKET_ID,
			&lt.CREATED_BY,
			&lt.TOPIC,
			&lt.DESCRIPTION,
			&lt.CREATED_AT,
			&lt.EXPIRED_IN,
			&lt.CLOSED_IN,
			&lt.PRIORITY,
			&lt.STATUS,
			&lt.ASSIGNED_DPT,
			&lt.ASSIGNED_USER,
		)

		if err != nil {
			log.Println("Error al aplicar el escanner", err.Error())
			return nil, err
		}

		vw_task = append(vw_task, lt)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vw_task, nil

}

// RETORNO DE VSITAS POR TICKET_ID
func (r *TaskRequestRepo) GetVWTaskRequestByTicket(ctx context.Context, ticket_id string) ([]*business.VWTaskRequest, error) {

	query := `
		SELECT 
			id,
			ticket_id,
			created_by,
			topic,
			description,
			created_at,
			expired_in,
			closed_in,
			priority_name,
			status_name,
			assigned_departament_name,
			assigned_user_name
		FROM vw_task_request_details where ticket_id = $1 ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, ticket_id)

	if err != nil {
		log.Println("Error al correr el query", err.Error())
		return nil, err
	}

	var rows_task []*business.VWTaskRequest
	for rows.Next() {
		list := &business.VWTaskRequest{}
		if err := rows.Scan(
			&list.ID,
			&list.TICKET_ID,
			&list.CREATED_BY,
			&list.TOPIC,
			&list.DESCRIPTION,
			&list.CREATED_AT,
			&list.EXPIRED_IN,
			&list.CLOSED_IN,
			&list.PRIORITY,
			&list.STATUS,
			&list.ASSIGNED_DPT,
			&list.ASSIGNED_USER,
		); err != nil {
			log.Println("Error al aplicar el escanner")
			return nil, err
		}

		rows_task = append(rows_task, list)
	}

	return rows_task, nil

}
