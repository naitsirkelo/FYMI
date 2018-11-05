package main

import(
	"fmt"
	"net/http"
//	"encoding/json"
	"goproject/schema"
)


type Payload struct {
	Text string `schema:"text"`
}

func IdHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	id := new(Payload)
	decoder := schema.NewDecoder()

	err = decoder.Decode(id, r.Form)
	if err != nil {
		return
	}
	fmt.Println(id.Text)
}
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
