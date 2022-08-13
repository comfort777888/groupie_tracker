package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Execute function parsing data into a web-site.
func Execute(w http.ResponseWriter, parse string, data interface{}) {
	html, err := template.ParseFiles(parse)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func CheckMethod(r *http.Request, path, method string) int {
	if r.Method != method {
		return http.StatusMethodNotAllowed
	}
	if r.URL.Path != path {
		return http.StatusNotFound
	}
	return 200
}
