package organization

import (
	"context"
	"database/sql"
	"sydesk/pkg/domain"
)

type userRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) Create(ctx context.Context, user *domain.Users) error {
	query := `
		INSERT INTO users(id, alias, name,surname, email,  phone, departament_id, role_id, is_active)	
		VALUES(?,?,?,?,?,?,?,?,?)`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		user.ID,
		user.ALIAS,
		user.NAME,
		user.SURNAME,
		user.EMAIL,
		user.PHONE,
		user.DEPARTAMENT_ID,
		user.ROLE_ID,
		user.IS_ACTIVE,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetAll(ctx context.Context) ([]*domain.Users, error) {
	query := `
		SELECT
	      id, alias, name,surname, email, phone, departament_id, role_id, is_active
		FROM users`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.Users, 0)

	for rows.Next() {
		user := &domain.Users{}
		err := rows.Scan(
			&user.ID,
			&user.ALIAS,
			&user.NAME,
			&user.SURNAME,
			&user.EMAIL,
			&user.PHONE,
			&user.DEPARTAMENT_ID,
			&user.ROLE_ID,
			&user.IS_ACTIVE)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// validation error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	//return list data
	return users, nil
}
