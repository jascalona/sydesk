package components

import "context"

type Enviroment struct {
	ID        string `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"createdAt" db:"created_at"`
}

type ValidationEnviroment struct {
	NAME string `json:name validate:"required, max=200"`
}

type EnviromentRepo interface {
	GetAll(ctx context.Context) ([]*Enviroment, error)
}
