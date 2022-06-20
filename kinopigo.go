package kinopigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type KinopigoClient struct {
	Token   string
	BaseURL string
}

func NewKinopigoClient() (KinopigoClient, error) {
	var client KinopigoClient

	// attempt to load API token from env var
	token := os.Getenv("KINOPIO_API_KEY")
	if token == "" {
		return client, errors.New("Environment variable KINOPIO_API_KEY is unset")
	}

	// create and return a new instance of the client struct
	client = KinopigoClient{
		Token:   token,
		BaseURL: "https://api.kinopio.club",
	}
	return client, nil
}

// utility function for sending HTTP requests to the API
func (kc *KinopigoClient) sendHTTPRequest(method string, path string, body interface{}) (*http.Response, error) {
	// encode body
	var buffer io.ReadWriter = nil
	if body != nil {
		buffer = new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(body)
	}

	// create request
	req, _ := http.NewRequest(method, fmt.Sprintf("%s/%s", kc.BaseURL, path), buffer)
	req.Header.Set("Authorization", kc.Token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// dispatch request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	// check status code
	expectedStatusCode := http.StatusOK
	if method == "POST" {
		expectedStatusCode = http.StatusCreated
	}
	if resp.StatusCode != expectedStatusCode {
		return resp, fmt.Errorf("Non-success status code returned: %d", resp.StatusCode)
	}

	return resp, nil
}
