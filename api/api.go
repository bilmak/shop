package api

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ApiStuct struct {
	cs CustomerApi
	es EmployeesApi
	ps ProductApi
}

func (a *ApiStuct) Routs(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/customers", a.cs.CreateCustomers)
	r.Put("/customers/{id}", a.cs.UpdateCustomer)
	r.Delete("/customers/{id}", a.cs.DeleteCustomer)
	r.Get("/customers", a.cs.GetAllCustomers)

	r.Post("/employees", a.es.CreateEmployees)
	r.Put("/employees/{id}", a.es.UpdateEmployees)
	r.Delete("/employees/{id}", a.es.DeleteById)
	r.Get("/employees", a.es.GetAllEmployees)

	r.Post("/products", a.ps.CreateProduct)
	r.Put("/products/{id}", a.ps.UpdateProduct)
	r.Delete("/products/{id}", a.ps.DeleteProduct)
	r.Get("/products", a.ps.GetAllProduct)
	return r
}

func CreateCustomers(db *sql.DB) {
	panic("unimplemented")
}
