package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetIssues(owner, repo string) (*[]Issue, error) {
	url := strings.Join([]string{githubURL, "repos", owner, repo, "issues"}, "/")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
