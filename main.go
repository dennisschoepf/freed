package main

import (
	"fmt"
	"log"
	"net/http"

	home "github.com/dennisschoepf/freed/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route Handlers
	r.HandleFunc("/", home.RouteHandler).Methods("GET")

	fmt.Println("--- Starting server at port 8080 ---")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
