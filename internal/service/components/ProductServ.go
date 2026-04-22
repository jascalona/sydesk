package components

import (
	"context"
	"fmt"
	"log"
	"sydesk/pkg/domain/components"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]*components.Product, error)
	Created(ctx context.Context, product *components.Product) error
}

type ProductServiceImpl struct {
	Repo     components.ProductRepo
	validate *validator.Validate
}

// Instancia para los servicios comunes
func NewProductService(repo components.ProductRepo) ProductService {
	return &ProductServiceImpl{
		Repo:     repo,
		validate: validator.New(),
	}
}

func (s *ProductServiceImpl) Created(ctx context.Context, product *components.Product) error {
	// Invocamos la validacion del struct
	if err := s.validate.Struct(product); err != nil {
		log.Println("Error al procesar la solicitud %v:", err)
		return fmt.Errorf("El mensaje no cumple con los parametros definidos: %w", err)
	}

	// Asignacion del valor is_active, estado del producto post-registro (activo por defecto)
	product.IS_ACTIVE = 0 == 0

	err := s.Repo.Created(ctx, product)
	if err != nil {
		log.Println("Error al crear el registro:", err.Error())
		return fmt.Errorf("No se pudo crear el registro: ", err.Error())
	}
	return nil
}

func (s *ProductServiceImpl) GetAll(ctx context.Context) ([]*components.Product, error) {
	products, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los datos %V: ", err)
		return nil, fmt.Errorf("Error al obtener los datos %w", err)
	}
	return products, nil
}
