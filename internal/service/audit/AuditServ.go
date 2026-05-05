package audit

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"

	"github.com/google/uuid"
)

type AuditService interface {
	GetAuditCustomer(ctx context.Context, parentID uuid.UUID) ([]*audit.Audit, error)
}

type AuditServImp struct {
	Repo audit.AuditCustomer
}

func NewAuditServ(repo audit.AuditCustomer) AuditService {
	return &AuditServImp{
		Repo: repo,
	}
}

func (s *AuditServImp) GetAuditCustomer(ctx context.Context, parentID uuid.UUID) ([]*audit.Audit, error) {
	audit_serv, err := s.Repo.GetAuditCustomer(ctx, parentID)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return audit_serv, nil
}
