package main

import(
	"fmt"
	"net/http"
//	"encoding/json"
	"github.com/gorilla/schema"
)

type Payload struct {
	Text string `schema:"text"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Internet!")
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Help wanted!")
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
