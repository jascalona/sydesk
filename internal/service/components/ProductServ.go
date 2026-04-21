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
}

type ProductServiceImpl struct {
	Repo     components.ProductRepo
	validate *validator.Validate
}

func (s *ProductServiceImpl) GetAll(ctx context.Context) ([]*components.Product, error) {
	products, err := s.Repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error al obtener los datos %V: ", err)
		return nil, fmt.Errorf("Error al obtener los datos %w", err)
	}
	return products, nil
}
