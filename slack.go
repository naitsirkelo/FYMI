package main

import(
	"bytes"
	"encoding/json"
	"net/http"
)

var WebhookUrl string = "https://hooks.slack.com/services/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"
var OMDbkey 	 string = "9ef35ea6"

func SendPayload(payload string) error {
	val := map[string]string{"text": payload}
	jsonStr, err := json.Marshal(val)
	req, err := http.NewRequest("POST", WebhookUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return nil
}
