package onpremise

import (
	"context"
)

// GroupService handles Groups for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/server/#api/2/group
type GroupService service

// groupMembersResult is only a small wrapper around the Group* methods
// to be able to parse the results
type groupMembersResult struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Members    []GroupMember `json:"values"`
}

// Group represents a Jira group
type Group struct {
	ID                   string          `json:"id"`
	Title                string          `json:"title"`
	Type                 string          `json:"type"`
	Properties           groupProperties `json:"properties"`
	AdditionalProperties bool            `json:"additionalProperties"`
}

type groupProperties struct {
	Name groupPropertiesName `json:"name"`
}

type groupPropertiesName struct {
	Type string `json:"type"`
}

// GroupMember reflects a single member of a group
type GroupMember struct {
	Self         string `json:"self,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
	AccountType  string `json:"accountType,omitempty"`
}

// GroupSearchOptions specifies the optional parameters for the Get Group methods
type GroupSearchOptions struct {
	StartAt              int
	MaxResults           int
	IncludeInactiveUsers bool
}

// Get returns a paginated list of members of the specified group and its subgroups.
// Users in the page are ordered by user names.
// User of this resource is required to have sysadmin or admin permissions.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/server/#api/2/group-getUsersFromGroup
//
// WARNING: This API only returns the first page of group members
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *GroupService) Get(ctx context.Context, name string, options *GroupSearchOptions) ([]GroupMember, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// TODO use addOptions

// Add adds user to group
//
// Jira API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/group-addUserToGroup
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *GroupService) Add(ctx context.Context, groupname string, username string) (*Group, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Remove removes user from group
//
// Jira API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/group-removeUserFromGroup
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *GroupService) Remove(ctx context.Context, groupname string, username string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
