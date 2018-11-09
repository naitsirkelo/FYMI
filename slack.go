package main

import(
	//"bytes"
	"encoding/json"
	"net/http"
)

	//Function for sending information about movie as json
func SendPayload(w http.ResponseWriter, movie Movie) error {

	var attachment [1]interface{}	//Create empty array
	img := map[string]string{"fallback": "Poster", "image_url": movie.Poster} //Map for attachment info
	attachment[0] = img	//Add info to attachment array
	text := "Title: " + movie.Title + "\nGenre: " + movie.Genre + "\nReleased: " + movie.Released +  //Create the slack response text
		"\nDirector: " + movie.Director + "\nRuntime: " + movie.Runtime + "\nImdb Rating: " + movie.Rating +
		"\nType: " + movie.Type
	if movie.Type == "series" {
		text = text + "\nTotal seasons: " + movie.TotalSeasons
	}
	val := map[string]interface{}{"text": text, "attachments": attachment}
	w.Header().Set("Content-Type", "application/json")	//Set response header
	err := json.NewEncoder(w).Encode(val)	//Encode
	if err != nil {
		return err
	}
	return nil
}
