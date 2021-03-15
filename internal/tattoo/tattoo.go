package tattoo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type MultiText struct {
	EN string `json:"en"`
	JP string `json:"jp"`
}

type Artist struct {
	Name         string      `json:"name"`
	Location     string      `json:"location"`
	LocationLink string      `json:"location_link"`
	Instagram    string      `json:"instagram"`
	IsTattooed   bool        `json:"is_tattooed"`
	Notes        []MultiText `json:"notes"`
}

type Location struct {
	Name    MultiText `json:"name"`
	Artists []Artist  `json:"artists"`
}

type Locations struct {
	Contents []Location `json:"locations"`
}

func GetTattooArtists() Locations {
	file, err := os.Open("../../assets/docs/tattoo_artists.json")
	if err != nil {
		log.Println("tattoo: GetTattooArtists(): Unable to open tattoo_artists.json")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("tattoo: GetTattooArtists(): Unable to read bytes from JSON")
	}

	var Data Locations

	json.Unmarshal(bytes, &Data)

	return Data
}
