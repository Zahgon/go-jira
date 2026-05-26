package main

import (
	"context"
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func main() {
	jiraClient, _ := jira.NewClient("https://issues.apache.org/jira/", nil)

	// Running JQL query

	jql := "project = Mesos and type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)
	issues, resp, err := jiraClient.Issue.Search(context.Background(), jql, nil)
	if err != nil {
		panic(err)
	}
	outputResponse(issues, resp)

	fmt.Println("")
	fmt.Println("")

	// Running an empty JQL query to get all tickets
	jql = ""
	fmt.Printf("Usecase: Running an empty JQL query to get all tickets\n")
	issues, resp, err = jiraClient.Issue.Search(context.Background(), jql, nil)
	if err != nil {
		panic(err)
	}
	outputResponse(issues, resp)
}

func outputResponse(issues []jira.Issue, resp *jira.Response) { _ = "STUB: not implemented"; return }
