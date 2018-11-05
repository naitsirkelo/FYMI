package main

import(
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Internet!")
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Help wanted!")
}
