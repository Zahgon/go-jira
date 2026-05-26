package onpremise

import (
	"context"
)

// UserService handles users for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-group-Users
type UserService service

// User represents a Jira user.
type User struct {
	Self            string     `json:"self,omitempty" structs:"self,omitempty"`
	AccountID       string     `json:"accountId,omitempty" structs:"accountId,omitempty"`
	AccountType     string     `json:"accountType,omitempty" structs:"accountType,omitempty"`
	Name            string     `json:"name,omitempty" structs:"name,omitempty"`
	Key             string     `json:"key,omitempty" structs:"key,omitempty"`
	Password        string     `json:"-"`
	EmailAddress    string     `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	AvatarUrls      AvatarUrls `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName     string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active          bool       `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone        string     `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Locale          string     `json:"locale,omitempty" structs:"locale,omitempty"`
	ApplicationKeys []string   `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
}

// UserGroup represents the group list
type UserGroup struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

type userSearchParam struct {
	name  string
	value string
}

type userSearch []userSearchParam

type userSearchF func(userSearch) userSearch

// Get gets user info from Jira using its Account Id
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-user-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) Get(ctx context.Context, accountId string) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetByAccountID gets user info from Jira
// Searching by another parameter that is not accountId is deprecated,
// but this method is kept for backwards compatibility
// Jira API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-getUser
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) GetByAccountID(ctx context.Context, accountID string) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create creates an user in Jira.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-createUser
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) Create(ctx context.Context, user *User) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes an user from Jira.
// Returns http.StatusNoContent on success.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-user-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) Delete(ctx context.Context, accountId string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGroups returns the groups which the user belongs to
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-user-groups-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) GetGroups(ctx context.Context, accountId string) (*[]UserGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetSelf information about the current logged-in user
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-myself-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) GetSelf(ctx context.Context) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// WithMaxResults sets the max results to return
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithMaxResults(maxResults int) userSearchF {
	_ = "STUB: not implemented"
	return *new(userSearchF)
}

// WithStartAt set the start pager
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithStartAt(startAt int) userSearchF { _ = "STUB: not implemented"; return *new(userSearchF) }

// WithActive sets the active users lookup
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithActive(active bool) userSearchF { _ = "STUB: not implemented"; return *new(userSearchF) }

// WithInactive sets the inactive users lookup
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithInactive(inactive bool) userSearchF { _ = "STUB: not implemented"; return *new(userSearchF) }

// WithUsername sets the username to search
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithUsername(username string) userSearchF { _ = "STUB: not implemented"; return *new(userSearchF) }

// WithAccountId sets the account id to search
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithAccountId(accountId string) userSearchF {
	_ = "STUB: not implemented"
	return *new(userSearchF)
}

// WithProperty sets the property (Property keys are specified by path) to search
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func WithProperty(property string) userSearchF { _ = "STUB: not implemented"; return *new(userSearchF) }

// Find searches for user info from Jira:
// It can find users by email or display name using the query parameter
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-user-search-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) Find(ctx context.Context, property string, tweaks ...userSearchF) ([]User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
