package main

import(
  "fmt"
  "os"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Internet!")
}

// - - - - - - - - - -  Port  - - - - - - - - - - - - - -

func GetPort() string {
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
  fmt.Println("Hello World!")

  http.HandleFunc("/", handler)

  err := http.ListenAndServe(GetPort(), nil)
  if err != nil {
      log.Fatal("ListenAndServe Error: ", err)
  }
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - -
