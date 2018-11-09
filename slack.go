package main

import(
	"bytes"
	"encoding/json"
	"net/http"
)

func SendPayload(payload string, imageurl string) error {

	var attachment [1]interface{}
	img := map[string]string{"image_url": imageurl}
	attachment[0] = img
	val := map[string]interface{}{"text": payload, "attachments": attachment}

	jsonStr, err := json.Marshal(val)
	req, err := http.NewRequest("POST", SLACKURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return nil
}
