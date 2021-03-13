package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	httpPort := 8080
	portStr := ":" + strconv.Itoa(httpPort)
	fs := http.FileServer(http.Dir("assets/"))

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/about", aboutHandler)

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
