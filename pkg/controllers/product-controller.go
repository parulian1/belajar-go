package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/parulian1/belajar-go/pkg/models"
	"github.com/parulian1/belajar-go/pkg/utils"
)

var NewProduct models.Product

func ListProduct(w http.ResponseWriter, req *http.Request) {
	newBooks := models.ListProduct()
	resp, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DetailProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	productId := vars["productId"]
	prodId, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err)
	}
	productDetails, _ := models.DetailProduct(prodId)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	createProduct := &models.Product{}
	utils.ParseBody(req, createProduct)
	product := createProduct.CreateProduct()
	res, _ := json.Marshal(product)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func RemoveProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	productId := vars["productId"]
	pID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	product := models.RemoveProduct(pID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, req *http.Request) {
	updProduct := &models.Product{}
	utils.ParseBody(req, updProduct)
	vars := mux.Vars(req)
	productId := vars["productId"]
	pID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	productDetail, db := models.DetailProduct(pID)
	if updProduct.Name != productDetail.Name {
		productDetail.Name = updProduct.Name
	}
	if updProduct.Author != productDetail.Author {
		productDetail.Author = updProduct.Author
	}
	if updProduct.Publication != productDetail.Publication {
		productDetail.Publication = updProduct.Publication
	}
	db.Save(&productDetail)
	res, _ := json.Marshal(productDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
