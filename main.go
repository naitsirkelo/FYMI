package main

import(
  "fmt"
  "os"
  "log"
  "net/http"
)

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

  http.HandleFunc("/fymi/help", HelpHandler)
  http.HandleFunc("/fymi/id", IdHandler)
  http.HandleFunc("/fymi/title", TestMenu)
  err := http.ListenAndServe(getPort(), nil)
  if err != nil {
      log.Fatal("ListenAndServe Error: ", err)
  }
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - -
