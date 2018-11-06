package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

type Movie struct {
  Title    string `json="Title"`
  Genre    string `json="Genre"`
  Language string `json="Language"`
  Country  string `json="Country"`
  Runtime  string `json="Runtime"`
  Director string `json="Director"`
  Released string `json="Released"`
  Poster   string `json="Poster"`
}

func IdHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()	//Parse the form
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	id := r.Form["text"][0]		//Gets the id from slash command
	omdbUrl := MakeUrlId(id)	//Creates the url

	resp, err := http.Get(omdbUrl)	//Gets response from created omdb url
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	defer resp.Body.Close()

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)	//Decode json from omdb url
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Title: %v \nGenre: %v \nReleased: %v \n", movie.Title, movie.Genre, movie.Released) //Print info to slack
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello there, friend.")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Internet!")
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello! I am your friendly neighbourhood Movie-Info-Finding bot!")
		fmt.Fprintln(w, "\tAvailable commands:")
		fmt.Fprintln(w, "\t- /fymiid <IMDB movie id>     			Example: 'tt1790809'")
		fmt.Fprintln(w, "\t- /fymititle <IMDB movie title>  	Example: 'The Godfather'")
}
