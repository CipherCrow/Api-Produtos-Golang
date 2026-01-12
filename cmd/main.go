package main

import (
	CasosDeUso "api-produtos-golang/casosDeUso"
	"api-produtos-golang/controller"
	"api-produtos-golang/db"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}
	//repository

	// Casos de Uso
	produtoCasoDeUso := CasosDeUso.NewProductUseCase()

	//Controllers
	productController := controller.NewProductController(produtoCasoDeUso)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/produtos/", productController.GetProducts)

	server.Run(":8000")
}
