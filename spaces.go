package kinopigo

import (
	"encoding/json"
	"fmt"
)

type Space struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Cards       []Card       `json:"cards"`
	Connections []Connection `json:"connections"`
}

func (kc *KinopigoClient) GetSpace(id string) (Space, error) {
	var space Space

	// send request
	resp, err := kc.sendHTTPRequest("GET", fmt.Sprintf("space/%s", id), nil)
	if err != nil {
		return space, err
	}

	// attempt to decode response body into space
	err = json.NewDecoder(resp.Body).Decode(&space)
	return space, err
}
