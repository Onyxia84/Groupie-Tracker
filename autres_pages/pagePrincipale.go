package autres_pages

import (
	"groupietracker/Artists"
	"groupietracker/Tools"
	"groupietracker/getstruct"
	"image/color"
	"net/http"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func PagePrincipale(w *fyne.Window) {

	// Déclaration des variables

	var artist []getstruct.Artist
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

	/*searchResults := container.NewVBox()

	// Fonction pour mettre à jour les suggestions de recherche à chaque saisie dans le champ de recherche
	e_recherche.OnChanged = func(text string) {
		Tools.GenerateSearchSuggestions(text, searchResults, artist)
	}*/
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
				v := *w
				v.SetContent(splitPage)
				w = &v
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
				v := *w
				v.SetContent(splitPage)
				w = &v
			})

			contentText := widget.NewLabel("Recherchez quelque chose içi : ")
			searchZone = container.NewAdaptiveGrid(2, e_recherche, retour_btn)
			searchZoneStretch = container.New(layout.NewVBoxLayout(), contentText, searchZone)
			searchZoneFinished := container.New(layout.NewVBoxLayout(), searchZoneStretch)

			// END BOTTOM PAGE

			content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer(), searchZoneFinished)
			contentt := container.NewVScroll(content)

			// FIN ORGANISATION

			split := container.NewHSplit(
				listView,
				contentt)
			split.Offset = 0.9

			v := *w
			v.SetContent(split)
			w = &v
		}

		// Réinitialiser la page de droite et effacer les résultats de la recherche

		retour_btn = widget.NewButton("Effacer les résultats de la recherche", func() {

			searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
			searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone)
			searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
			searchZoneFinal := container.NewGridWithRows(8, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

			splitPage := container.NewHSplit(
				listView,
				searchZoneFinal)
			splitPage.Offset = 0.2

			// Affichage

			v := *w
			v.SetContent(splitPage)
			w = &v
			listR = listVierge
		})

		// Disposition de la page de droite

		splitPage := container.NewHSplit(
			listView,
			container.NewVBox(e_recherche, retour_btn, listViewR))
		splitPage.Offset = 0.2

		// Affichage

		v := *w
		v.SetContent(splitPage)
		w = &v
	})

	// Disposition de la page lorsque la listView (gauche) est activée

	listView.OnSelected = func(id widget.ListItemID) {

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

			// Affichage

			v := *w
			v.SetContent(splitPage)
			w = &v
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

		contentText := widget.NewLabel("Recherchez quelque chose içi : ")
		searchZone = container.NewAdaptiveGrid(2, e_recherche, subt_btn)
		searchZoneStretch = container.New(layout.NewVBoxLayout(), contentText, searchZone)
		searchZoneFinished := container.New(layout.NewVBoxLayout(), searchZoneStretch)

		// END BOTTOM PAGE

		content := container.NewVBox(pageTop, rightMidContainer_Items, layout.NewSpacer(), searchZoneFinished)
		contentt := container.NewVScroll(content)

		// FIN ORGANISATION

		split := container.NewHSplit(
			listView,
			contentt)
		split.Offset = 0.2

		// Affichage

		v := *w
		v.SetContent(split)
		w = &v
	}

	// Disposition de la page par défaut

	searchZone := container.NewAdaptiveGrid(2, e_recherche, subt_btn)
	searchZoneStretch := container.New(layout.NewVBoxLayout(), searchText, searchZone)
	searchZoneStretchCentered := container.NewGridWithColumns(3, layout.NewSpacer(), searchZoneStretch, layout.NewSpacer())
	searchZoneFinal := container.NewGridWithRows(5, layout.NewSpacer(), layout.NewSpacer(), searchZoneStretchCentered, layout.NewSpacer(), layout.NewSpacer())

	splitPage := container.NewHSplit(
		listView,
		searchZoneFinal)
	splitPage.Offset = 0.2

	// Affichage

	v := *w
	v.SetContent(splitPage)
	w = &v
}
