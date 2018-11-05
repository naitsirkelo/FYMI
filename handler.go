package main

import(
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Internet!")
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Help wanted!")
}
