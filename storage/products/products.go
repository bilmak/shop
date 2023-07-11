package products

import (
	"database/sql"
	"fmt"
	"shop/core"
)

type ProductStorage struct {
	DB *sql.DB
}

func (ps ProductStorage) CreateProduct(product core.Product) error {
	query := "insert into product(name, price)values($1, $2)"
	_, err := ps.DB.Exec(query, product.Name, product.Price)
	if err != nil {
		fmt.Println(" CreateProduct, query err", err)
		return err
	}
	return nil
}

func (ps ProductStorage) UpdateProduct(product core.Product) error {
	query := "update product set name = $1, price = $2 where id = $3"
	_, err := ps.DB.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println("  UpdateProduct, query err", err)
		return err
	}
	return nil
}

func (ps ProductStorage) DeleteProduct(id int) error {
	query := "delete from product where id = $1"
	result, _ := ps.DB.Exec(query, id)
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("DeleteProduct error", err)
		return err
	}
	if row == 0 {
		return err
	}
	return nil
}

func (ps ProductStorage) GetAllProduct() ([]core.Product, error) {
	query := "select id, name, price from product"
	row, err := ps.DB.Query(query)
	if err != nil {
		fmt.Println(" GetAllProduct, query err", err)
		return []core.Product{}, err
	}

	var products []core.Product
	for row.Next() {
		var product core.Product

		row.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}
	return products, nil
}
