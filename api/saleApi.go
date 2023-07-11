package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"shop/core"
	"shop/storage/sales"
	"strconv"
)

type SaleApi struct {
	SS sales.SaleStorage
}

func (SA SaleApi) CreateSale(w http.ResponseWriter, r *http.Request) {
	var sales core.Sales
	err := json.NewDecoder(r.Body).Decode(&sales)
	if err != nil {
		fmt.Println("CreateSales json err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = sales.ValidateSale()
	if err != nil {
		fmt.Println("CreateSale validate error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SA.SS.CreateSale(sales)
}

func (SA SaleApi) UpdateSale(w http.ResponseWriter, r *http.Request) {
	var sales core.Sales
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("CreateSale id int error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sales.ID = idInt
	err = json.NewDecoder(r.Body).Decode(&sales)
	if err != nil {
		fmt.Println("CreateSale json decoder error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SA.SS.UpdateSale(idInt, sales)

}

func (SA SaleApi) DeleteSale(w http.ResponseWriter, r *http.Request) {
	idSting := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idSting)
	if err != nil {
		fmt.Println("DeleteSale id int error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SA.SS.DeleteSale(idInt)

}

func (SA SaleApi) GetAllSales(w http.ResponseWriter, r *http.Request) {
	sales, _ := SA.SS.GetAllSales()
	err := json.NewEncoder(w).Encode(&sales)
	if err != nil {
		fmt.Println("Gett all sales json error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
