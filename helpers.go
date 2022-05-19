package helly

import (
	"encoding/json"
	"net/http"
)

const additionalOptions = "?v=10&encoding=json"

func getDiscordGatewayURL() (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://discord.com/api/v10/gateway", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "DiscordBot (https://github.com/denkylabs/disgo, 0.0.1)")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	// Parse the response object into a string
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return "", err
	}

	return response["url"] + additionalOptions, nil
}
