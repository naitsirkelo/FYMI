package main

import(
	"fmt"
	"net/http"
//	"encoding/json"
	"github.com/gorilla/schema"
)


type Payload struct {
	Text string `schema:"&text"`
}

func IdHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	var payload Payload
	decoder := schema.NewDecoder()

	err = decoder.Decode(&payload, r.PostForm)
	if err != nil {
		return
	}
	fmt.Fprintln(w, payload.Text)
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
