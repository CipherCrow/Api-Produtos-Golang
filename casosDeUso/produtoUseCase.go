package CasosDeUso

import "api-produtos-golang/model"

type ProdutoCasoDeUso struct {
}

func NewProductUseCase() ProdutoCasoDeUso {
	return ProdutoCasoDeUso{}
}

func (pu *ProdutoCasoDeUso) GetAllProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
