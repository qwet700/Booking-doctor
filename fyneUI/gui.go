package ui

import (
	"context"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/qwet700/Booking-doctor/pkg/db"
	"github.com/qwet700/Booking-doctor/pkg/models"
)

func GUI() fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()

	submitButton := widget.NewButton("Submit", func() {
		name := nameEntry.Text
		phone := phoneEntry.Text
		age := ageEntry.Text

		// validate entry
		if name == "" || phone == "" || age == "" {
			return
		}

		ageInt, err := strconv.Atoi(age)
		if err != nil {
			log.Println(err)
			return
		}

		// add to MongoDB
		user := models.User{
			Name:  name,
			Phone: phone,
			Age:   ageInt,
		}
		if err := addUser(user); err != nil {
			log.Println(err)
			return
		}
		nameEntry.SetText("")
		phoneEntry.SetText("")
		ageEntry.SetText("")
	})

	content := container.NewVBox(
		widget.NewLabel("Name"),
		nameEntry,
		widget.NewLabel("Phone"),
		phoneEntry,
		widget.NewLabel("Age"),
		ageEntry,
		submitButton,
	)

	return content
}

var client = db.Dbconnect()

func addUser(user models.User) error {
	collection := client.Database("user").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}
