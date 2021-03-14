package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gqo/ggmfdev/internal/randalbum"
)

func main() {
	log.Println("main: main(): Setting up website.")
	httpPort := 8080
	portStr := ":" + strconv.Itoa(httpPort)

	_, err := os.Stat(filepath.Join("../../assets/css/main.css"))
	if err != nil {
		log.Println(err)
	}

	// http.Handle("/assets/", http.StripPrefix("/assets/",
	// 	http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/main.css", cssHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/music", musicHandler)

	log.Println("main: main(): Finished. Beginning execution.")
	log.Fatal(http.ListenAndServe(portStr, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.ServeFile(w, r, "../../web/static/index.html")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../assets/images/favicon.ico")
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	http.ServeFile(w, r, "../../assets/css/main.css")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/static/about.html")
}

func musicHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../../web/template/music.html"))
	Data := randalbum.GetAlbum()
	t.Execute(w, Data)
}
