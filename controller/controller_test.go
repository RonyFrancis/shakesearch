package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"pulley.com/shakesearch/models"
	"testing"
)

func TestHandleSearch(t *testing.T) {
	searcher := models.Searcher{}
	err := searcher.Load("../completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}
	searcher.Work, err = searcher.LoadPlays()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/search?q=Hamlet", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleSearch(searcher))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
