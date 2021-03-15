package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gqo/ggmfdev/internal/langsupport"
	"github.com/gqo/ggmfdev/internal/randalbum"
	"github.com/gqo/ggmfdev/internal/tattoo"
	"golang.org/x/text/language"
)

func main() {
	log.Println("main: main(): Setting up website.")
	httpPort := 8080
	portStr := ":" + strconv.Itoa(httpPort)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/main.css", cssHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/music", musicHandler)
	http.HandleFunc("/tattoo", tattooHandler)
	http.HandleFunc("/eng", engHandler)
	http.HandleFunc("/jp", jpHandler)

	log.Println("main: main(): Finished. Beginning execution.")
	log.Fatal(http.ListenAndServe(portStr, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	filepath := getPagePath(r, "index", "static")

	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, r, filepath)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../assets/images/favicon.ico")
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	http.ServeFile(w, r, "../../assets/css/main.css")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	filepath := getPagePath(r, "about", "static")

	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, r, filepath)
}

func musicHandler(w http.ResponseWriter, r *http.Request) {
	filepath := getPagePath(r, "music", "template")

	t := template.Must(template.ParseFiles(filepath))
	Data := randalbum.GetAlbum()

	t.Execute(w, Data)
}

func tattooHandler(w http.ResponseWriter, r *http.Request) {
	filepath := getPagePath(r, "tattoo", "template")

	t := template.Must(template.ParseFiles(filepath))
	Data := tattoo.GetTattooArtists()

	t.Execute(w, Data)
}

func engHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "lang",
		Value: "en",
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func jpHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "lang",
		Value: "ja",
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func getPagePath(r *http.Request, pageName, pageType string) string {
	lang := langsupport.DetermineLanguage(r)

	filepath := "../../web/" + pageType + "/"
	switch lang {
	case language.AmericanEnglish:
		filepath += "en/"
	case language.Japanese:
		filepath += "jp/"
	}
	filepath += pageName + ".html"

	return filepath
}
