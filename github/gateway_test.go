package github_test

import (
	"fmt"
	"testing"

	"github.com/tjgurwara99/label-maker/github"
)

func TestGetLabel(t *testing.T) {
	labels, err := github.GetLabels("https://api.github.com/repos/tjgurwara99/label-maker")

	if err != nil {
		t.Errorf("Something wrong with GetLabel function: %v", err)
	}

	fmt.Println(labels[0].Name)
}
