package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gqo/ggmfdev/internal/langsupport"
	"github.com/gqo/ggmfdev/internal/randalbum"
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

	serveStaticMultiLang(w, r, "index")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../assets/images/favicon.ico")
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	http.ServeFile(w, r, "../../assets/css/main.css")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "../../web/static/about.html")
	serveStaticMultiLang(w, r, "about")
}

func musicHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/template/music.html"))
	Data := randalbum.GetAlbum()
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

func serveStaticMultiLang(w http.ResponseWriter, r *http.Request, page string) {
	lang := langsupport.DetermineLanguage(r)

	filepath := "../../web/static/"
	switch lang {
	case language.AmericanEnglish:
		filepath += "en/"
	case language.Japanese:
		filepath += "jp/"
	}
	filepath += page + ".html"

	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, r, filepath)
}
