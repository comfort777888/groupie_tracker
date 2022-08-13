package handlers

import (
	"net/http"
)

type Error struct {
	Message string
	Status  int
}

func ErrorHandler(w http.ResponseWriter, status int) {
	errHandler := setError(status)
	w.WriteHeader(status)
	Execute(w, "ui/templates/error.html", errHandler)
}

func setError(status int) *Error {
	return &Error{
		Status:  status,
		Message: http.StatusText(status),
	}
}
