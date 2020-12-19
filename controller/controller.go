package controller

import (
	"net/http"
	"pulley.com/shakesearch/models"
	"strings"
)

func HandleSearch(searcher models.Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch query from the url params
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}

		// searching the request query is present in the contents List
		var index int
		index = -1
		for i, play := range searcher.Contents {
			if index != -1 { break }
			// For case insensitive search and substring search
			if strings.Contains(strings.ToLower(play), strings.ToLower(query[0])) {
				index = i
			}
		}
		if index == -1  {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("couldn't find the requested play"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(searcher.Work.Plays[index].Content)
	}
}