package tools

import (
	"strings"
)

func ToLower(s string) string {
	var x []rune
	x = []rune(s)
	if s == "" {
		return ""
	}
	h := 0
	for h <= len(s)-1 {
		if rune(s[h]) >= 'A' && rune(s[h]) <= 'Z' {
			x[h] += 32
		}
		h++
	}
	return (string(x))
}

func stripStars(getConcertDates []string) []string {
	strippedDates := make([]string, len(getConcertDates))
	for i, date := range getConcertDates {
		if strings.HasPrefix(date, "*") {
			strippedDates[i] = strings.TrimPrefix(date, "*")
		} else {
			strippedDates[i] = date
		}
	}
	return strippedDates
}

func StringAppend(Tab []string) string {
	var Newstring string
	tabString := stripStars(Tab)
	for i := 0; i < len(Tab); i++ {
		// Newstring = Newstring + tabString[i] + ", "
		if i == len(Tab)-1 {
			Newstring = Newstring + tabString[i] + " "
		} else {
			Newstring = Newstring + tabString[i] + ", "
		}
	}
	return Newstring
}

func MapString(ma map[string][]string) []string {
	var tabString []string
	for k, v := range ma {
		tabString = append(tabString, k+" ", StringAppend(v)+" ")
	}
	return tabString
}

func Recherche(comparer string, recherche string) bool {
	maxC := len(comparer)
	maxR := len(recherche)
	recherche = ToLower(recherche)
	comparer = ToLower(comparer)
	dif := len(comparer) - len(recherche)
	if dif < 0 {
		return false
	}
	if maxR == 0 {
		return false
	}
	if maxC < maxR {
		return false
	}
	if comparer == recherche {
		return true
	} else {
		for i := 0; i < dif; i++ {
			// fmt.Println(comparer[i:dif])
			if comparer[i:maxR+i] == recherche[0:maxR] {
				return true
			}
		}
	}
	return false
}

/*func GenerateSearchSuggestions(text string, searchResults *fyne.Container, artists []getstruct.Artist) {
	searchResults.Objects = nil

	if text == "" {
		return
	}

	var found bool
	var correspondingResultAdded bool

	for id, artist := range artists {
		if strings.Contains(strings.ToLower(Artists.GetArtist()[id].Name), strings.ToLower(text)) {
			found = true

			if !correspondingResultAdded {
				correspondingResultLabel := widget.NewLabel("Corresponding result: ")
				searchResults.Add(correspondingResultLabel)

				correspondingResultAdded = true
			}

			artistButton := widget.NewButton(Artists.GetArtist()[id].Name, func() {
				artist = artist
			})

			searchResults.Add(layout.NewSpacer())

			searchResults.Add(artistButton)
		}
	}

	if !found {
		noResultLabel := widget.NewLabel("No result")
		searchResults.Add(noResultLabel)
	}
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}*/


// // filtre par nombre de membre croissant au décroissant
// croissant := widget.NewButton("ordre décroissant", func() {
// 	sort.Slice(Artist, func(i, j int) bool {
// 		if len(Artist[i].Members) != len(Artist[j].Members) {
// 			return len(Artist[i].Members) > len(Artist[j].Members)
// 		}
// 		return Artist[i].Name < Artist[j].Name
// 	})
// 	updateArtistGrille(grille, Artist, w)
// })
// décroissant := widget.NewButton("Ordre croissant", func() {
// 	sort.Slice(Artist, func(i, j int) bool {
// 		if len(Artist[i].Members) != len(Artist[j].Members) {
// 			return len(Artist[j].Members) > len(Artist[i].Members)
// 		}
// 		return Artist[j].Name < Artist[i].Name
// 	})
// 	updateArtistGrille(grille, Artist, w)
// })
// // création des filtres par date de création
// searchEntry := widget.NewEntry()
// creationDateEntry := widget.NewEntry()
// creationDateEntry.SetPlaceHolder("Date")
// creationDateButton := widget.NewButton("Rechercher", func() {
// 	creationDate, err := strconv.Atoi(creationDateEntry.Text)
// 	if err != nil {
// 		creationDate = 0
// 	}
// 	filteredArtists := filtreArtists(Artist, searchEntry.Text, creationDate, true != false)
// 	updateArtistGrille(grille, filteredArtists, w)
// })
// // Créer un bouton pour chaque nombre de membres unique
// var memberCountButtons []*widget.Button
// for memberCount := 1; memberCount <= 8; memberCount++ {
// 	if uniqueMemberCounts[memberCount] {
// 		memberCountButton := widget.NewButton(fmt.Sprintf("%d membres", memberCount), func(count int) func() {
// 			return func() {
// 				// Filtre les artistes en fonction du nombre de membres
// 				filteredArtists := filtreArtists(Artist, "", 0, false)
// 				var nonEmptyArtists []Artists
// 				for _, artist := range filteredArtists {
// 					if len(artist.Members) == count {
// 						nonEmptyArtists = append(nonEmptyArtists, artist)
// 					}
// 				}
// 				// mise à jour de la grille pour les artistes avec les artistes filtrés
// 				updateArtistGrille(grille, nonEmptyArtists, w)
// 			}
// 		}(memberCount))
// 		memberCountButtons = append(memberCountButtons, memberCountButton)
// 	}
// }
