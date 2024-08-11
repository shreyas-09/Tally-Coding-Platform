package main

import (
	"fmt"
	"log"
	"net/http"
	"y/router"

	"github.com/rs/cors"
	// "github.com/gorilla/mux"
)

func main() {

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
