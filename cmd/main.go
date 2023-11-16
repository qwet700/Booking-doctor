package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter().StrictSlash(true)
	// routes.ImpItem(r)
	// http.Handle("/", r)
	// fmt.Printf("Running on port 33060")
	// log.Fatal(http.ListenAndServe("localhost:33060", r))

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
