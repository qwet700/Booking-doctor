package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/qwet700/Booking-doctor/pkg/controller"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))
	// r.Use(middleware.Logger)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", controller.CreateUser)
		r.Get("/", controller.GetAllUsers)
		// r.Get("/{name}", controller.GetUserName)
		r.Put("/{id}", controller.UpdateUser)
		// r.Delete("/{id}", controller.DeleteUser)
	})

	r.Route("/doctor", func(r chi.Router) {
		r.Post("/", controller.CreateDoc)
		// r.Get("/{id}", controller.GetPatient)
		r.Put("/{docid}", controller.UpdateDoc)
		r.Delete("/{docid}", controller.DeleteDoc)
	})

	http.ListenAndServe(":8000", r)
}
