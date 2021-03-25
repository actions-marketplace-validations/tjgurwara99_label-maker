package github

import (
	"encoding/json"
	"fmt"
)

// Payload struct for storing the relavant
// information from the events payload from the json file
type Payload struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

// Issue struct for storing the relavant Issue
// information from the events payload
type Issue struct {
	RepositoryURL string `json:"repository_url"`
	URL           string `json:"url"`
	Title         string `json:"title"`
	// Body  string
}

// GetEventInfo Returns Event  using the json event payload
func GetPayloadInfo(payloadBytes []byte) (*Payload, error) {
	var payload Payload

	err := json.Unmarshal(payloadBytes, &payload)

	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling the []byte of payload to Payload: %v", err)
	}

	return &payload, nil
}
