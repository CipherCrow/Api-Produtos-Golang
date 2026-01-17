package repository

import (
	"api-produtos-golang/model"
	"database/sql"
	"fmt"
)

type ProdutoRepository struct {
	connection *sql.DB
}

func NewProdutoRepository(psql *sql.DB) ProdutoRepository {
	return ProdutoRepository{connection: psql}
}

func (pr *ProdutoRepository) GetAllProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var listaProdutos []model.Product
	var objetoProduto model.Product

	for rows.Next() {
		err = rows.Scan(
			&objetoProduto.ID,
			&objetoProduto.Name,
			&objetoProduto.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		listaProdutos = append(listaProdutos, objetoProduto)
	}

	rows.Close()

	return listaProdutos, nil
}
func (pr *ProdutoRepository) Save(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name,price)" +
		"values ($1,$2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProdutoRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &produto, nil
}
