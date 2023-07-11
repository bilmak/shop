package sales

import (
	"database/sql"
	"fmt"
	"shop/core"
)

type SaleStorage struct {
	DB *sql.DB
}

func (ss SaleStorage) CreateSale(sale core.Sales) {
	query := "insert into sales(id, date, customerid, employeesid, productid, unitprice, quantity, totalprice) values ($1,$2, $3, $4, $5, $6, $7,$8)"
	_, err:= ss.DB.Exec(query, sale.ID, sale.Date, sale.CustomerID, sale.EmployeesID, sale.ProductID, sale.UnitPrice, sale.Quantity, sale.TotalPrice)
	if err != nil{
		fmt.Println("Create sales err", err)
	}
	
}
