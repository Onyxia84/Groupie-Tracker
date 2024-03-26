package main

import (
	autres_pages "groupietracker/Autres_Pages"

	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
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

	w.Resize(fyne.NewSize(1600, 800))

	autres_pages.LoginPage(&w)

	w.ShowAndRun()
}
