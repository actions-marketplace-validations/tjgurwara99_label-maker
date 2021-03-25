package github_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/tjgurwara99/label-maker/github"
)

func TestGetPayloadInfo(t *testing.T) {
	currentWD, err := os.Getwd()
	if err != nil {
		t.Errorf("Couldn't get the current working directory: %v", err)
	}

	// The following payload is an example payload which has been
	// taken from the GitHub API documentation - you can find it at:
	// https://docs.github.com/en/developers/webhooks-and-events/webhook-events-and-payloads#issues
	jsonFile, err := os.Open(fmt.Sprintf("%s/payload.json", currentWD))

	if err != nil {
		t.Errorf("Error opening payload json: %v", err)
	}

	payloadInBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		t.Errorf("Conversion failed: *os.File to []Byte: %v", err)
	}

	event, err := github.GetPayloadInfo(payloadInBytes)

	if err != nil {
		t.Errorf("Failed to store payload in Event struct: %v", err)
	}

	testCaseIssue := github.Issue{
		Title:         "Spelling error in the README file",
		URL:           "https://api.github.com/repos/Codertocat/Hello-World/issues/1",
		RepositoryURL: "https://api.github.com/repos/Codertocat/Hello-World",
	}

	testCase := github.Payload{
		Action: "edited",
		Issue:  testCaseIssue,
	}

	if event.Issue != testCase.Issue {
		t.Errorf("Test Failed: Issue not equivalent")
	}

	if event.Action != testCase.Action {
		t.Errorf("Test Failed: Repository URL not equivalent")
	}
}
