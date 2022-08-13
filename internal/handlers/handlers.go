package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie/internal/logic"
)

// Home function is a main page of the site, which displays list of artist
// their picture and link for more information.
func Home(w http.ResponseWriter, r *http.Request) {
	status := CheckMethod(r, "/", http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		log.Printf("Error %d - Not allowed method\n", status)
		return
	}
	entries, err := logic.GetData()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		log.Printf("Error %d - Unable unparse data - %s\n", http.StatusInternalServerError, err)
	}
	Execute(w, "ui/templates/index.html", &entries)
}

// Artist function is responsive for artist's info.
func Artist(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	status := CheckMethod(r, "/artist/"+id, http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		log.Printf("Error %d - User tried to use not allowed method - %s \n", http.StatusMethodNotAllowed, r.Method)
		return
	}
	idint, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	entries, err := logic.GetData()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		log.Printf("Error %d - Unable unparse data - %s\n", http.StatusInternalServerError, err)
	}
	if idint < 0 || idint > len(entries.Artist) {
		ErrorHandler(w, http.StatusNotFound)
		log.Printf("Error %d - Page (http://localhost:8080%s) not found", http.StatusNotFound, r.URL.Path)
		return
	}
	Execute(w, "ui/templates/artist.html", &entries.Artist[idint-1])
}

// Relation function is responsive for displaying information about location and date artist had.
func Relation(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/relations/")
	status := CheckMethod(r, "/relations/"+id, http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		log.Printf("Error %d - User tried to use not allowed method - %s \n", http.StatusMethodNotAllowed, r.Method)
		return
	}
	idint, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		log.Printf("Error %d - User tried to use not allowed method - %s \n", http.StatusNotFound, r.Method)
		return
	}
	entries, err := logic.GetData()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		log.Printf("Error %d - Unable unparse data - %s\n", http.StatusInternalServerError, err)
		return
	}
	if idint < 0 || idint > len(entries.Artist) {
		ErrorHandler(w, http.StatusBadRequest)
		log.Printf("Error %d - Page (http://localhost:8080%s) not found", http.StatusNotFound, r.URL.Path)
		return
	}
	Execute(w, "ui/templates/concert.html", entries.Artist[idint-1])
}

// function Searchbar is initialized only if
func SearchBar(w http.ResponseWriter, r *http.Request) {
	status := CheckMethod(r, "/search", http.MethodGet)
	if status != 200 {
		ErrorHandler(w, status)
		return
	}
	search := r.FormValue("send")
	entries, err := logic.SearchBar(search)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	Execute(w, "ui/templates/index.html", &entries)
}
