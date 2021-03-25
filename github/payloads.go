package github

import (
	"encoding/json"
	"fmt"
)

// Payload struct for storing the relevant
// information from the events payload from the json file
type Payload struct {
	Action      string      `json:"action"`
	Issue       Issue       `json:"issue"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

// Issue struct for storing the relevant Issue
// information from the events payload
type Issue struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PullRequest struct for storing the relevant Pull Request
// information from the events payload
type PullRequest struct {
	IssueURL string `json:"issue_url"`
	URL      string `json:"url"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

// Repository struct for storing the relevant Repository
// information from the events payload. This is a common
// field for both issues and pull_request payload so it
// will make things easier to work with.
type Repository struct {
	URL string `json:"url"`
}

// GetPayloadInfo Returns Event  using the json event payload
func GetPayloadInfo(payloadBytes []byte) (*Payload, error) {
	var payload Payload

	err := json.Unmarshal(payloadBytes, &payload)

	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling the []byte of payload to Payload: %v", err)
	}

	return &payload, nil
}
