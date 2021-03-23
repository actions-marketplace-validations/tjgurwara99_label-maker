package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getLabels(repositoryURL interface{}) ([]map[string]interface{}, error) {

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
		fmt.Println(err)
	}

	fmt.Println(body)

	var labels []map[string]interface{}

	err = json.Unmarshal(body, &labels)

	if err != nil {
		// return error
		fmt.Println(err)
	}
	fmt.Printf("%s", labels)

	return labels, nil
}

func main() {
	eventString := os.Getenv("GITHUB_EVENT_PATH")
	jsonFile, err := os.Open(eventString)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var event map[string]interface{}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(byteValue, &event)

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

	var updateLabels []string

	for key, value := range labels {
		fmt.Printf("%s %s\n", key, value)
	}

	fmt.Printf("%s\n", updateLabels)

	labelResponse, err := json.Marshal(map[string]interface{}{
		"label": updateLabels,
	})

	if err != nil {
		// something
	}

	responseBody := bytes.NewBuffer(labelResponse)

	url := fmt.Sprintf("%s%s", URL.(string), "/labels")

	fmt.Println(url)

	request, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		fmt.Println("couldn't make a new request")
		os.Exit(1)
	}

	key := fmt.Sprintf("Bearer %v", os.Getenv("GITHUB_TOKEN"))
	request.Header.Add("Authorization", key)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	defer response.Body.Close()
	fmt.Println(ioutil.ReadAll(response.Body))
	fmt.Println("Successfully added label")
}
