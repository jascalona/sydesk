package audit

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"
)

type ChannelServ interface {
	GetAll(ctx context.Context) ([]*audit.ChannelRep, error)
}

type ChannelServiceImpl struct {
	Repo audit.InterfaceChannel
}

func NewChannelServ(repo audit.InterfaceChannel) ChannelServ {
	return &ChannelServiceImpl{
		Repo: repo,
	}
}

func (s *ChannelServiceImpl) GetAll(ctx context.Context) ([]*audit.ChannelRep, error) {
	channel, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los registros %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros: %v", err.Error())
	}
	return channel, nil
}
