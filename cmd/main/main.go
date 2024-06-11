package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joaorodrs/go-books/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3333", r))
}
