package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/tjgurwara99/label-maker/github"
)

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	jsonFile, err := os.Open(eventPath)

	if err != nil {
		log.Fatalf("Error opening event Payload: %v\n", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatalf("Error converting json payload to []byte: %v", err)
	}

	payload, err := github.GetPayloadInfo(byteValue)

	if err != nil {
		log.Fatalf("Couldn't get event payload stored in Event struct: %v", err)
	}

	token := os.Getenv("INPUT_TOKEN")

	if token == "" {
		log.Fatal("Couldn't get environment variable TOKEN")
	}

	token = fmt.Sprintf("bearer %v", token)

	labels, err := github.GetLabels(payload.Issue.RepositoryURL, token)

	if err != nil {
		log.Fatalf("Couldn't fetch labels: %v", err)
	}

	var newLabels []string

	for _, label := range labels {
		if !strings.Contains(strings.ToLower(payload.Issue.Title), strings.ToLower(label.Name)) {
			continue
		}
		newLabels = append(newLabels, label.Name)
	}

	response, err := github.AddLabels(newLabels, payload.Issue.URL, token)

	if err != nil {
		log.Fatalf("Response error: %v", err)
	}
	defer response.Body.Close()

	fmt.Println("Successfully added label")
}
