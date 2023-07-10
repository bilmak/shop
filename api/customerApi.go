package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"shop/core"
	"shop/storage/customer"
	"strconv"
)

type CustomerApi struct {
	cusromerStor customer.CustomerStorage
}

func (c CustomerApi) CreateCustomers(w http.ResponseWriter, r *http.Request) {
	var customers core.Customer
	err := json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	c.cusromerStor.CreateCustomers(customers)
	w.WriteHeader(http.StatusOK)
}

func (c CustomerApi) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	var customers core.Customer
	customers.ID = idInt

	err = json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	c.cusromerStor.UpdateCustomer(customers)
}

func (c CustomerApi) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	c.cusromerStor.DeleteCustomer(idInt)
}

func (c CustomerApi) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customer := c.cusromerStor.GetAllCustomers()
	err := json.NewEncoder(w).Encode(customer)
	if err != nil {
		fmt.Println("getAll: cant read json", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
