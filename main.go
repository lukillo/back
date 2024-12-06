package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"email-marketing-api/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}