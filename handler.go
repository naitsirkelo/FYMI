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
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	id := r.Form["text"][0]
	omdbUrl := MakeUrlId(id)
	fmt.Fprintln(w, id, "  Url: ", omdbUrl)

	resp, err := http.Get(omdbUrl)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	defer resp.Body.Close()

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, movie.Title)

//	fmt.Fprintf(w, "%v as %T len %v\n", r.Form["text"], r.Form["text"], len(r.Form["text"]))
//	SendPayload(payload.Text)
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
