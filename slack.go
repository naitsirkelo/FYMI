package main

import(
	"encoding/json"
	"net/http"
)

// - - - - - - - - - -  Send info as json  - - - - - - - - - - - - - -

func SendPayload(w http.ResponseWriter, movie Movie) error {

	var attachment [1]interface{}		//Create empty array
																	//Initialize map for attachment info
	img := map[string]string{"fallback": "Poster", "image_url": movie.Poster}
	attachment[0] = img							//Add info to attachment array
																	//Creating the slack response text
	text :=
		"Title: " + movie.Title + "\nGenre: " + movie.Genre + "\nReleased: " + movie.Released +
		"\nDirector: " + movie.Director + "\nImdb Rating: " + movie.ImdbRating +
		"\nType: " + movie.Type

	if movie.Type == "series" {
		text = text + "\nTotal seasons: " + movie.TotalSeasons
	} else if movie.Type == "movie" {
		text = text + "\nRuntime: " + movie.Runtime
	}
	val := map[string]interface{}{"text": text, "attachments": attachment}

	w.Header().Set("Content-Type", "application/json")	//Set response header
	err := json.NewEncoder(w).Encode(val)	//Encode
	if err != nil {
		return err
	}
	return nil
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
