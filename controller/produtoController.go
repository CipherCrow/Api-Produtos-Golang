package controller

import (
	usecase "api-produtos-golang/casosDeUso"
	"api-produtos-golang/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProdutoController struct {
	produtoUseCase usecase.ProdutoCasoDeUso
}

func NewProductController(puc usecase.ProdutoCasoDeUso) ProdutoController {
	return ProdutoController{
		produtoUseCase: puc,
	}
}

func (p *ProdutoController) GetProducts(ctx *gin.Context) {

}

func (p *ProdutoController) GetAllProducts(ctx *gin.Context) {
	produtos, err := p.produtoUseCase.GetAllProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, produtos)
}

func (p *ProdutoController) CadastrarProduto(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.produtoUseCase.Save(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}
