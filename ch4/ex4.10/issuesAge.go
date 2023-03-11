package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const issuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	query := url.QueryEscape(strings.Join(terms, " "))
	response, err := http.Get(issuesURL + "?q=" + query)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", response.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		response.Body.Close()
		return nil, err

	}

	response.Body.Close()

	return &result, nil
}

func main() {
	result, err := searchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	var issuesOlderThanOneYear []Issue
	var issuesInThePastYear []Issue
	var IssuesInThePastMonth []Issue

	today := time.Now()
	oneMonthAgo := today.AddDate(0, 0, -29)
	OneYearAgo := today.AddDate(-1, 0, 0)

	for _, item := range result.Items {

		switch {
		case item.CreatedAt.Before(OneYearAgo):
			issuesOlderThanOneYear = append(issuesOlderThanOneYear, *item)

		case item.CreatedAt.After(OneYearAgo) && item.CreatedAt.Before(oneMonthAgo):
			issuesInThePastYear = append(issuesInThePastYear, *item)

		case item.CreatedAt.After(oneMonthAgo):
			IssuesInThePastMonth = append(IssuesInThePastMonth, *item)

		}

	}

	if len(issuesOlderThanOneYear) > 0 {
		fmt.Println("Issues older than one year")
		for _, item := range issuesOlderThanOneYear {
			fmt.Printf("#%-5d %9.9s  %s\n",
				item.Number, item.User.Login, item.CreatedAt.Format("2006/01/02"))
		}
	}

	if len(issuesInThePastYear) > 0 {
		fmt.Println("Issues in the past year")
		for _, item := range issuesInThePastYear {
			fmt.Printf("#%-5d %9.9s  %s\n",
				item.Number, item.User.Login, item.CreatedAt.Format("2006/01/02"))
		}
	}

	if len(IssuesInThePastMonth) > 0 {
		fmt.Println("issues in the past month")
		for _, item := range IssuesInThePastMonth {
			fmt.Printf("#%-5d %9.9s  %s\n",
				item.Number, item.User.Login, item.CreatedAt.Format("2006/01/02"))
		}
	}

}
