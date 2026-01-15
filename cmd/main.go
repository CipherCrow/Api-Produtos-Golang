package main

import (
	"api-produtos-golang/casosDeUso"
	"api-produtos-golang/controller"
	"api-produtos-golang/db"
	"api-produtos-golang/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}
	//repository
	produtoRepository := repository.NewProdutoRepository(dbConnection)

	// Casos de Uso
	produtoCasoDeUso := casosDeUso.NewProductUseCase(produtoRepository)

	//Controllers
	productController := controller.NewProductController(produtoCasoDeUso)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/produtos/", productController.GetProducts)
	server.POST("/produtos/", productController.CadastrarProduto)

	server.Run(":8000")
}
