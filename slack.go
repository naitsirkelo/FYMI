package main

import(
	"bytes"
	"fmt"
	"net/http"
)

var WebhookUrl string = "https://hooks.slack.com/services/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"

func SendPayload(payload string) error {
	fmt.Println(payload)
	var jsonStr = []byte(`{"text": payload}`)
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
