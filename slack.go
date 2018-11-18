package main

import(
	"encoding/json"
	"net/http"
)

// - - - - - - - - - -  Send info as json  - - - - - - - - - - - - - - -

func SendPayload(w http.ResponseWriter, movie Movie) error {

	var attachment [1]interface{}		//Create empty array
																	//Initialize map for attachment info
	img := map[string]string{"fallback": "Poster", "image_url": movie.Poster}
	attachment[0] = img							//Add info to attachment array
																	//Creating the slack response text
	text :=
		"Title: " + movie.Title + "\nGenre:\t" + movie.Genre + "\nReleased:\t" + movie.Released +
		"\nImdb Rating:\t" + movie.ImdbRating +	"\nType: " + movie.Type

	if movie.Type == "Series" {			//Changes response if result is a series or a movie
		text = text + "\nTotal seasons:\t" + movie.TotalSeasons
	} else if movie.Type == "Movie" {
		text = text + "\nDirector:\t" + movie.Director + "\nRuntime:\t" + movie.Runtime
	}
																	//Prepare variables as map before encoding
	val := map[string]interface{}{"text": text, "attachments": attachment}

	w.Header().Set("Content-Type", "application/json")	//Set response header
	err := json.NewEncoder(w).Encode(val)	//Encode
	if err != nil {
		return err
	}
	return nil
}

// - - - - - - - - - -  Send drop-down list  - - - - - - - - - - - - - - -

func SendMovieMenu(w http.ResponseWriter, titles []string) error{
	var attachment [1]interface{}			//Search function gets 10 the first titles.
	var actions [1]interface{}				//Initialize arrays to be able to return
	var options []interface{}					//correct JSON response to Slack

	for i:= 0; i < len(titles); i++ {
		options = append(options, map[string]string{"text": titles[i], "value": titles[i]})
	}
								//Filling maps with correct structure, to show the title
								//elements in the options array in the drop-down list.
	actions[0] = 		map[string]interface{}{"name": "titles_list", "text": "Pick a movie", "type": "select", "options": options}
	attachment[0] = map[string]interface{}{"fallback": "Menu", "actions": actions}
								//Creating a system message to be displayed in the Slack channel
	text := "Several matches for that title. Please select correct from the menu."

	val := map[string]interface{}{"text": text, "attachments": attachment}

	w.Header().Set("Content-Type", "application/json")	//Set response header
	err := json.NewEncoder(w).Encode(val)	//Encode
	if err != nil {
		return err
	}
	return nil
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
