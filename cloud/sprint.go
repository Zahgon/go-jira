package cloud

import (
	"context"
)

// SprintService handles sprints in Jira Agile API.
// See https://docs.atlassian.com/jira-software/REST/cloud/
type SprintService service

// IssuesWrapper represents a wrapper struct for moving issues to sprint
type IssuesWrapper struct {
	Issues []string `json:"issues"`
}

// IssuesInSprintResult represents a wrapper struct for search result
type IssuesInSprintResult struct {
	Issues []Issue `json:"issues"`
}

// MoveIssuesToSprint moves issues to a sprint, for a given sprint Id.
// Issues can only be moved to open or active sprints.
// The maximum number of issues that can be moved in one operation is 50.
//
// Jira API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/sprint-moveIssuesToSprint
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *SprintService) MoveIssuesToSprint(ctx context.Context, sprintID int, issueIDs []string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetIssuesForSprint returns all issues in a sprint, for a given sprint Id.
// This only includes issues that the user has permission to view.
// By default, the returned issues are ordered by rank.
//
// Jira API Docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/sprint-getIssuesForSprint
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *SprintService) GetIssuesForSprint(ctx context.Context, sprintID int) ([]Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetIssue returns a full representation of the issue for the given issue key.
// Jira will attempt to identify the issue by the issueIdOrKey path parameter.
// This can be an issue id, or an issue key.
// If the issue cannot be found via an exact match, Jira will also look for the issue in a case-insensitive way, or by looking to see if the issue was moved.
//
// # The given options will be appended to the query string
//
// Jira API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/issue-getIssue
//
// TODO: create agile service for holding all agile apis' implementation
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *SprintService) GetIssue(ctx context.Context, issueID string, options *GetQueryOptions) (*Issue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
