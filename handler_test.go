package main

import (
	"testing"
	"net/http"
	"strings"
)


func Test_idHandler(t *testing.T) {

	//Payload containing all fields sent as a POST request when using
	//a slash command in Slack to this API, to simulate the APIs response.
	//In this case "Star Wars" is passed as text parametre from Slack.
	payload := strings.NewReader("token=f&%26team_id=f&%26team_domain=d&%26enterprise_id=&%26enterprise_name=&%26channel_id=&%26channel_name=&%26user_id=&%26user_name=&%26command=&%26text=Star%20Wars&%26response_url=&%26trigger_id=")

	req, err := http.NewRequest("POST", "/fymi/title", payload)
	if err != nil {
		t.Errorf("Error making the POST request, %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
