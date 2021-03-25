package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Label struct for GitHub labels
type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

// GetLabels Get GitHub labels of the repository
func GetLabels(repositoryURL string, token string) ([]Label, error) {
	URL := fmt.Sprintf("%v/labels", repositoryURL)

	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return nil, fmt.Errorf("Couldn't make a new request in GetLabel: %v", err)
	}

	request.Header.Add("Authorization", token)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("Response error in GetLabel: %v", err)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("couldn't convert response body to []byte: %v", err)
	}

	var labels []Label

	err = json.Unmarshal(body, &labels)

	if err != nil {
		return nil, fmt.Errorf("problem unmarshalling the response body: %v", err)
	}

	return labels, nil
}

// AddLabels adds given labels to the issue/pull request and returns the server's response for post processing
func AddLabels(newLabels []string, issueURL, authToken string) (*http.Response, error) {
	labelResponse, err := json.Marshal(map[string][]string{
		"labels": newLabels,
	})

	if err != nil {
		return nil, fmt.Errorf("Error marshalling labels to map[string][]string: %v", err)
	}

	// converting labelResponse to bytes for making a new request
	responseBody := bytes.NewBuffer(labelResponse)

	url := fmt.Sprintf("%s%s", issueURL, "/labels")

	request, err := http.NewRequest("POST", url, responseBody)

	if err != nil {
		return nil, fmt.Errorf("Error: Writing a new request with labels as bytes buffer: %v", err)
	}

	request.Header.Add("Authorization", authToken)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	return http.DefaultClient.Do(request)
}
