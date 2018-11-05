package main

import(
	"fmt"
	"net/http"
)

func BotHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello there, friend.")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Internet!")
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello! I am your friendly neighbourhood Movie-Info-Finding bot!\n")
		fmt.Fprintln(w, "\tAvailable commands:\n")
		fmt.Fprintln(w, "\t- /fymiid <IMDB movie id>     			Example: 'tt1790809'\n")
		fmt.Fprintln(w, "\t- /fymititle <IMDB movie title>  	Example: 'The Godfather'\n")
}
