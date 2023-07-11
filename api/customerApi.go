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
	CC customer.CustomerStorage
}

func (c CustomerApi) CreateCustomers(w http.ResponseWriter, r *http.Request) {
	var customers core.Customer
	err := json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	err = customers.ValidateCustomer()
	if err != nil {
		fmt.Println("CreateProduct validate error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.CC.CreateCustomers(customers)
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

	c.CC.UpdateCustomer(customers)
}

func (c CustomerApi) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	c.CC.DeleteCustomer(idInt)
}

func (c CustomerApi) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customer, err := c.CC.GetAllCustomers()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		fmt.Println("getAll: cant read json", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
