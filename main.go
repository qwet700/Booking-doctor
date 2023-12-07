package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ui "github.com/qwet700/Booking-doctor/UI"
	"github.com/qwet700/Booking-doctor/pkg/controller"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))
	// r.Use(middleware.Logger)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", controller.CreateUser)
		r.Get("/", controller.GetAllUsers)
		r.Get("/{id}", controller.GetUserID)
		r.Put("/{id}", controller.UpdateUser)
		// r.Delete("/{id}", controller.DeleteUser)
	})

	r.Route("/doctor", func(r chi.Router) {
		r.Post("/", controller.CreateDoc)
		// r.Get("/{docid}", controller.GetPatient)
		r.Put("/{docid}", controller.UpdateDoc)
		r.Delete("/{docid}", controller.DeleteDoc)
	})
	ui.RunUI()

	http.ListenAndServe(":8000", r)
}

// var client *http.Client

// type userFunc struct {
// 	Text string `json:"text"`
// }

// func getuserFunc() (userFunc, error) {
// 	var fact userFunc
// 	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
// 	if err != nil {
// 		return userFunc{}, err
// 	}

// 	defer resp.Body.Close()

// 	err = json.NewDecoder(resp.Body).Decode(&fact)
// 	if err != nil {
// 		return userFunc{}, err
// 	}

// 	return fact, nil
// }
