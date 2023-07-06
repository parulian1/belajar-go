package models

import (
	"github.com/jinzhu/gorm"
	"github.com/parulian1/belajar-go/pkg/config"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Name        string `json:"nama"`
	Author      string `json:"penulis"`
	Publication string `json:"penerbit"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (prod *Product) CreateProduct() *Product {
	db.NewRecord(prod)
	db.Create(&prod)
	return prod
}

func ListProduct() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func DetailProduct(productId int64) (*Product, *gorm.DB) {
	var detailProduct Product
	db := db.Where("ID=?", productId).Find(&detailProduct)
	return &detailProduct, db
}

func RemoveProduct(productId int64) Product {
	var remProduct Product
	db.Where("ID=?", productId).Delete(&remProduct)
	return remProduct
}
