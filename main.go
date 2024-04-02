package main

import (
	"groupietracker/Artists"
	"groupietracker/autres_pages"
	"groupietracker/login"
	"groupietracker/api"
	"fmt"
	"net/url"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Groupie Tracker")

	menu := fyne.NewMainMenu(
		// Thème de la page
		fyne.NewMenu("Thèmes",
			fyne.NewMenuItem("Thème sombre", func() {
				myApp.Settings().SetTheme(theme.DarkTheme())
			}),
			fyne.NewMenuItem("Thème clair", func() {
				myApp.Settings().SetTheme(theme.LightTheme())
			}),
		),
		fyne.NewMenu("Filtres",
			fyne.NewMenuItem("Afficher les filtres", func() {
				checkbox1 := widget.NewCheck("Artiste", func(b bool) {
					fmt.Println(fmt.Sprintf("%t", b))
					if b == true {
						for _, artist := range Artists.GetArtist(){
							if len(artist.Members) == 1 {
								a := artist.Name
								fmt.Println(a)
							}
						}
					}

				})
				checkbox2 := widget.NewCheck("Groupe", func(b bool) {
					fmt.Println(fmt.Sprintf("%t", b))
					if b == true {
						for _, artist := range Artists.GetArtist(){
							if len(artist.Members) > 1{
							a := artist.Members
							fmt.Println(a)
							}
						}
					}
				})
				checkbox3 := widget.NewCheck("Date de création", func(b bool) {
					fmt.Println(fmt.Sprintf("%t", b))
					if b == true {
						for _, artist := range Artists.GetArtist(){
							a := artist.CreationDate
							fmt.Println(a)
						}
					}
				})
				checkbox4 := widget.NewCheck("Lieu de concert", func(b bool) {
					fmt.Println(fmt.Sprintf("%t", b))
					if b == true {
						for _, artist := range api.ApiLocations(){
							a := artist.Locations
							fmt.Println(a)
						}
					}
				})
				checkbox5 := widget.NewCheck("premier album", func(b bool) {
					fmt.Println(fmt.Sprintf("%t", b))
					if b == true {
						for _, artist := range Artists.GetArtist(){
							a := artist.FirstAlbum
							fmt.Println(a)
						}
					}
				})
				checkbox6 := widget.NewCheck("date de concert", func(b bool) {
					fmt.Println(fmt.Sprintf("%t", b))
					if b == true {
						for _, artist := range api.ApiConcertDates(){
							a := artist.ConcertDates
							fmt.Println(a)
						}
					}
				})

				checkboxes := container.NewVBox(
					checkbox1,
					checkbox2,
					checkbox3,
					checkbox4,
					checkbox5,
					checkbox6,
				)

				// retour_btn = widget.NewButton("Retour", func() {
				// 	v := *w
				// 	v.SetContent(splitPage)
				// 	w = &v
				// })

				w.SetContent(checkboxes)
			}),
		),
		fyne.NewMenu("En savoir plus",
			fyne.NewMenuItem("Spotify", func() {
				lien, _ := url.Parse("https://developer.spotify.com/documentation/embeds")
				_ = myApp.OpenURL(lien)
			}),
		),
	)

	w.SetMainMenu(menu)

	w.Resize(fyne.NewSize(1200, 600))
	autres_pages.PagePrincipale(&w)
	login.LoginPage(&w)

	w.ShowAndRun()
}
