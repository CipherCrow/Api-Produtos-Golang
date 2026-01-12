package controller

import (
	usecase "api-produtos-golang/casosDeUso"

	"github.com/gin-gonic/gin"
)

type ProdutoController struct {
	produtoUseCase usecase.ProdutoCasoDeUso
}

func NewProductController(puc usecase.ProdutoCasoDeUso) ProdutoController {
	return ProdutoController{
		puc,
	}
}

func (p *ProdutoController) GetProducts(ctx *gin.Context) {

}

func (p *ProdutoController) GetAllProducts(ctx *gin.Context) {
	// Implementation will go here
}
