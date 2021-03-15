package loadtext

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/language"
)

type PageType string

const (
	Index PageType = "Index"
	About PageType = "About"
	Music PageType = "Music"
)

type FileText struct {
	ENText []string `json:"en"`
	JPText []string `json:"jp"`
}

type LangText struct {
	Text []string
}

func GetPageText(page PageType, lang language.Tag) LangText {
	filepath := "../../assets/text/"
	switch page {
	case Index:
		filepath += "index"
	case About:
		filepath += "about"
	case Music:
		filepath += "music"
	}
	filepath += ".json"

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("loadtext: GetPageText(): Unable to open file at:", filepath)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("loadtext: GetPageText(): Unable to read bytes from JSON.")
	}

	var fileText FileText

	json.Unmarshal(bytes, &fileText)

	var Data LangText

	switch lang {
	case language.AmericanEnglish:
		Data.Text = fileText.ENText
	case language.Japanese:
		Data.Text = fileText.JPText
	}

	return Data
}
