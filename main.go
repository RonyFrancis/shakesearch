package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pulley.com/shakesearch/controller"
	"pulley.com/shakesearch/models"
)

func main() {
	searcher := models.Searcher{}
	// Load completworks to searcher
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}
	// load plays from completework into searcher in structured format
	searcher.Work, err = searcher.LoadPlays()
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// search API
	http.HandleFunc("/search", controller.HandleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}


