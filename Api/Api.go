package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Artist représente un groupe ou un artiste
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID           int      `json:"id"`
	Locations    []string `json:"locations"`
	ConcertDates string   `json:"dates"`
}

type Dates struct {
	ID           int      `json:"id"`
	ConcertDates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type IndexResponseloc struct {
	Indexloc []Locations `json:"index"`
}

type IndexResponsedat struct {
	Indexdat []Dates `json:"index"`
}

type IndexResponserel struct {
	Indexrel []Relations `json:"index"`
}

func ApiArtist() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Erreur lors de la requête GET: %s", err)
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		log.Fatalf("Erreur lors de la lecture du JSON: %s", err)
	}
	var data []Artist
	for _, artist := range artists {
		data = append(data, artist)
	}
	return data
}

func ApiLocations() []Locations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatalf("Erreur lors de la requête GET: %s", err)
	}
	defer resp.Body.Close()

	var response IndexResponseloc
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Erreur lors de la lecture du JSON: %s", err)
	}
	var data []Locations
	for _, location := range response.Indexloc {
		data = append(data, location)
	}
	return data
}

func ApiConcertDates() []Dates {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatalf("Erreur lors de la requête GET: %s", err)
	}
	defer resp.Body.Close()

	var response IndexResponsedat
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Erreur lors de la lecture du JSON: %s", err)
	}
	var data []Dates
	for _, dates := range response.Indexdat {
		data = append(data, dates)
	}
	return data
}

func ApiRelations() []Relations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatalf("Erreur lors de la requête GET: %s", err)
	}
	defer resp.Body.Close()

	var response IndexResponserel
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Erreur lors de la lecture du JSON: %s", err)
	}

	var data []Relations
	for _, relation := range response.Indexrel {
		data = append(data, relation)
		/*for lieu, dates := range relation.DatesLocations {
			data = append(data, lieu)
		}*/

	}
	return data
}
