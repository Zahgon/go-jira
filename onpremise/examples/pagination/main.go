package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

// GetAllIssues will implement pagination of api and get all the issues.
// Jira API has limitation as to maxResults it can return at one time.
// You may have usecase where you need to get all the issues according to jql
// This is where this example comes in.
func GetAllIssues(client *jira.Client, searchString string) ([]jira.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Max results can go up to 1000

func main() {
	jiraClient, err := jira.NewClient("https://issues.apache.org/jira/", nil)
	if err != nil {
		panic(err)
	}

	jql := "project = Mesos and type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)

	issues, err := GetAllIssues(jiraClient, jql)
	if err != nil {
		panic(err)
	}
	fmt.Println(issues)

}
