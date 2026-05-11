package audit

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/audit"

	"github.com/google/uuid"
)

type AsService interface {
	GetAll(ctx context.Context, cpr_id uuid.UUID) ([]*audit.AuditServ, error)
	Created(ctx context.Context, auditStatus *audit.AuditServ) error
	UpdateAS(ctx context.Context, id int64, input audit.ValidateAS) error
}

type AsServiceImpl struct {
	Repo audit.AuditServInterface
}

func NewAsServ(repo audit.AuditServInterface) AsService {
	return &AsServiceImpl{
		Repo: repo,
	}
}

func (s *AsServiceImpl) GetAll(ctx context.Context, cpr_id uuid.UUID) ([]*audit.AuditServ, error) {

	audit_status, err := s.Repo.GetAll(ctx, cpr_id)
	if err != nil {
		log.Printf("Error al obtener los registros: %v", err.Error())
		return nil, fmt.Errorf("error al obtener los registros")
	}
	return audit_status, nil
}

func (s *AsServiceImpl) Created(ctx context.Context, auditStatus *audit.AuditServ) error {
	// despues se pueda integrar reglas de negocio

	// persistencia de data
	err := s.Repo.Created(ctx, auditStatus)
	if err != nil {
		log.Printf("Error al procesar la solicitud %v", err.Error())
		return fmt.Errorf("no se pudo crear el registro, por favor verifique la traza de la operacion")
	}
	return nil
}

func (s *AsServiceImpl) UpdateAS(ctx context.Context, id int64, input audit.ValidateAS) error {

	// VERIFICAR SI EL REGISTRO EXISTE
	existingAudit, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("no se encontró el registro de auditoría: %w", err)
	}

	// REGLAS DE NEGOCIO (VERIFICACION DE QUE END_AT SEA MAYOR A START_AT)

	// CONVERSION DE LOS DATOS PARSEADOS EN EL JSON PARA LA BD
	existingAudit.CPR_ID = input.CPR_ID
	existingAudit.START_AT = &input.START_AT
	existingAudit.ENVIRONMENT = input.ENVIRONMENT
	existingAudit.SERVICES = &input.SERVICES
	existingAudit.DESCRIPTION = input.DESCRIPTION
	existingAudit.ACTIVITIES = input.ACTIVITIES
	existingAudit.END_AT = &input.END_AT
	existingAudit.LAST_OPERATION = &input.LAST_OPERATION

	// PERSISTIR
	// Aquí llamamos al Repo enviando la entidad completa
	return s.Repo.UpdateAS(ctx, existingAudit)
}
