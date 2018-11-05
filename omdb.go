package main

import (
	"strings"
)
func MakeUrlId(id string) string {
	url := OMDB + "apikey=" + APIKEY + "&i=" + id
	return url
}

func MakeUrlTitle(title string) string {
	replacer := strings.NewReplacer(" ", "_")
	title = replacer.Replace(title)
	url := OMDB + "apikey=" + APIKEY + "&t=" + title
	return url
}
