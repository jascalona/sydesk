package organization

import (
	"context"
	"database/sql"
	domain "sydesk/pkg/domain/organization"
)

type userRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepo {
	return &userRepo{DB: db}
}

// --- REGISTRO DE USUARIO ---//
func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users(id, alias, name, surname, email, phone, departament_id, password_hash, role_id, is_active)	
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9, $10)`

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
		user.PASSWORD_HASH,
		user.ROLE_ID,
		user.IS_ACTIVE,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetAll(ctx context.Context) ([]*domain.User, error) {
	query := `
		SELECT
	      id, alias, name,surname, email, phone, departament_id, role_id, is_active
		FROM users`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0)

	for rows.Next() {
		user := &domain.User{}
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
