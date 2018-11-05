package main

import(
  "fmt"
  "os"
  "log"
  "net/http"
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

// - - - - - - - - - -  Port  - - - - - - - - - - - - - -

func getPort() string {
	 	var port = os.Getenv("PORT")
 				// Port sets to :8080 as a default
 		if (port == "") {
 			port = "8080"
			fmt.Println("No PORT variable detected, defaulting to " + port)
 		}
 		return (":" + port)
}

// - - - - - - - - - -  Main  - - - - - - - - - - - - -

func main() {

  http.HandleFunc("/hello", BotHandler)
  http.HandleFunc("/fymi", RootHandler)
  http.HandleFunc("/fymi/help", HelpHandler)

  err := http.ListenAndServe(getPort(), nil)
  if err != nil {
      log.Fatal("ListenAndServe Error: ", err)
  }
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - -
