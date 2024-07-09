package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

var films = map[string][]Film{
	"Films": {
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "Anaconda 2", Director: "Kasperskeye Coppola"},
		{Title: "Dunkirk", Director: "Christoph Coppola"},
	},
}

func main() {

	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/add-film/", addFilm)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("index.html"))

	temp.Execute(w, films)
}

func addFilm(w http.ResponseWriter, r *http.Request) {

	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	temp := template.Must(template.ParseFiles("index.html"))
	temp.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
