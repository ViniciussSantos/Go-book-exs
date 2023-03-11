package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

type IssueResponse struct {
	Issues []Issue
}

type MilestoneResponse struct {
	Milestones []Milestone
}

func newIssueResponse(owner, repo string) (ir IssueResponse, err error) {
	issues, err := GetIssues(owner, repo)
	if err != nil {
		return
	}
	ir.Issues = *issues

	return ir, err
}

func (ir IssueResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Print(issueListTemplate.Execute(w, ir))

	return
}

func main() {

	owner := os.Args[1]
	repo := os.Args[2]

	issues, err := newIssueResponse(owner, repo)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/issues", issues)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
