package main

import(
	"bytes"
	"encoding/json"
	"net/http"
)

func SendPayload(payload string, imageurl string) error {
	attachment := map[string]string{"image_url": imageurl}
	val := map[string]interface{}{"response_type": "in_channel", "text": payload, "attachments": attachment}
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
