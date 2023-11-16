package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	routes.ImpItem(r)
	http.Handle("/", r)
	fmt.Printf("Running on port 33060")
	log.Fatal(http.ListenAndServe("localhost:33060", r))
}
