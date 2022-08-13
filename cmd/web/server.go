package web

import (
	"log"
	"net/http"
	"time"

	"groupie/internal/handlers"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	mux := http.NewServeMux()
	a.httpServer = &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artist/", handlers.Artist)
	mux.HandleFunc("/relations/", handlers.Relation)
	mux.HandleFunc("/search", handlers.SearchBar)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("ui/assets"))))
	log.Println("Server is on: http://localhost:8080/")
	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
