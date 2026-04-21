package handler

import "sydesk/internal/service/components"

type ProductHandler struct {
	Service components.ProductService
}

func NewProductHandler(service components.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}
