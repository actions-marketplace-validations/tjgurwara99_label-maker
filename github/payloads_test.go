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
		Title: "Spelling error in the README file",
		URL:   "https://api.github.com/repos/Codertocat/Hello-World/issues/1",
		Body:  "It looks like you accidentally spelled 'commit' with two 't's.",
	}

	testRepository := github.Repository{
		URL: "https://api.github.com/repos/Codertocat/Hello-World",
	}

	testCasePullRequest := github.PullRequest{}

	testCase := github.Payload{
		Action:     "edited",
		Issue:      testCaseIssue,
		Repository: testRepository,
	}

	if event.Issue != testCase.Issue {
		t.Errorf("Test Failed: Issue not equivalent: Expected: %#v,\nReceived : %#v\n", testCaseIssue, event.Issue)
	}

	if event.Action != testCase.Action {
		t.Errorf("Test Failed: Repository URL not equivalent: Expected: %#v,\nReceived : %#v\n", testCase.Action, event.Action)
	}

	if event.PullRequest != testCasePullRequest {
		t.Errorf("Test Failed: Pull Request information not equivalent: Expected: %#v,\nReceived: %#v\n", testCasePullRequest, event.PullRequest)
	}
}
