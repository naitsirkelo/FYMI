package main

import (
	"testing"
)

func Test_makeUrlId(t *testing.T) {
	url := MakeUrlId("tt12345678")
	if url != "http://www.omdbapi.com/?apikey=9ef35ea6&i=tt12345678" {
		t.Errorf("Url was malformed, got: %v", url)
	}
}

func Test_makeUrlTitle(t *testing.T) {
        url := MakeUrlTitle("Star Wars: A New Hope")
        if url != "http://www.omdbapi.com/?apikey=9ef35ea6&t=Star_Wars:_A_New_Hope" {
                t.Errorf("Url was malformed, got: %v", url)
        }
}

func Test_makeUrlSearch(t *testing.T) {
        url := MakeUrlSearch("The Godfather")
        if url != "http://www.omdbapi.com/?apikey=9ef35ea6&s=The_Godfather" {
                t.Errorf("Url was malformed, got: %v", url)
        }
}

