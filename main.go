package main

import (
	"database/sql"
	"fmt"
	"shop/config"

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

	// apiStruct := api.ApiStuct{
	// 	api.CustomerApi{customer.CustomerStorage{DB: db}},
	// 	api.EmployeesApi{employees.EmployeesStorage{DB: db}},
	// 	api.ProductApi{products.ProductStorage{DB: db}},
	// }

	// err = http.ListenAndServe(":8010", Routs(db))
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
