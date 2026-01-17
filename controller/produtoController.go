package controller

import (
	usecase "api-produtos-golang/casosDeUso"
	"api-produtos-golang/model"
	"net/http"
	"strconv"

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

func (p *ProdutoController) FindProductById(ctx *gin.Context) {

	id := ctx.Param("idProduto")

	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	idProduto, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	produto, err := p.produtoUseCase.GetProductById(idProduto)

	if produto != nil {
		response := model.Response{
			Message: "Produto nao foi encontrado!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, produto)
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
