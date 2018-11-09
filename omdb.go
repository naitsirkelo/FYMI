package main

import (
	"strings"
)

func MakeUrlId(id string) string {
	replacer := strings.NewReplacer(" ", "")	//Necessary to remove spaces for correct error handling
	id = replacer.Replace(id)
	url := OMDB + "apikey=" + APIKEY + "&i=" + id //Builds the url
	return url
}

func MakeUrlTitle(title string) string {
	replacer := strings.NewReplacer(" ", "_") //Replaces spaces with underscore in title
	title = replacer.Replace(title)
	url := OMDB + "apikey=" + APIKEY + "&t=" + title //Builds the url
	return url
}
