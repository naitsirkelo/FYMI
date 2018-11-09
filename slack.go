package main

import(
	//"bytes"
	"encoding/json"
	"net/http"
)

func SendPayload(w http.ResponseWriter, movie Movie) error {

	var attachment [1]interface{}
	img := map[string]string{"fallback": "Poster", "image_url": movie.Poster}
	attachment[0] = img
	text := "Title: " + movie.Title + "\nGenre: " + movie.Genre + "\nReleased: " + movie.Released +
		"\nDirector: " + movie.Director + "\nRuntime: " + movie.Runtime
	val := map[string]interface{}{"text": text, "attachments": attachment}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(val)
	if err != nil {
		return err
	}
	/*jsonStr, err := json.Marshal(val)
	req, err := http.NewRequest("POST", SLACKURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()*/
	return nil
}
