package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"strings"
)

var ContentNotFound = errors.New("contents not found")

type Searcher struct {
	// CompleteWorks contains the entire data of the book
	CompleteWorks string
	SuffixArray   *suffixarray.Index
	// Work structured format of the completeWorks
	// CompleteWork is parsed into a more readable format
	Work          Work
	// Contents list of plays for faster access in searching
	Contents      []string

}

// Work list of play from  William Shakespeare
type Work struct {
	Plays []Play `json:"plays"`
}

// Play contains details related to the play like name and content(script)
type Play struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

// Load the file into searcher struct
func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	return nil
}

//
func (s * Searcher) LoadPlays() (work Work, err error) {
	// fetch all play list
  	contents, err := s.loadContents()
	if err != nil {
		return work, err
	}
	s.Contents = contents
	var content string
	var contentList []string
	var play Play
	for i, val := range contents {
		play.Name = strings.TrimSpace(val)
		if i < len(contents) - 1 {
			content = s.loadPlay(strings.TrimSpace(val), strings.TrimSpace(contents[i + 1]))
		} else {
			content = s.loadPlay(strings.TrimSpace(val),"")
		}
		contentList = s.DeleteEmpty(strings.Split(content, "\r\n"))
		buff, err := json.Marshal(contentList)
		if err != nil {
			fmt.Println(err.Error())
		}
		play.Content = buff
		work.Plays = append(work.Plays, play)
	}
	return work, err
}

func  (s *Searcher) loadContents() ([]string, error) {
	idxs := s.SuffixArray.Lookup([]byte("Contents"), 1)
	results := []string{}
	for _, idx := range idxs {
		results = append(results, s.CompleteWorks[idx + 8 :2920])
	}
	if len(results) == 0 {
		return results, ContentNotFound
	}
	results = strings.Split(strings.TrimSpace(results[0]), "\r\n\r\n")
	return results, nil
}

func (s *Searcher) loadPlay(currentPlay, nextPlay string) string {
	var currentPlayIdxs, nextPlayIdxs []int
	currentPlayIdxs = s.fetchStartingIndex(currentPlay, false, 0)
	if nextPlay != "" {
		nextPlayIdxs = s.fetchStartingIndex(nextPlay, true, currentPlayIdxs[0])
	} else {
		nextPlayIdxs = s.SuffixArray.Lookup([]byte("FINIS"), 1)
	}
	results := append(currentPlayIdxs, nextPlayIdxs...)
	if len(results) < 2 {
		return ""
	}
	// return everything from heading of the current play and nextplay
	return s.CompleteWorks[results[0]:results[1]]
}

func (s *Searcher) fetchStartingIndex(play string, isNextPlay bool, currentPlayIdx int) (results []int){
	idxs := s.SuffixArray.Lookup([]byte(play), 2)
	if len(idxs) == 1 {
		substrings := strings.Split(play, " ")
		play = substrings[len(substrings) - 1 ]
		idxs = s.SuffixArray.Lookup([]byte(play), 2)
	}
	for _, idx := range idxs {
		if idx < 2920 {
			continue
		}
		if isNextPlay &&   idx < currentPlayIdx {
			continue
		}
		results = append(results, idx)
	}
	return
}

func (searcher Searcher)DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}




