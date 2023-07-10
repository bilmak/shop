package products

import (
	"database/sql"
	"fmt"
	"shop/core"
)

type ProductStorage struct {
	DB *sql.DB
}

func (ps ProductStorage) CreateProduct(product core.Product) {
	query := "insert into product(name, price)values($1, $2)"
	ps.DB.Exec(query, product.Name, product.Price)
}

func (ps ProductStorage) UpdateProduct(product core.Product) {
	query := "update product set name = $1, price = $2 where id = $3"
	ps.DB.Exec(query, product.Name, product.Price, product.ID)
}

func (ps ProductStorage) DeleteProduct(id int) {
	query := "delete from product where id = $1"
	result, _ := ps.DB.Exec(query, id)
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("DeleteProduct error", err)
		return
	}
	if row == 0 {
		return
	}
}

func (ps ProductStorage) GetAllProduct() []core.Product {
	query := "select id, name, price from product"
	row, err := ps.DB.Query(query)
	if err != nil {
		fmt.Println(" GetAllProduct, query err", err)
	}

	var products []core.Product
	for row.Next() {
		var product core.Product

		row.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}
	return products
}
