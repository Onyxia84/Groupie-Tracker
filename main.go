package main

import (
	autres_pages "groupietracker/Autres_Pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()

	w := myApp.NewWindow("Groupie Tracker")

	w.Resize(fyne.NewSize(1600, 800))

	autres_pages.LoginPage(&w)

	w.ShowAndRun()
}
