package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ankit/go-book-managememnt/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) 
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) 
}
