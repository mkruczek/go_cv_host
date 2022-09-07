package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("starting server....\n")
	http.Handle("/", http.FileServer(http.Dir("./html")))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("server down: %s", err)
	}
}
