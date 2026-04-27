package middlerware

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"sydesk/pkg/domain"
	"time"
)

type authRepo struct {
	DB *sql.DB
}

func NewAuthRepo(db *sql.DB) domain.AuthRepo {
	return &authRepo{DB: db}
}

func (r *authRepo) GetByEmail(email string) (*domain.Auth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT id, email, password_hash FROM users WHERE email = $1`

	var user domain.Auth
	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.EMAIL,
		&user.PASSWORD_HASH,
	)

	if err != nil {
		log.Printf("Error QueryContext %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Usuario no encontrado")
		}
	}
	// retorno del objeto para UX
	return &user, err

}
