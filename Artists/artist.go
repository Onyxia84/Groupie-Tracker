package Artists

import (
	"encoding/json"
	"fmt"
	"groupietracker/getstruct"
	"log"
	"net/http"
	"os"
)

func GetArtist() []getstruct.Artist {
	templink := getstruct.GetLinks()
	var artists []getstruct.Artist
	response, err := http.Get(templink.Artists)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// responseData, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
	}
	return artists
}

func GetLocations(url []getstruct.Artist, id int) getstruct.Locations {
	var locations getstruct.Locations
	response, err := http.Get(url[id].Locations)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// responseData, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	err = json.NewDecoder(response.Body).Decode(&locations)
	if err != nil {
		log.Fatal(err)
	}
	return locations
}
func GetLocationsDate(url getstruct.Locations) getstruct.LocationsDate { // du coup ici il va chercher les dates pour les locations du coup il faut l'information sur qu'elle artiste puis lieux
	var locationsDate getstruct.LocationsDate
	response, err := http.Get(url.Dates)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// responseData, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	err = json.NewDecoder(response.Body).Decode(&locationsDate)
	if err != nil {
		log.Fatal(err)
	}
	return locationsDate
}
func GetConcertDates(url []getstruct.Artist, Id int) getstruct.ConcertDates {
	var ConcertDates getstruct.ConcertDates
	response, err := http.Get(url[Id].ConcertDates)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// responseData, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	err = json.NewDecoder(response.Body).Decode(&ConcertDates)
	if err != nil {
		log.Fatal(err)
	}
	return ConcertDates
}

func GetLocationsRelation(url []getstruct.Artist, id int) getstruct.Relation {
	var Relation getstruct.Relation
	response, err := http.Get(url[id].Relations)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// responseData, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	err = json.NewDecoder(response.Body).Decode(&Relation)
	if err != nil {
		log.Fatal(err)
	}
	return Relation
}
