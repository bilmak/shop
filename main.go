package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"shop/api"
	"shop/config"
	"shop/storage/customer"
	"shop/storage/employees"
	"shop/storage/products"
	"shop/storage/sales"

	_ "github.com/lib/pq"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	apiStruct := api.ApiStuct{
		CS: api.CustomerApi{CC: customer.CustomerStorage{DB: db}},
		ES: api.EmployeesApi{ES: employees.EmployeesStorage{DB: db}},
		PS: api.ProductApi{PS: products.ProductStorage{DB: db}},
		SA: api.SaleApi{SS: sales.SaleStorage{DB: db}},
	}

	err = http.ListenAndServe(":8010", apiStruct.Routs(db))
	if err != nil {
		fmt.Println(err)
	}

}
