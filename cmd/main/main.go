package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iDominate/go-bookstore/pkg/models"
	"github.com/iDominate/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	models.InitDB()
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
