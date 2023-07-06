package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/parulian1/belajar-go/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterProductRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
