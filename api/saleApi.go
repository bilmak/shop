package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shop/core"
	"shop/storage/sales"
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
		fmt.Println("CreateProduct validate error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SA.SS.CreateSale(sales)
}
