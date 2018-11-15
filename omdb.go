package main

import (
	"strings"
)

// - - - - - - - - - -  Url with IMDb ID  - - - - - - - - - - - - - - - -

func MakeUrlId(id string) string {
	replacer := strings.NewReplacer(" ", "")	//Necessary to remove spaces for correct error handling
	id = replacer.Replace(id)
	url := OMDB + "apikey=" + APIKEY + "&i=" + id //Builds the url
	return url
}

// - - - - - - - - - -  Url with IMDb title  - - - - - - - - - - - - - -

func MakeUrlTitle(title string) string {
	replacer := strings.NewReplacer(" ", "_") //Replaces spaces with underscore in title
	title = replacer.Replace(title)
	url := OMDB + "apikey=" + APIKEY + "&t=" + title //Builds the url
	return url
}

func MakeUrlSearch(title string) string {
	replacer := strings.NewReplacer(" ", "_")
	title = replacer.Replace(title)
        url := OMDB + "apikey=" + APIKEY + "&s=" + title
	return url
}
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
