package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"shop/core"
	"shop/storage/employees"
	"strconv"
)

type EmployeesApi struct {
	es employees.EmployeesStorage
}

func (em EmployeesApi) CreateEmployees(w http.ResponseWriter, r *http.Request) {
	var employees core.Employees

	err := json.NewDecoder(r.Body).Decode(&employees)
	if err != nil {
		fmt.Println("CreateEmployees, decoder error", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	em.es.CreateEmployees(employees)
}

func (em EmployeesApi) UpdateEmployees(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	var employees core.Employees
	employees.ID = idInt

	err = json.NewDecoder(r.Body).Decode(&employees)
	if err != nil {
		fmt.Println("UpdateEmployees, decoder erorr", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	em.es.UpdateEmployees(employees)

}

func (em EmployeesApi) DeleteById(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("DeleteById, cant convert id to int", err)
		return
	}
	em.es.DeleteById(idInt)
}

func (em EmployeesApi) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := em.es.GetAllEmployees()
	err := json.NewEncoder(w).Encode(employees)
	if err != nil {
		fmt.Println("GetAllEmployees: cant read json", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
