package autres_pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"groupietracker/Artists"
	Getstruct "groupietracker/getstruct"
	"groupietracker/tools"
	"image/color"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PagePrincipale() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Hello")
	filtres := container.NewVBox()
	filtres2 := container.NewScroll(filtres)
	var artistButtons []*widget.Button // Garder une référence aux boutons ajoutés

	checkbox1 := widget.NewCheck("Artiste", func(b bool) {
		if b == true {
			for _, artist1 := range Artists.GetArtist() {
				if len(artist1.Members) == 1 {
					artistImageURL := artist1.Image
					artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
					buttonText := artist1.Name + " - Artiste "

					// Créez un bouton avec l'image et le texte
					// Créez un bouton avec l'image et le texte
					id := 0

					imageButton := widget.NewButton("", func() {
						for i := 0; i < 52; i++ {
							if buttonText == Artists.GetArtist()[i].Name+" - Artiste " {
								id = i
								break
							}
							id = i
						}

						// CREATION DES CANVAS
						// Récupérer l'image de l'artiste

						tempo, _ := http.Get(Artists.GetArtist()[id].Image)
						contentImage := canvas.NewImageFromReader(tempo.Body, "image")
						contentImage.FillMode = canvas.ImageFillOriginal

						// Créer et initliaser les textes

						contentName := canvas.NewText(Artists.GetArtist()[id].Name, color.RGBA{255, 255, 255, 1})
						textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
						contentmember := canvas.NewText(""+tools.StringAppend(Artists.GetArtist()[id].Members), color.RGBA{255, 255, 255, 1})
						textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
						contentCreationDate := canvas.NewText(strconv.Itoa(Artists.GetArtist()[id].CreationDate), color.RGBA{255, 255, 255, 1})
						textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
						contentFirstAlbum := canvas.NewText(""+Artists.GetArtist()[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
						myMapRelation := Artists.GetLocationsRelation(Artists.GetArtist(), id)
						textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
						contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

						// FIN DE CREATION DES CANVAS

						// Styliser les canvas

						contentName.Alignment = fyne.TextAlignCenter
						contentName.TextStyle = fyne.TextStyle{Bold: true}
						contentName.TextSize = 30

						textMember.TextSize = 20
						textMember.Alignment = fyne.TextAlignCenter
						textMember.TextStyle = fyne.TextStyle{Bold: true}

						contentmember.TextSize = 18
						contentmember.Alignment = fyne.TextAlignCenter

						textCreationDate.TextSize = 20
						textCreationDate.Alignment = fyne.TextAlignCenter
						textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

						contentCreationDate.TextSize = 18
						contentCreationDate.Alignment = fyne.TextAlignCenter

						textAlbum.TextSize = 20
						textAlbum.Alignment = fyne.TextAlignCenter
						textAlbum.TextStyle = fyne.TextStyle{Bold: true}

						contentFirstAlbum.TextSize = 18
						contentFirstAlbum.Alignment = fyne.TextAlignCenter

						textLocation.TextSize = 20
						textLocation.Alignment = fyne.TextAlignCenter
						textLocation.TextStyle = fyne.TextStyle{Bold: true}

						contentLocation.Wrapping = fyne.TextWrapWord
						contentLocation.Alignment = fyne.TextAlignCenter

						// Organiser les canvas et containers
						memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
						creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
						albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
						locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

						// TOP PAGE

						artistNameAndImage := container.NewVBox(contentName, contentImage)
						spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer())
						pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

						// MIDDLE PAGE

						rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

						// END MIDDLE PAGE
						// BOTTOM PAGE

						// END BOTTOM PAGE

						content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
						contentt := container.NewVScroll(content)

						// FIN ORGANISATION

						w.SetContent(contentt)
					})
					imageButton.Importance = widget.LowImportance
					imageButton.SetIcon(artistImageResource)
					imageButton.SetText(buttonText)
					filtres.Add(imageButton)
					artistButtons = append(artistButtons, imageButton)
				}
			}
			filtres2.Show()
		} else {
			filtres2.Hide()
		}
	})

	checkbox2 := widget.NewCheck("Groupe", func(b bool) {
		if b == true {
			for _, artist := range Artists.GetArtist() {
				if len(artist.Members) > 2 {
					artistImageURL := artist.Image
					artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
					buttonText := artist.Name + " - Groupe "

					// Créez un widget d'image
					// Créez un bouton avec l'image et le texte
					id := 0

					imageButton := widget.NewButton("", func() {
						for i := 0; i < 52; i++ {
							if buttonText == Artists.GetArtist()[i].Name+" - Groupe " {
								id = i
								break
							}
							id = i
						}

						// CREATION DES CANVAS
						// Récupérer l'image de l'artiste

						tempo, _ := http.Get(Artists.GetArtist()[id].Image)
						contentImage := canvas.NewImageFromReader(tempo.Body, "image")
						contentImage.FillMode = canvas.ImageFillOriginal

						// Créer et initliaser les textes

						contentName := canvas.NewText(Artists.GetArtist()[id].Name, color.RGBA{255, 255, 255, 1})
						textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
						contentmember := canvas.NewText(""+tools.StringAppend(Artists.GetArtist()[id].Members), color.RGBA{255, 255, 255, 1})
						textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
						contentCreationDate := canvas.NewText(strconv.Itoa(Artists.GetArtist()[id].CreationDate), color.RGBA{255, 255, 255, 1})
						textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
						contentFirstAlbum := canvas.NewText(""+Artists.GetArtist()[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
						myMapRelation := Artists.GetLocationsRelation(Artists.GetArtist(), id)
						textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
						contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

						// FIN DE CREATION DES CANVAS

						// Styliser les canvas

						contentName.Alignment = fyne.TextAlignCenter
						contentName.TextStyle = fyne.TextStyle{Bold: true}
						contentName.TextSize = 30

						textMember.TextSize = 20
						textMember.Alignment = fyne.TextAlignCenter
						textMember.TextStyle = fyne.TextStyle{Bold: true}

						contentmember.TextSize = 18
						contentmember.Alignment = fyne.TextAlignCenter

						textCreationDate.TextSize = 20
						textCreationDate.Alignment = fyne.TextAlignCenter
						textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

						contentCreationDate.TextSize = 18
						contentCreationDate.Alignment = fyne.TextAlignCenter

						textAlbum.TextSize = 20
						textAlbum.Alignment = fyne.TextAlignCenter
						textAlbum.TextStyle = fyne.TextStyle{Bold: true}

						contentFirstAlbum.TextSize = 18
						contentFirstAlbum.Alignment = fyne.TextAlignCenter

						textLocation.TextSize = 20
						textLocation.Alignment = fyne.TextAlignCenter
						textLocation.TextStyle = fyne.TextStyle{Bold: true}

						contentLocation.Wrapping = fyne.TextWrapWord
						contentLocation.Alignment = fyne.TextAlignCenter

						// Organiser les canvas et containers
						memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
						creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
						albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
						locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

						// TOP PAGE

						artistNameAndImage := container.NewVBox(contentName, contentImage)
						spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer())
						pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

						// MIDDLE PAGE

						rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

						// END MIDDLE PAGE
						// BOTTOM PAGE

						// END BOTTOM PAGE

						content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
						contentt := container.NewVScroll(content)

						// FIN ORGANISATION

						w.SetContent(contentt)
					})
					imageButton.Importance = widget.LowImportance
					imageButton.SetIcon(artistImageResource)
					imageButton.SetText(buttonText)
					filtres.Add(imageButton)
				}
			}
			filtres2.Show()
		} else {
			filtres2.Hide()
		}
	})
	checkbox3 := widget.NewCheck("Date de création", func(b bool) {
		if b == true {
			for _, artist := range Artists.GetArtist() {
				artistImageURL := artist.Image
				artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
				buttonText := strconv.Itoa(artist.CreationDate) + " - Date de création "

				id := 0

				imageButton := widget.NewButton("", func() {
					for i := 0; i < 52; i++ {
						if buttonText == strconv.Itoa(Artists.GetArtist()[i].CreationDate)+" - Date de création " {
							id = i
							break
						}
						id = i
					}

					// CREATION DES CANVAS
					// Récupérer l'image de l'artiste

					tempo, _ := http.Get(Artists.GetArtist()[id].Image)
					contentImage := canvas.NewImageFromReader(tempo.Body, "image")
					contentImage.FillMode = canvas.ImageFillOriginal

					// Créer et initliaser les textes

					contentName := canvas.NewText(Artists.GetArtist()[id].Name, color.RGBA{255, 255, 255, 1})
					textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
					contentmember := canvas.NewText(""+tools.StringAppend(Artists.GetArtist()[id].Members), color.RGBA{255, 255, 255, 1})
					textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
					contentCreationDate := canvas.NewText(strconv.Itoa(Artists.GetArtist()[id].CreationDate), color.RGBA{255, 255, 255, 1})
					textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
					contentFirstAlbum := canvas.NewText(""+Artists.GetArtist()[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
					myMapRelation := Artists.GetLocationsRelation(Artists.GetArtist(), id)
					textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
					contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

					// FIN DE CREATION DES CANVAS

					// Styliser les canvas

					contentName.Alignment = fyne.TextAlignCenter
					contentName.TextStyle = fyne.TextStyle{Bold: true}
					contentName.TextSize = 30

					textMember.TextSize = 20
					textMember.Alignment = fyne.TextAlignCenter
					textMember.TextStyle = fyne.TextStyle{Bold: true}

					contentmember.TextSize = 18
					contentmember.Alignment = fyne.TextAlignCenter

					textCreationDate.TextSize = 20
					textCreationDate.Alignment = fyne.TextAlignCenter
					textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

					contentCreationDate.TextSize = 18
					contentCreationDate.Alignment = fyne.TextAlignCenter

					textAlbum.TextSize = 20
					textAlbum.Alignment = fyne.TextAlignCenter
					textAlbum.TextStyle = fyne.TextStyle{Bold: true}

					contentFirstAlbum.TextSize = 18
					contentFirstAlbum.Alignment = fyne.TextAlignCenter

					textLocation.TextSize = 20
					textLocation.Alignment = fyne.TextAlignCenter
					textLocation.TextStyle = fyne.TextStyle{Bold: true}

					contentLocation.Wrapping = fyne.TextWrapWord
					contentLocation.Alignment = fyne.TextAlignCenter

					// Organiser les canvas et containers
					memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
					creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
					albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
					locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

					// TOP PAGE

					artistNameAndImage := container.NewVBox(contentName, contentImage)
					spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer())
					pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

					// MIDDLE PAGE

					rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

					// END MIDDLE PAGE
					// BOTTOM PAGE

					// END BOTTOM PAGE

					content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
					contentt := container.NewVScroll(content)

					// FIN ORGANISATION

					w.SetContent(contentt)
				})
				imageButton.Importance = widget.LowImportance
				imageButton.SetIcon(artistImageResource)
				imageButton.SetText(buttonText)
				filtres.Add(imageButton)
			}
			filtres2.Show()
		} else {
			filtres2.Hide()
		}
	})

	checkbox4 := widget.NewCheck("premier album", func(b bool) {
		if b == true {
			for _, artist1 := range Artists.GetArtist() {
				artistImageURL := artist1.Image
				artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
				buttonText := artist1.FirstAlbum + " - Premier Albulm "

				// Créez un widget d'image
				// Créez un bouton avec l'image et le texte
				id := 0
				// Créez un bouton avec l'image et le texte
				imageButton := widget.NewButton("", func() {
					for i := 0; i < 52; i++ {
						if buttonText == Artists.GetArtist()[i].FirstAlbum+" - Premier Albulm " {
							id = i
							break
						}
						id = i
					}

					// CREATION DES CANVAS
					// Récupérer l'image de l'artiste

					tempo, _ := http.Get(Artists.GetArtist()[id].Image)
					contentImage := canvas.NewImageFromReader(tempo.Body, "image")
					contentImage.FillMode = canvas.ImageFillOriginal

					// Créer et initliaser les textes

					contentName := canvas.NewText(Artists.GetArtist()[id].Name, color.RGBA{255, 255, 255, 1})
					textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
					contentmember := canvas.NewText(""+tools.StringAppend(Artists.GetArtist()[id].Members), color.RGBA{255, 255, 255, 1})
					textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
					contentCreationDate := canvas.NewText(strconv.Itoa(Artists.GetArtist()[id].CreationDate), color.RGBA{255, 255, 255, 1})
					textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
					contentFirstAlbum := canvas.NewText(""+Artists.GetArtist()[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
					myMapRelation := Artists.GetLocationsRelation(Artists.GetArtist(), id)
					textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
					contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

					// FIN DE CREATION DES CANVAS

					// Styliser les canvas

					contentName.Alignment = fyne.TextAlignCenter
					contentName.TextStyle = fyne.TextStyle{Bold: true}
					contentName.TextSize = 30

					textMember.TextSize = 20
					textMember.Alignment = fyne.TextAlignCenter
					textMember.TextStyle = fyne.TextStyle{Bold: true}

					contentmember.TextSize = 18
					contentmember.Alignment = fyne.TextAlignCenter

					textCreationDate.TextSize = 20
					textCreationDate.Alignment = fyne.TextAlignCenter
					textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

					contentCreationDate.TextSize = 18
					contentCreationDate.Alignment = fyne.TextAlignCenter

					textAlbum.TextSize = 20
					textAlbum.Alignment = fyne.TextAlignCenter
					textAlbum.TextStyle = fyne.TextStyle{Bold: true}

					contentFirstAlbum.TextSize = 18
					contentFirstAlbum.Alignment = fyne.TextAlignCenter

					textLocation.TextSize = 20
					textLocation.Alignment = fyne.TextAlignCenter
					textLocation.TextStyle = fyne.TextStyle{Bold: true}

					contentLocation.Wrapping = fyne.TextWrapWord
					contentLocation.Alignment = fyne.TextAlignCenter

					// Organiser les canvas et containers
					memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
					creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
					albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
					locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

					// TOP PAGE

					artistNameAndImage := container.NewVBox(contentName, contentImage)
					spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer())
					pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

					// MIDDLE PAGE

					rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

					// END MIDDLE PAGE
					// BOTTOM PAGE

					// END BOTTOM PAGE

					content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
					contentt := container.NewVScroll(content)

					// FIN ORGANISATION

					w.SetContent(contentt)
				})
				imageButton.Importance = widget.LowImportance
				imageButton.SetIcon(artistImageResource)
				imageButton.SetText(buttonText)
				filtres.Add(imageButton)
			}
			filtres2.Show()
		} else {
			filtres2.Hide()
		}
	})

	filtres2.SetMinSize(fyne.NewSize(300, 300))
	// Ajout des checkboxes à un conteneur VBox
	checkboxes := container.NewVBox(
		checkbox1,
		checkbox2,
		checkbox3,
		checkbox4,
		filtres2,
		// Ajoutez d'autres checkboxes si nécessaire
	)

	// Créez le menu avec les checkboxes
	menu := fyne.NewMenu("Filtres",
		fyne.NewMenuItem("Afficher les filtres", func() {
			// Affiche les checkboxes lorsque l'élément de menu est sélectionné
			// Remplacez les commentaires suivants par votre propre logique pour afficher les checkboxes

			w.SetContent(checkboxes)
		}),
	)
	menu2 := fyne.NewMenu("En savoir plus",
		fyne.NewMenuItem("Spotify", func() {
			lien, _ := url.Parse("https://developer.spotify.com/documentation/embeds")
			_ = a.OpenURL(lien)
		}),
	)

	menu1 := fyne.NewMenu("Thèmes",
		fyne.NewMenuItem("Thème sombre", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		fyne.NewMenuItem("Thème clair", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	// Ajoutez le menu à la fenêtre
	w.SetMainMenu(fyne.NewMainMenu(menu, menu1, menu2))
	// Déclaration des variables

	var artist []Getstruct.Artist
	artist = Artists.GetArtist()
	var listR []int
	var listVierge []int
	var n_recherche string

	// Création de la listView (gauche)

	listView := widget.NewList(
		// first argument is item count</div>
		// len() function to get myStudentData slice len</div>
		func() int { return len(artist) },
		// 2nd argument is for widget choice. I want to use label</div>
		func() fyne.CanvasObject { return widget.NewLabel("") },
		// 3rd argument is to update widget with our new data</div>
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(artist[lii].Name)
		},
	)

	// Création de la recherche (droite)

	searchText := canvas.NewText("Recherchez quelque chose içi : ", color.White)
	e_recherche := widget.NewEntry()
	e_recherche.SetPlaceHolder("Cliquez ici pour rechercher :")
	retour_btn := widget.NewButton("Retour", func() {
	})
	subt_btn := widget.NewButton("", func() {
	})

	// Dans la fonction pagePrincipale

	// Bouton "rechercher" avec la fonctionnalité rechercher

	subt_btn = widget.NewButton("Rechercher", func() {
		n_recherche = e_recherche.Text
		n_recherche = tools.ToLower(n_recherche)
		for comID := 0; len(artist) > comID; comID++ {
			if tools.Recherche(artist[comID].Name, n_recherche) {
				listR = append(listR, comID)
			}
		}

		// Création de la mini listView déroulante en fonction de la recherche

		listViewR := widget.NewList(
			// first argument is item count</div>
			// len() function to get myStudentData slice len</div>
			func() int { return len(listR) },
			// 2nd argument is for widget choice. I want to use label</div>
			func() fyne.CanvasObject { return widget.NewLabel("") },
			// 3rd argument is to update widget with our new data</div>
			func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(artist[listR[lii]].Name)
			},
		)

		// Disposition de la page lorsque une option de la mini listView est activée

		listViewR.OnSelected = func(id widget.ListItemID) {

			// Disposition de la page par défaut

			searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
			searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone)
			searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
			searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

			splitPage := container.NewHSplit(
				listView,
				searchZoneFinal)
			splitPage.Offset = 0.2

			retour_btn = widget.NewButton("Retour", func() {
				w.SetContent(splitPage)
			})

			// CREATION DES CANVAS

			// Récupérer l'image de l'artiste

			tempo, _ := http.Get(artist[listR[id]].Image)
			contentImage := canvas.NewImageFromReader(tempo.Body, "image")
			contentImage.FillMode = canvas.ImageFillOriginal

			// Créer et initliaser les textes

			contentName := canvas.NewText(artist[listR[id]].Name, color.RGBA{255, 255, 255, 1})
			textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
			contentmember := canvas.NewText(""+tools.StringAppend(artist[listR[id]].Members), color.RGBA{255, 255, 255, 1})
			textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
			contentCreationDate := canvas.NewText(strconv.Itoa(artist[listR[id]].CreationDate), color.RGBA{255, 255, 255, 1})
			textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
			contentFirstAlbum := canvas.NewText(""+artist[listR[id]].FirstAlbum, color.RGBA{255, 255, 255, 1})
			myMapRelation := Artists.GetLocationsRelation(artist, listR[id])
			textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
			contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

			// FIN DE CREATION DES CANVAS

			// Styliser les canvas

			contentName.Alignment = fyne.TextAlignCenter
			contentName.TextStyle = fyne.TextStyle{Bold: true}
			contentName.TextSize = 30

			textMember.TextSize = 20
			textMember.Alignment = fyne.TextAlignCenter
			textMember.TextStyle = fyne.TextStyle{Bold: true}

			contentmember.TextSize = 18
			contentmember.Alignment = fyne.TextAlignCenter

			textCreationDate.TextSize = 20
			textCreationDate.Alignment = fyne.TextAlignCenter
			textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

			contentCreationDate.TextSize = 18
			contentCreationDate.Alignment = fyne.TextAlignCenter

			textAlbum.TextSize = 20
			textAlbum.Alignment = fyne.TextAlignCenter
			textAlbum.TextStyle = fyne.TextStyle{Bold: true}

			contentFirstAlbum.TextSize = 18
			contentFirstAlbum.Alignment = fyne.TextAlignCenter

			textLocation.TextSize = 20
			textLocation.Alignment = fyne.TextAlignCenter
			textLocation.TextStyle = fyne.TextStyle{Bold: true}

			contentLocation.Wrapping = fyne.TextWrapWord
			contentLocation.Alignment = fyne.TextAlignCenter

			// Organiser les canvas et containers
			memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
			creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
			albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
			locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

			// TOP PAGE

			artistNameAndImage := container.NewVBox(contentName, contentImage)
			spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), retour_btn)
			pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

			// MIDDLE PAGE

			rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

			// END MIDDLE PAGE
			// BOTTOM PAGE

			retour_btn = widget.NewButton("Effacer les résultats de la recherche", func() {
				w.SetContent(splitPage)
			})

			// END BOTTOM PAGE

			content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
			contentt := container.NewVScroll(content)

			// FIN ORGANISATION

			split := container.NewHSplit(
				listView,
				contentt)
			split.Offset = 0.2

			w.SetContent(split)

		}

		// Réinitialiser la page de droite et effacer les résultats de la recherche

		retour_btn = widget.NewButton("Effacer les résultats de la recherche", func() {

			sugg := container.NewVBox()
			sugg2 := container.NewScroll(sugg)
			e_recherche.OnChanged = func(query string) {
				searchText1 := strings.ToLower(query)
				if len(query) > 0 {
					sugg.Objects = nil

					for _, artist1 := range Artists.GetArtist() {
						if strings.Contains(strings.ToLower(artist1.Name), searchText1) {
							artistImageURL := artist1.Image
							artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
							buttonText := artist1.Name + " - Groupe "

							// Créez un widget d'image
							id := 0
							// Créez un bouton avec l'image et le texte
							imageButton := widget.NewButton("", func() {
								for i := 0; i < 52; i++ {
									if buttonText == artist[i].Name+" - Groupe " {
										id = i
										break
									}
									id = i
								}
								searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
								sugg2.SetMinSize(fyne.NewSize(200, 200))
								searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone, sugg2)
								searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
								searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

								splitPage := container.NewHSplit(
									listView,
									searchZoneFinal)
								splitPage.Offset = 0.2

								retour_btn = widget.NewButton("Retour", func() {
									w.SetContent(splitPage)
									w.SetFullScreen(false)
									w.SetFullScreen(true)
								})

								// CREATION DES CANVAS

								// Récupérer l'image de l'artiste

								tempo, _ := http.Get(artist[id].Image)
								contentImage := canvas.NewImageFromReader(tempo.Body, "image")
								contentImage.FillMode = canvas.ImageFillOriginal

								// Créer et initliaser les textes

								contentName := canvas.NewText(artist[id].Name, color.RGBA{255, 255, 255, 1})
								textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
								contentmember := canvas.NewText(""+tools.StringAppend(artist[id].Members), color.RGBA{255, 255, 255, 1})
								textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
								contentCreationDate := canvas.NewText(strconv.Itoa(artist[id].CreationDate), color.RGBA{255, 255, 255, 1})
								textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
								contentFirstAlbum := canvas.NewText(""+artist[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
								myMapRelation := Artists.GetLocationsRelation(artist, id)
								textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
								contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

								// FIN DE CREATION DES CANVAS

								// Styliser les canvas

								contentName.Alignment = fyne.TextAlignCenter
								contentName.TextStyle = fyne.TextStyle{Bold: true}
								contentName.TextSize = 30

								textMember.TextSize = 20
								textMember.Alignment = fyne.TextAlignCenter
								textMember.TextStyle = fyne.TextStyle{Bold: true}

								contentmember.TextSize = 18
								contentmember.Alignment = fyne.TextAlignCenter

								textCreationDate.TextSize = 20
								textCreationDate.Alignment = fyne.TextAlignCenter
								textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

								contentCreationDate.TextSize = 18
								contentCreationDate.Alignment = fyne.TextAlignCenter

								textAlbum.TextSize = 20
								textAlbum.Alignment = fyne.TextAlignCenter
								textAlbum.TextStyle = fyne.TextStyle{Bold: true}

								contentFirstAlbum.TextSize = 18
								contentFirstAlbum.Alignment = fyne.TextAlignCenter

								textLocation.TextSize = 20
								textLocation.Alignment = fyne.TextAlignCenter
								textLocation.TextStyle = fyne.TextStyle{Bold: true}

								contentLocation.Wrapping = fyne.TextWrapWord
								contentLocation.Alignment = fyne.TextAlignCenter

								// Organiser les canvas et containers
								memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
								creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
								albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
								locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

								// TOP PAGE

								artistNameAndImage := container.NewVBox(contentName, contentImage)
								spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), retour_btn)
								pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

								// MIDDLE PAGE

								rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

								// END MIDDLE PAGE
								// BOTTOM PAGE

								retour_btn = widget.NewButton("Effacer les résultats de la recherche", func() {
									w.SetContent(splitPage)
								})

								// END BOTTOM PAGE

								content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
								contentt := container.NewVScroll(content)

								// FIN ORGANISATION

								split := container.NewHSplit(
									listView,
									contentt)
								split.Offset = 0.2

								w.SetContent(split)

							})

							// Ajoutez l'image et le texte au bouton
							imageButton.Importance = widget.LowImportance
							imageButton.SetIcon(artistImageResource)
							imageButton.SetText(buttonText)
							sugg.Add(imageButton)
						}
					}
					sugg.Show()
					sugg2.Show()
				} else {
				}
			}

			searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
			sugg2.SetMinSize(fyne.NewSize(200, 200))
			searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone, sugg2)
			searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
			searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())
			// Affichage
			splitPage := container.NewHSplit(
				listView,
				searchZoneFinal)
			splitPage.Offset = 0.2

			w.SetContent(splitPage)
			listR = listVierge
		})

		// Disposition de la page de droite

		splitPage := container.NewHSplit(
			listView,
			container.NewVBox(e_recherche, retour_btn, listViewR))
		splitPage.Offset = 0.2

		// Affichage

		w.SetContent(splitPage)
	})

	// Disposition de la page lorsque la listView (gauche) est activée

	listView.OnSelected = func(id widget.ListItemID) {

		// Disposition de la page par défaut
		sugg := container.NewVBox()
		sugg2 := container.NewScroll(sugg)
		e_recherche.OnChanged = func(query string) {
			searchText1 := strings.ToLower(query)
			if len(query) > 0 {
				sugg.Objects = nil
				for _, artist1 := range Artists.GetArtist() {
					if strings.Contains(strings.ToLower(artist1.Name), searchText1) {
						artistImageURL := artist1.Image
						artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
						buttonText := artist1.Name + " - Groupe "

						// Créez un widget d'image
						id := 0
						// Créez un bouton avec l'image et le texte
						imageButton := widget.NewButton("", func() {
							for i := 0; i < 52; i++ {
								if buttonText == artist[i].Name+" - Groupe " {
									id = i
									break
								}
								id = i
							}
							searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
							sugg2.SetMinSize(fyne.NewSize(200, 200))
							searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone, sugg2)
							searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
							searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

							splitPage := container.NewHSplit(
								listView,
								searchZoneFinal)
							splitPage.Offset = 0.2

							retour_btn = widget.NewButton("Retour", func() {
								w.SetContent(splitPage)
								w.SetFullScreen(false)
								w.SetFullScreen(true)
							})

							// CREATION DES CANVAS

							// Récupérer l'image de l'artiste

							tempo, _ := http.Get(artist[id].Image)
							contentImage := canvas.NewImageFromReader(tempo.Body, "image")
							contentImage.FillMode = canvas.ImageFillOriginal

							// Créer et initliaser les textes

							contentName := canvas.NewText(artist[id].Name, color.RGBA{255, 255, 255, 1})
							textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
							contentmember := canvas.NewText(""+tools.StringAppend(artist[id].Members), color.RGBA{255, 255, 255, 1})
							textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
							contentCreationDate := canvas.NewText(strconv.Itoa(artist[id].CreationDate), color.RGBA{255, 255, 255, 1})
							textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
							contentFirstAlbum := canvas.NewText(""+artist[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
							myMapRelation := Artists.GetLocationsRelation(artist, id)
							textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
							contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

							// FIN DE CREATION DES CANVAS

							// Styliser les canvas

							contentName.Alignment = fyne.TextAlignCenter
							contentName.TextStyle = fyne.TextStyle{Bold: true}
							contentName.TextSize = 30

							textMember.TextSize = 20
							textMember.Alignment = fyne.TextAlignCenter
							textMember.TextStyle = fyne.TextStyle{Bold: true}

							contentmember.TextSize = 18
							contentmember.Alignment = fyne.TextAlignCenter

							textCreationDate.TextSize = 20
							textCreationDate.Alignment = fyne.TextAlignCenter
							textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

							contentCreationDate.TextSize = 18
							contentCreationDate.Alignment = fyne.TextAlignCenter

							textAlbum.TextSize = 20
							textAlbum.Alignment = fyne.TextAlignCenter
							textAlbum.TextStyle = fyne.TextStyle{Bold: true}

							contentFirstAlbum.TextSize = 18
							contentFirstAlbum.Alignment = fyne.TextAlignCenter

							textLocation.TextSize = 20
							textLocation.Alignment = fyne.TextAlignCenter
							textLocation.TextStyle = fyne.TextStyle{Bold: true}

							contentLocation.Wrapping = fyne.TextWrapWord
							contentLocation.Alignment = fyne.TextAlignCenter

							// Organiser les canvas et containers
							memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
							creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
							albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
							locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

							// TOP PAGE

							artistNameAndImage := container.NewVBox(contentName, contentImage)
							spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), retour_btn)
							pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

							// MIDDLE PAGE

							rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

							// END MIDDLE PAGE
							// BOTTOM PAGE

							retour_btn = widget.NewButton("Effacer les résultats de la recherche", func() {
								w.SetContent(splitPage)
							})

							// END BOTTOM PAGE

							content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
							contentt := container.NewVScroll(content)

							// FIN ORGANISATION

							split := container.NewHSplit(
								listView,
								contentt)
							split.Offset = 0.2

							w.SetContent(split)
						})

						// Ajoutez l'image et le texte au bouton
						imageButton.Importance = widget.LowImportance
						imageButton.SetIcon(artistImageResource)
						imageButton.SetText(buttonText)
						sugg.Add(imageButton)
					}
				}
			} else {
			}
		}

		searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
		sugg2.SetMinSize(fyne.NewSize(200, 200))
		searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone, sugg2)
		searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
		searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

		splitPage := container.NewHSplit(
			listView,
			searchZoneFinal)
		splitPage.Offset = 0.2

		retour_btn = widget.NewButton("Retour", func() {

			// Affichage

			w.SetContent(splitPage)
			w.SetFullScreen(false)
			w.SetFullScreen(true)
		})

		// CREATION DES CANVAS

		// Récupérer l'image de l'artiste

		tempo, _ := http.Get(artist[id].Image)
		contentImage := canvas.NewImageFromReader(tempo.Body, "image")
		contentImage.FillMode = canvas.ImageFillOriginal

		// Créer et initialiaser les textes

		contentName := canvas.NewText(artist[id].Name, color.RGBA{255, 255, 255, 1})
		textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
		contentmember := canvas.NewText(""+tools.StringAppend(artist[id].Members), color.RGBA{255, 255, 255, 1})
		textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
		contentCreationDate := canvas.NewText(strconv.Itoa(artist[id].CreationDate), color.RGBA{255, 255, 255, 1})
		textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
		contentFirstAlbum := canvas.NewText(""+artist[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
		myMapRelation := Artists.GetLocationsRelation(artist, id)
		textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
		contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

		// FIN DE CREATION DES CANVAS

		// Styliser les canvas

		contentName.Alignment = fyne.TextAlignCenter
		contentName.TextStyle = fyne.TextStyle{Bold: true}
		contentName.TextSize = 30

		textMember.TextSize = 20
		textMember.Alignment = fyne.TextAlignCenter
		textMember.TextStyle = fyne.TextStyle{Bold: true}

		contentmember.TextSize = 15
		contentmember.Alignment = fyne.TextAlignCenter

		textCreationDate.TextSize = 20
		textCreationDate.Alignment = fyne.TextAlignCenter
		textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

		contentCreationDate.TextSize = 15
		contentCreationDate.Alignment = fyne.TextAlignCenter

		textAlbum.TextSize = 20
		textAlbum.Alignment = fyne.TextAlignCenter
		textAlbum.TextStyle = fyne.TextStyle{Bold: true}

		contentFirstAlbum.TextSize = 15
		contentFirstAlbum.Alignment = fyne.TextAlignCenter

		textLocation.TextSize = 20
		textLocation.Alignment = fyne.TextAlignCenter
		textLocation.TextStyle = fyne.TextStyle{Bold: true}

		contentLocation.Wrapping = fyne.TextWrapWord
		contentLocation.Alignment = fyne.TextAlignCenter

		// Organiser les canvas et containers
		memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
		creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
		albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
		locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

		// TOP PAGE

		artistNameAndImage := container.NewVBox(contentName, contentImage)
		spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), retour_btn)
		pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

		// MIDDLE PAGE

		rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

		// END MIDDLE PAGE
		// BOTTOM PAGE

		// END BOTTOM PAGE

		content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
		contentt := container.NewVScroll(content)

		// FIN ORGANISATION

		split := container.NewHSplit(
			listView,
			contentt)
		split.Offset = 0.2

		// Affichage

		w.SetContent(split)
	}

	// Disposition de la page par défaut

	sugg := container.NewVBox()
	sugg2 := container.NewScroll(sugg)
	e_recherche.OnChanged = func(query string) {
		searchText1 := strings.ToLower(query)
		if len(query) > 0 {
			sugg.Objects = nil

			for _, artist1 := range Artists.GetArtist() {
				if strings.Contains(strings.ToLower(artist1.Name), searchText1) {
					artistImageURL := artist1.Image
					artistImageResource, _ := fyne.LoadResourceFromURLString(artistImageURL)
					buttonText := artist1.Name + " - Groupe "

					// Créez un widget d'image
					id := 0
					// Créez un bouton avec l'image et le texte
					imageButton := widget.NewButton("", func() {
						for i := 0; i < 52; i++ {
							if buttonText == artist[i].Name+" - Groupe " {
								id = i
								break
							}
							id = i
						}
						searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
						sugg2.SetMinSize(fyne.NewSize(200, 200))
						searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone, sugg2)
						searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
						searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

						splitPage := container.NewHSplit(
							listView,
							searchZoneFinal)
						splitPage.Offset = 0.2

						retour_btn = widget.NewButton("Retour", func() {
							w.SetContent(splitPage)
							w.SetFullScreen(false)
							w.SetFullScreen(true)
						})

						// CREATION DES CANVAS

						// Récupérer l'image de l'artiste

						tempo, _ := http.Get(artist[id].Image)
						contentImage := canvas.NewImageFromReader(tempo.Body, "image")
						contentImage.FillMode = canvas.ImageFillOriginal

						// Créer et initliaser les textes

						contentName := canvas.NewText(artist[id].Name, color.RGBA{255, 255, 255, 1})
						textMember := canvas.NewText("Liste des membres : ", color.RGBA{255, 0, 0, 1})
						contentmember := canvas.NewText(""+tools.StringAppend(artist[id].Members), color.RGBA{255, 255, 255, 1})
						textCreationDate := canvas.NewText("Date de création : ", color.RGBA{255, 0, 0, 1})
						contentCreationDate := canvas.NewText(strconv.Itoa(artist[id].CreationDate), color.RGBA{255, 255, 255, 1})
						textAlbum := canvas.NewText("Premier album publié en : ", color.RGBA{255, 0, 0, 1})
						contentFirstAlbum := canvas.NewText(""+artist[id].FirstAlbum, color.RGBA{255, 255, 255, 1})
						myMapRelation := Artists.GetLocationsRelation(artist, id)
						textLocation := canvas.NewText("Dates et lieux de concerts : ", color.RGBA{255, 0, 0, 1})
						contentLocation := widget.NewLabel("" + tools.StringAppend(tools.MapString(myMapRelation.Locations)))

						// FIN DE CREATION DES CANVAS

						// Styliser les canvas

						contentName.Alignment = fyne.TextAlignCenter
						contentName.TextStyle = fyne.TextStyle{Bold: true}
						contentName.TextSize = 30

						textMember.TextSize = 20
						textMember.Alignment = fyne.TextAlignCenter
						textMember.TextStyle = fyne.TextStyle{Bold: true}

						contentmember.TextSize = 18
						contentmember.Alignment = fyne.TextAlignCenter

						textCreationDate.TextSize = 20
						textCreationDate.Alignment = fyne.TextAlignCenter
						textCreationDate.TextStyle = fyne.TextStyle{Bold: true}

						contentCreationDate.TextSize = 18
						contentCreationDate.Alignment = fyne.TextAlignCenter

						textAlbum.TextSize = 20
						textAlbum.Alignment = fyne.TextAlignCenter
						textAlbum.TextStyle = fyne.TextStyle{Bold: true}

						contentFirstAlbum.TextSize = 18
						contentFirstAlbum.Alignment = fyne.TextAlignCenter

						textLocation.TextSize = 20
						textLocation.Alignment = fyne.TextAlignCenter
						textLocation.TextStyle = fyne.TextStyle{Bold: true}

						contentLocation.Wrapping = fyne.TextWrapWord
						contentLocation.Alignment = fyne.TextAlignCenter

						// Organiser les canvas et containers
						memberContainer := container.New(layout.NewGridLayoutWithRows(2), textMember, contentmember)
						creationDateContainer := container.New(layout.NewGridLayoutWithRows(2), textCreationDate, contentCreationDate)
						albumContainer := container.New(layout.NewGridLayoutWithRows(2), textAlbum, contentFirstAlbum)
						locationContainer := container.New(layout.NewGridLayoutWithRows(2), textLocation, contentLocation)

						// TOP PAGE

						artistNameAndImage := container.NewVBox(contentName, contentImage)
						spacerAndButton := container.New(layout.NewGridLayoutWithRows(6), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), retour_btn)
						pageTop := container.New(layout.NewGridLayoutWithColumns(3), layout.NewSpacer(), artistNameAndImage, spacerAndButton)

						// MIDDLE PAGE

						rightMidContainer_Items := container.NewGridWithRows(4, memberContainer, creationDateContainer, albumContainer, locationContainer)

						// END MIDDLE PAGE
						// BOTTOM PAGE

						retour_btn = widget.NewButton("Effacer les résultats de la recherche", func() {
							w.SetContent(splitPage)
						})

						// END BOTTOM PAGE

						content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer())
						contentt := container.NewVScroll(content)

						// FIN ORGANISATION

						split := container.NewHSplit(
							listView,
							contentt)
						split.Offset = 0.2

						w.SetContent(split)
					})

					// Ajoutez l'image et le texte au bouton
					imageButton.Importance = widget.LowImportance
					imageButton.SetIcon(artistImageResource)
					imageButton.SetText(buttonText)
					sugg.Add(imageButton)
				}
			}
			sugg.Show()
			sugg2.Show()
		} else {
		}
	}

	searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
	sugg2.SetMinSize(fyne.NewSize(200, 200))
	searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone, sugg2)
	searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
	searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

	splitPage := container.NewHSplit(
		listView,
		searchZoneFinal)
	splitPage.Offset = 0.2
	(w).SetContent(splitPage)
	(w).SetFullScreen(false)
	(w).SetFullScreen(true)
	w.ShowAndRun()
}
