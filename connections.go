package main

import "encoding/json"

type Connection struct {
	SpaceID          string `json:"spaceId"`
	ConnectionTypeID string `json:"connectionTypeId"`
	StartCardID      string `json:"startCardId"`
	EndCardID        string `json:"endCardId"`
}

func (kc *KinopigoClient) CreateConnection(connection Connection) (Connection, error) {
	// send request
	resp, err := kc.sendHTTPRequest("POST", "connection", connection)
	if err != nil {
		return connection, err
	}

	// attempt to decode response into connection and return
	err = json.NewDecoder(resp.Body).Decode(&connection)
	return connection, err
}
