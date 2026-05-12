package audit

import "context"

type ChannelRep struct {
	ID        int    `json:"id" db:"id"`
	NAME      string `json:"name" db:"name"`
	CREATEDAT string `json:"created_at" db:"created_at"`
}

type ValidationChannel struct {
	NAME string `json:"name" binding:"required, max=100"`
}

type InterfaceChannel interface {
	GetAll(ctx context.Context) ([]*ChannelRep, error)
}
