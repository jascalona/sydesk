package components

import "context"

type TypeRequest struct {
	ID        int    `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"created_at" db:"created_at"`
}

type ValidationTypeRequest struct {
	NAME string `json:"name" binding:"required, max=225"`
}

type InterfaceTypeRequest interface {
	GetAll(ctx context.Context) ([]*TypeRequest, error)
	//Created(ctx context.Context, tp_request *TypeRequest) error
}
