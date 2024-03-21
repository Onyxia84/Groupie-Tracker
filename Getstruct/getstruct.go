package getstruct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Links struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}
type Artist struct {
	Id           int      `json:"id"`
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
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type LocationsDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relation struct {
	Id        int                 `json:"id"`
	Locations map[string][]string `json:"datesLocations"`
}

func GetLinks() *Links {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	tempLink := new(Links)
	err = json.Unmarshal(responseData, tempLink)
	if err != nil {
		log.Fatal(err)
	}

	return tempLink
}
func GetArtist() *[]Artist {
	templink := GetLinks()
	response, err := http.Get(templink.Artists)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	tempArtist := new([]Artist)
	err = json.Unmarshal(responseData, &tempArtist)
	if err != nil {
		log.Fatal(err)
	}
	return tempArtist
}
