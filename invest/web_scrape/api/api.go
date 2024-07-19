package api

import (
	"fmt"
	"io"
	"net/http"
)

// config.ConfigInfo.GoldConfig.API.Url
// config.ConfigInfo.GoldConfig.API.ApiKey

func CallApi(url string, key string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error making request\n%w", err)
	}

	// Add headers to the request
	req.Header.Add("x-access-token", key)

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request\n%w", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body\n%w", err)
	}

	// Print the response body
	return string(body), nil
}
