package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"shop/core"
	"shop/storage/products"
	"strconv"
)

type ProductApi struct {
	PS products.ProductStorage
}

func (pa ProductApi) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product core.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println("CreateProduct decoder error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = product.ValidateProduct()
	if err != nil {
		fmt.Println("CreateProduct validate error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pa.PS.CreateProduct(product)
}

func (pa ProductApi) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)

	var product core.Product
	product.ID = idInt

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println("UpdateProduct decoder error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pa.PS.UpdateProduct(product)
}

func (pa ProductApi) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := path.Base(r.URL.String())
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	pa.PS.DeleteProduct(idInt)

}

func (pa ProductApi) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products := pa.PS.GetAllProduct()
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		fmt.Println("GetAllProduct, newEncoder err", err)
	}
}
