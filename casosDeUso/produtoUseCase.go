package casosDeUso

import (
	"api-produtos-golang/model"
	"api-produtos-golang/repository"
)

type ProdutoCasoDeUso struct {
	repository repository.ProdutoRepository
}

func NewProductUseCase(us repository.ProdutoRepository) ProdutoCasoDeUso {
	return ProdutoCasoDeUso{repository: us}
}

func (pu *ProdutoCasoDeUso) GetAllProducts() ([]model.Product, error) {
	return pu.repository.GetAllProducts()
}

func (pu *ProdutoCasoDeUso) Save(product model.Product) (model.Product, error) {
	productId, err := pu.repository.Save(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}
