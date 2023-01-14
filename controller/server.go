package controller

import (
	"github.com/gorilla/mux"
	"golang-books-app/model"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	Repository *model.Repository
}

func NewController(repository *model.Repository) *Controller {
	return &Controller{repository}
}

var router *mux.Router

func Init(repository *model.Repository) {
	controller := NewController(repository)

	router = mux.NewRouter()
	initHandlers(controller)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started on port 8080")
	log.Fatal(srv.ListenAndServe())
}

func initHandlers(controller *Controller) {
	router.HandleFunc("/api/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controller.FindBook).Methods("GET")
	router.HandleFunc("/api/books", controller.CreateBook).Methods("POST")
	//router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	//router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
}
