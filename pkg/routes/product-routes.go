package routes

import (
	"github.com/gorilla/mux"
	"github.com/parulian1/belajar-go/pkg/controllers"
)

var RegisterProductRoutes = func(router *mux.Router) {
	router.HandleFunc("/product/", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/", controllers.ListProduct).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.DetailProduct).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{productId}", controllers.RemoveProduct).Methods("DELETE")
}
