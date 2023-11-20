package routes

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/qwet700/Booking-doctor/controller"
)

var Handle = func(router *chi.Router) {
	r := chi.NewRouter()

	r.Post("/user", controller.CreateUser)
	r.Get("/users", controller.GetAllUsers)
	r.Get("/user/{id}", controller.GetUsers)

	r.Use(middleware.Timeout(60 * time.Second))
}
