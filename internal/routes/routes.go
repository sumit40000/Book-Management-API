package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sumit40000/book-management-api/internal/controllers"
)

func CreateRoutes() *mux.Router{
	// creating the router using the gorilla mux library 

	router := mux.NewRouter()
	//settting up the routes 
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1> Welcome to the Book Management API"))
	})
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST") // Creating a new book in the db
	router.HandleFunc("/api/books",controllers.GetAllBooks).Methods("GET") // getting the list of all books from the db 
	router.HandleFunc("/api/books/{id}", controllers.FindOneBook).Methods("GET") // getting a book with specified id from the db
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT") // updating a book with it's id 
	router.HandleFunc("/api/books/{id}", controllers.DeleteOneBook ).Methods("DELETE") // deleting a book with id 
	router.HandleFunc("/api/books", controllers.DeleteAllBooks).Methods("DELETE") // deleting all the books from the db

	return router
	
}