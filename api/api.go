package api

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ApiStuct struct {
	CS CustomerApi
	ES EmployeesApi
	PS ProductApi
	SA SaleApi
}

func (a *ApiStuct) Routs(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/customers", a.CS.CreateCustomers)
	r.Put("/customers/{id}", a.CS.UpdateCustomer)
	r.Delete("/customers/{id}", a.CS.DeleteCustomer)
	r.Get("/customers", a.CS.GetAllCustomers)

	r.Post("/employees", a.ES.CreateEmployees)
	r.Put("/employees/{id}", a.ES.UpdateEmployees)
	r.Delete("/employees/{id}", a.ES.DeleteById)
	r.Get("/employees", a.ES.GetAllEmployees)

	r.Post("/products", a.PS.CreateProduct)
	r.Put("/products/{id}", a.PS.UpdateProduct)
	r.Delete("/products/{id}", a.PS.DeleteProduct)
	r.Get("/products", a.PS.GetAllProduct)

	r.Post("/sales", a.SA.CreateSale)
	return r
}

func CreateCustomers(db *sql.DB) {
	panic("unimplemented")
}
