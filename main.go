package main

import (
	//  "groupietracker/autres_pages"
	// "groupietracker/Filtres"
	"groupietracker/login"
	"net/url"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Groupie Tracker")

	menu := fyne.NewMainMenu(

		// Theme de le la page
		fyne.NewMenu("Thèmes",
			fyne.NewMenuItem("Thèmes sombre", func() {
				myApp.Settings().SetTheme(theme.DarkTheme())
			}),

			fyne.NewMenuItem("Thème clair", func() {
				myApp.Settings().SetTheme(theme.LightTheme())
			}),
		),

		fyne.NewMenu("En savoir plus",
			fyne.NewMenuItem("Spotify", func() {
				lien, _ := url.Parse("https://developer.spotify.com/documentation/embeds")
				_ = myApp.OpenURL(lien)
			}),
		))

	w.SetMainMenu(menu)


	checkbox1 := widget.NewCheck("Artiste", func(b bool) {
		fmt.Println(fmt.Sprintf("%t", b))
	})
	checkbox2 := widget.NewCheck("Groupe", func(b bool) {
		fmt.Println(fmt.Sprintf("%t", b))
	})
	checkbox3 := widget.NewCheck("Date de création", func(b bool) {
		fmt.Println(fmt.Sprintf("%t", b))
	})
	checkbox4 := widget.NewCheck("Lieu de concert", func(b bool) {
		fmt.Println(fmt.Sprintf("%t", b))
	})
	checkbox5 := widget.NewCheck("Nombre de Membre", func(b bool) {
		fmt.Println(fmt.Sprintf("%t", b))
	})

	checkboxes := container.NewVBox(
		checkbox1,
		checkbox2,
		checkbox3,
		checkbox4,
		checkbox5,
	)
	w.SetContent(checkboxes)

	w.Resize(fyne.NewSize(1600, 800))

	login.LoginPage(&w)

	w.ShowAndRun()
}
