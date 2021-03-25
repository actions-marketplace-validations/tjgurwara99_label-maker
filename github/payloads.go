package github

import (
	"encoding/json"
	"fmt"
)

// Event struct for storing the relavant
// information from the events payload from the json file
type Event struct {
	RepositoryURL string
	Issue         Issue
}

// Issue struct for storing the relavant Issue
// information from the events payload
type Issue struct {
	URL   string
	Title string
	// Body  string
}

// GetEventInfo Returns Event  using the json event payload
func GetEventInfo(payload []byte) (*Event, error) {
	var event map[string]interface{}
	err := json.Unmarshal(payload, &event)

	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling the []byte of payload to map[string]interface{}: %v", err)
	}

	// try to get issue from parsed payload
	fullIssueInfomation, ok := event["issue"]

	if !ok {
		return nil, fmt.Errorf("Couldn't find 'issue' field in the payload: %v", err)
	}

	var repositoryURL interface{}
	var issueURL interface{}
	var issueTitle interface{}
	// var issueBody interface{} // Commenting this out because I don't think the payload contains the body of the issue when issue is edited.

	switch fullIssueInfomation.(type) {
	case map[string]interface{}:
		repositoryURL, ok = fullIssueInfomation.(map[string]interface{})["repository_url"]
		if !ok {
			return nil, fmt.Errorf("Conversion failed: Repository URL field in payload: %v", err)
		}
		issueTitle, ok = fullIssueInfomation.(map[string]interface{})["title"]
		if !ok {
			return nil, fmt.Errorf("Conversion failed: Issue's Title field in payload: %v", err)
		}
		issueURL, ok = fullIssueInfomation.(map[string]interface{})["url"]
		if !ok {
			return nil, fmt.Errorf("Conversion failed: Issue's URL field in payload: %v", err)
		}
	default:
		return nil, fmt.Errorf("Type assertion failed: fullIssuePayload not of type map[string]interface{}")
	}

	// Type assertion block - uneeded but better to be safe then sorry
	// following can be simplified using a function and reflect package
	// but I tend to avoid reflect.

	_, ok = issueTitle.(string)

	if !ok {
		return nil, fmt.Errorf("Type assertion failed: issueTitle not of type string")
	}

	_, ok = issueURL.(string)

	if !ok {
		return nil, fmt.Errorf("Type assertion failed: issueURL not of type string")
	}

	_, ok = repositoryURL.(string)

	if !ok {
		return nil, fmt.Errorf("Type assertion failed: repositoryURL not of type string")
	}

	issue := Issue{
		Title: issueTitle.(string),
		URL:   issueURL.(string),
	}

	return &Event{
		Issue:         issue,
		RepositoryURL: repositoryURL.(string),
	}, nil
}
