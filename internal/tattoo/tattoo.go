package tattoo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

// Implements the sort.Interface
func (a Location) Len() int { return len(a.Artists) }
func (a Location) Less(i, j int) bool {
	leftName, rightName := a.Artists[i].Name, a.Artists[j].Name
	switch {
	case leftName == rightName:
		return a.Artists[i].Instagram < a.Artists[j].Instagram
	case leftName == "":
		return false
	case rightName == "":
		return true
	default:
		return leftName < rightName
	}
}
func (a Location) Swap(i, j int) {
	a.Artists[i], a.Artists[j] = a.Artists[j], a.Artists[i]
}

type Locations struct {
	Contents []Location `json:"locations"`
}

// Implements the sort.Interface
func (a Locations) Len() int { return len(a.Contents) }
func (a Locations) Less(i, j int) bool {
	return a.Contents[i].Name.EN < a.Contents[j].Name.EN
}
func (a Locations) Swap(i, j int) {
	a.Contents[i], a.Contents[j] = a.Contents[j], a.Contents[i]
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

	for i := range Data.Contents {
		sort.Sort(Data.Contents[i])
	}
	sort.Sort(Data)

	return Data
}
