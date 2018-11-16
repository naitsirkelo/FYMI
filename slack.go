package main

import(
	"encoding/json"
	"net/http"
//	"fmt"
//	"str"
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

func SendMovieMenu(w http.ResponseWriter, titles []string) error{
	var attachment [1]interface{}
	var actions [1]interface{}
	var options [1]interface{}

/*	for i:= 0; i < len(titles); i++ {
		options[i] = map[string]string{"text": titles[i], "value": titles[i]}
	}
*/

	options[0] = map[string]string{"text": "testText", "value": "testValue"}

	actions[0] = map[string]interface{}{"name": "titles_list", "text": "Pick a movie", "type": "select", "options": options}

	attachment[0] =  map[string]interface{}{"fallback": "Menu", "actions": actions}
	text := "Several matches for that title. Please select correct from the menu."

	val := map[string]interface{}{"text": text, "attachments": attachment}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(val)	//Encode
	if err != nil {
		return err
	}
//	fmt.Println(str(json.Marshal(val)))
	return nil
}
