package search

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    // "fyne.io/fyne/v2/canvas"
    // "fyne.io/fyne/v2/container"
    // "fyne.io/fyne/v2/widget"
    // "fyne.io/fyne/v2"
    // "fyne.io/fyne/v2/widget/x"

)

type GeoLocation struct {
    Lat string `json:"lat"`
    Lon string `json:"lon"`
}

func ConvertAddressToGeoCoord(address string) (GeoLocation, error) {
    var geoLocation GeoLocation
    baseURL := "https://nominatim.openstreetmap.org/search"
    query := fmt.Sprintf("?q=%s&format=json&limit=1", url.QueryEscape(address))

    // Set custom User-Agent
    client := &http.Client{}
    req, err := http.NewRequest("GET", baseURL+query, nil)
    if err != nil {
        return geoLocation, err
    }
    req.Header.Set("User-Agent", "GroupieTrackerApp")

    resp, err := client.Do(req)
    if err != nil {
        return geoLocation, err
    }
    defer resp.Body.Close()

    var result []GeoLocation
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return geoLocation, err
    }

    if len(result) > 0 {
        geoLocation = result[0]
    } else {
        return geoLocation, fmt.Errorf("no results found for address: %s", address)
    }

    return geoLocation, nil
}

// func CreateMapWithGeoLocation(latitude, longitude float64) fyne.CanvasObject {

//     mapObject := canvas.NewMapWithMarker("latitude", "longitude") // prendre les coordonnées de l'api

//     mapObject.SetZoomLevel(15)

//     mapContainer := container.NewBorder(
//         widget.NewLabel("Map"),
//         nil,
//         nil,
//         mapObject,
//     )

//     return mapContainer
// }

func GetLocationsData(ID int) GeoLocation { // a utiliser
    var geoLocation GeoLocation
    url := "https://groupietrackers.herokuapp.com/api/locations"
    data, err := http.Get(url + "/" + strconv.Itoa(ID))
    if err != nil {
        print(err)
    }
    defer data.Body.Close()
    err = json.NewDecoder(data.Body).Decode(&geoLocation)
    if err != nil {
        print(err)
    }
    return geoLocation
}