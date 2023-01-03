package routes

import (
	"bookstore-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/", controllers.GetUser).Methods("GET")

	router.HandleFunc("/authors/", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/authors/", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{AuthorId}", controllers.GetAuthorById).Methods("GET")
	router.HandleFunc("/authors/{AuthorId}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{AuthorId}", controllers.UpdateAuthor).Methods("DELETE")

	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
