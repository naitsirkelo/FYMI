package main

import(
	"fmt"
	"strings"
	"net/http"
	"encoding/json"

	// "github.com/victorgama/go-unfurl"
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
	parts := strings.Split(r.URL.Path, "/")

	err := r.ParseForm()	//Parse the form
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	id := r.Form["text"][0]			//Gets the movie variable from slack

	var omdbUrl string
	if (parts[2] == "title") {
		omdbUrl = MakeUrlTitle(id) //Creates the url from the movie title
	} else {
		omdbUrl = MakeUrlId(id)		//Creates the url from IMDB ID
	}

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
	SendPayload(w, movie)
	//fmt.Fprintf(w, "Title: %v \nGenre: %v \nReleased: %v \nDirector: %v \nRuntime: %v \nPoster: %v",
	//						movie.Title, movie.Genre, movie.Released, movie.Director, movie.Runtime, movie.Poster)
}


/*
func PosterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()	//Parse the form
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	title := r.Form["text"][0]			//Gets the id from slash command
	omdbUrl := MakeUrlTitle(title) 	//Creates the url from the movie title

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

//	SendPayload(movie.Title, movie.Poster)
}

*/
func BotHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello there, friend.")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Internet!")
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello! I am the FYMI-bot here to Find Your Movie Info for you! These are the available commands:")
		fmt.Fprintln(w, "\t- /fymihelp			(Show the bot functionalities)")
		fmt.Fprintln(w, "\t- /fymiid <IMDB movie id>	(Example: 'tt1790809')")
		fmt.Fprintln(w, "\t- /fymititle <IMDB movie title>	(Example: 'The Godfather')")
}
