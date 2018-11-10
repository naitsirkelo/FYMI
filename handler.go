package main

import(
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
)

type Movie struct {
  Title			    string `json="Title"`
  Genre    			string `json="Genre"`
  Language 			string `json="Language"`
  Country  			string `json="Country"`
  Runtime 			string `json="Runtime"`
  Director 			string `json="Director"`
  Released 			string `json="Released"`
  Poster   			string `json="Poster"`
  Response 			string `json="Response"`
  Error    			string `json="Error"`
  ImdbRating   	string `json="imdbRating"`
  Type	   			string `json="Type"`
  TotalSeasons 	string `json="totalSeasons"`
}

// - - - - - - - - - -  Parsing Movie  - - - - - - - - - - - - - -

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
	} else if (parts[2] == "id") {
		omdbUrl = MakeUrlId(id)		//Creates the url from IMDB ID
	} else {
		fmt.Fprintln(w, "Invalid Request")
		return
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

	if movie.Response == "True" {	//Checks if omdb found the movie in imdb
		err = SendPayload(w, movie) //Send info about movie as response
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	} else {
		fmt.Fprintln(w, movie.Error)
	}
	return
}

// - - - - - - - - - -  Help function  - - - - - - - - - - - - - -

func HelpHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello! I am the FYMI-bot here to Find Your Movie Info for you! These are the available commands:")
		fmt.Fprintln(w, "\t- /fymihelp			(Show the bot functionalities)")
		fmt.Fprintln(w, "\t- /fymiid <IMDB movie id>	(Example: 'tt1790809')")
		fmt.Fprintln(w, "\t- /fymititle <IMDB movie title>	(Example: 'The Godfather')")
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
