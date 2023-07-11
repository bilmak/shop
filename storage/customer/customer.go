package customer

import (
	"database/sql"
	"fmt"
	"shop/core"
)

type CustomerStorage struct {
	DB *sql.DB
}

func (c CustomerStorage) CreateCustomers(customers core.Customer) {
	query := "insert into customers(name) values ($1);"
	_, err := c.DB.Exec(query, customers.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (c CustomerStorage) UpdateCustomer(customers core.Customer) {

	query := "UPDATE customers SET name = $1 WHERE id = $2;"
	_, err := c.DB.Exec(query, customers.Name, customers.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c CustomerStorage) DeleteCustomer(id int) {
	query := "delete from customers where id = $1"
	result, err := c.DB.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if row == 0 {
		return
	}
}

func (c CustomerStorage) GetAllCustomers() ([]core.Customer, error) {
	query := "select id, name from customers c "

	row, err := c.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return []core.Customer{}, err
	}

	var customers []core.Customer

	for row.Next() {
		var customer core.Customer
		err := row.Scan(&customer.ID, &customer.Name)
		if err != nil {
			fmt.Println(err)
			return []core.Customer{}, err
		}
		customers = append(customers, customer)

	}
	return customers, nil

}
