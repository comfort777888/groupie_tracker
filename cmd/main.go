package main

import (
	"log"

	"groupie/cmd/web"
)

func main() {
	app := web.NewApp()
	if err := app.Run(); err != nil {
		log.Fatalf("Server run error: %v\n", err)
	}
}
