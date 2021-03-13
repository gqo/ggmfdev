package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("main: main(): Setting up website.")
	httpPort := 8080
	portStr := ":" + strconv.Itoa(httpPort)
	fs := http.FileServer(http.Dir("assets/"))

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
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
	http.ServeFile(w, r, "./web/static/index.html")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/images/favicon.ico")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/static/about.html")
}

func musicHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./web/template/music.html"))
	Data := getAlbum()
	t.Execute(w, Data)
}
