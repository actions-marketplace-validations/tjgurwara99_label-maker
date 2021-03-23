package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	var event map[string]interface{}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatalf("Error converting json payload to []byte: %v", err)
	}

	err = json.Unmarshal(byteValue, &event)

	if err != nil {
		log.Fatalf("Error Unmarshalling the []byte of json payload to map[string]interface{}: %v", err)
	}

	issue, ok := event["issue"]

	if !ok {
		log.Fatalf("Couldn't find issues field in payload: %v", err)
	}

	var repositoryURL interface{}

	var URL interface{}

	var issueTitle interface{}

	switch issue.(type) {
	case map[string]interface{}:
		repositoryURL, ok = issue.(map[string]interface{})["repository_url"]
		if !ok {
			log.Fatalf("Couldn't repository url field in issue map[string]interface{}: %v", err)
		}
		issueTitle, ok = issue.(map[string]interface{})["title"]
		if !ok {
			log.Fatalf("Couldn't find title field in issue map[string]interface{}: %v", err)
		}
		URL, ok = issue.(map[string]interface{})["url"]
		if !ok {
			log.Fatalf("Couldn't find url field in issue map[string]interface{}: %v", err)
		}
	default:
		log.Fatal("Issue payload is not of type map[string]interface{}")
	}

	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("Couldn't get environment variable repo-token")
	}

	token = fmt.Sprintf("bearer %v", token)

	labels, err := github.GetLabels(repositoryURL.(string), token) // maybe add a check to make sure repositoryURL type is string

	if err != nil {
		log.Fatalf("Couldn't fetch labels: %v", err)
	}

	var updateLabels []string

	for _, label := range labels {
		if !strings.Contains(issueTitle.(string), label.Name) {
			continue
		}
		updateLabels = append(updateLabels, label.Name)
	}

	labelResponse, err := json.Marshal(map[string]interface{}{
		"label": updateLabels,
	})

	if err != nil {
		log.Fatalf("Error marshalling labels: %v", err)
	}

	responseBody := bytes.NewBuffer(labelResponse)

	url := fmt.Sprintf("%s%s", URL.(string), "/labels")

	request, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		log.Fatalf("Error write a new request with labels as buffer: %v", err)
	}

	key := fmt.Sprintf("Bearer %v", os.Getenv("repo-token"))
	request.Header.Add("Authorization", key)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatalf("Response error: %v", err)
	}
	defer response.Body.Close()
	fmt.Printf("%v", response.Body)
	fmt.Println("Successfully added label")
}
