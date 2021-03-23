package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getLabels(repositoryURL interface{}) ([]interface{}, error) {

	var url string

	switch repositoryURL.(type) {
	case string:
		url = repositoryURL.(string)
	default:
		fmt.Println("url not a string")
		os.Exit(1)
		//return error instead
	}
	url = fmt.Sprintf("%v/labels", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("couldn't make a new request")
		os.Exit(1)

		// return error instead
	}

	key := fmt.Sprintf("Bearer %v", os.Getenv("GITHUB_TOKEN"))
	request.Header.Add("Authorization", key)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		// something
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		// something
	}

	var labels []interface{}

	err = json.Unmarshal(body, &labels)

	if err != nil {
		// return error
	}

	return labels, nil
}

func main() {
	eventString := os.Getenv("GITHUB_EVENT_PATH")
	var event map[string]interface{}
	err := json.Unmarshal([]byte(eventString), &event)

	if err != nil {
		fmt.Println(err)
		os.Exit(1) // so that it raises error in the github action
	}

	issue, ok := event["issue"]

	if !ok {
		fmt.Println("issue info: error parsing event hooks")
		os.Exit(1)
	}

	var repositoryURL interface{}

	var URL interface{}

	var issueTitle interface{}

	switch issue.(type) {
	case map[string]interface{}:
		repositoryURL, ok = issue.(map[string]interface{})["repository_url"]
		if !ok {
			fmt.Println("repository url: unable to obtain repo url: problem with Unmarshalling")
			os.Exit(1)
		}
		issueTitle, ok = issue.(map[string]interface{})["title"]
		if !ok {
			// something
		}
		URL, ok = issue.(map[string]interface{})["url"]
		if !ok {
			// Something
		}
	default:
		fmt.Println("repository url: unable to obtain repo url: problem with Unmarshalling")
		os.Exit(1)
	}

	labels, err := getLabels(repositoryURL)

	if err != nil {
		// something
	}

	var updateLabels []map[string]interface{}

	for _, label := range labels {
		var stringLabel string
		switch label.(type) {
		case map[string]interface{}:
			stringLabel = label.(map[string]interface{})["name"].(string)
			if strings.Contains(issueTitle.(string), stringLabel) {
				updateLabels = append(updateLabels, label.(map[string]interface{}))
			}
		default:
			fmt.Println("error")
			os.Exit(1)
		}
	}

	labelResponse, err := json.Marshal(map[string]interface{}{
		"label": updateLabels,
	})

	if err != nil {
		// something
	}

	responseBody := bytes.NewBuffer(labelResponse)

	_, err = http.Post(URL.(string), "application/json", responseBody)

	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	fmt.Println("Successfully added label")
}
