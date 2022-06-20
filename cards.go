package kinopigo

import (
	"encoding/json"
)

type Card struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	SpaceID         string `json:"spaceId"`
	ParentID        string `json:"parentId"`
	BackgroundColor string `json:"backgroundColor"`
	X               int    `json:"x"`
	Y               int    `json:"y"`
	Z               int    `json:"z"`
}

func (kc *KinopigoClient) CreateCard(card Card) (Card, error) {
	// send request
	resp, err := kc.sendHTTPRequest("POST", "card", card)
	if err != nil {
		return card, err
	}

	// attempt to decode response into card and return
	err = json.NewDecoder(resp.Body).Decode(&card)
	return card, err
}
