package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iDominate/go-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)

	router.HandleFunc("/book/", controllers.GetBooks).Methods(http.MethodGet)

	router.HandleFunc("/book/{id}", controllers.GetBook).Methods(http.MethodGet)

	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods(http.MethodPut)

	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
