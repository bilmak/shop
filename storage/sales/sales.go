package sales

import (
	"database/sql"
	"fmt"
	"shop/core"
)

type SaleStorage struct {
	DB *sql.DB
}

func (ss SaleStorage) CreateSale(sale core.Sales) error {
	query := "insert into sales(date, customerid, employeesid, productid, unitprice, quantity, totalprice) values ($1,$2, $3, $4, $5, $6, $7)"
	_, err := ss.DB.Exec(query, sale.Date, sale.CustomerID, sale.EmployeesID, sale.ProductID, sale.UnitPrice, sale.Quantity, sale.TotalPrice)
	if err != nil {
		fmt.Println("Create sales err", err)
		return err
	}
	return nil

}

func (ss SaleStorage) UpdateSale(id int, sales core.Sales) error {
	querty := "update sales set date= $1, customerid =$2, employeesid=$3, productid=$4, unitprice=$5, quantity=$6, totalprice=$7 where id= $8"

	_, err := ss.DB.Exec(querty, sales.Date, sales.CustomerID, sales.EmployeesID, sales.ProductID, sales.UnitPrice, sales.Quantity, sales.TotalPrice, id)
	if err != nil {
		fmt.Println("Create sales err", err)
		return err
	}
	return nil
}

func (ss SaleStorage) DeleteSale(id int) error {
	query := "delete from sales where id =$1"

	_, err := ss.DB.Exec(query, id)
	if err != nil {
		fmt.Println("Create sales err", err)
		return err
	}
	return nil

}

func (ss SaleStorage) GetAllSales() ([]core.Sales, error) {
	var sales []core.Sales
	query := "select id, date, customerid, employeesid, productid, unitprice, quantity, totalprice from sales"

	row, err := ss.DB.Query(query)
	if err != nil {
		fmt.Println("delete sales err storage", err)
		return []core.Sales{}, err
	}
	for row.Next() {
		var sale core.Sales
		row.Scan(&sale.ID, &sale.Date, &sale.CustomerID, &sale.EmployeesID, &sale.ProductID, &sale.UnitPrice, &sale.Quantity, &sale.TotalPrice)
		sales = append(sales, sale)
	}
	return sales, nil

}
