package Tools

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
