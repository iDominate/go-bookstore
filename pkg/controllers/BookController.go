package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iDominate/go-bookstore/pkg/models"
	"github.com/iDominate/go-bookstore/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(models.GetAllBooks())
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := models.GetBookById(ID)
	response, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	bookResponse := book.CreateBook()
	res, _ := json.Marshal(bookResponse)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error While parsing")
	}
	book := models.DeleteBook(bookId)
	w.WriteHeader(http.StatusAccepted)
	res, _ := json.Marshal(book)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	updateBook := &models.Book{}
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error parsing id")
	}
	bookDetails, db := models.GetBookById(bookId)
	utils.ParseBody(r, updateBook)
	if updateBook.Title != "" {
		bookDetails.Title = updateBook.Title
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
