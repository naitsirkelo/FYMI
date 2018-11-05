package main

import(
	"bytes"
	"encoding/json"
	"net/http"
)

func SendPayload(payload string) error {
	val := map[string]string{"text": payload}
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
