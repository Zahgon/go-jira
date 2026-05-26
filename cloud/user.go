package cloud

import (
	"context"
)

// UserService handles users for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-group-Users
type UserService service

// User represents a Jira user.
type User struct {
	Self             string           `json:"self,omitempty" structs:"self,omitempty"`
	AccountID        string           `json:"accountId,omitempty" structs:"accountId,omitempty"`
	AccountType      string           `json:"accountType,omitempty" structs:"accountType,omitempty"`
	Name             string           `json:"name,omitempty" structs:"name,omitempty"`
	Key              string           `json:"key,omitempty" structs:"key,omitempty"`
	Password         string           `json:"-"`
	EmailAddress     string           `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	AvatarUrls       AvatarUrls       `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName      string           `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active           bool             `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone         string           `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Locale           string           `json:"locale,omitempty" structs:"locale,omitempty"`
	Groups           UserGroups       `json:"groups,omitempty" structs:"groups,omitempty"`
	ApplicationRoles ApplicationRoles `json:"applicationRoles,omitempty" structs:"applicationRoles,omitempty"`
}

// UserGroup represents the group list
type UserGroup struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

// Groups is a wrapper for UserGroup
type UserGroups struct {
	Size  int         `json:"size,omitempty" structs:"size,omitempty"`
	Items []UserGroup `json:"items,omitempty" structs:"items,omitempty"`
}

// ApplicationRoles is a wrapper for ApplicationRole
type ApplicationRoles struct {
	Size  int               `json:"size,omitempty" structs:"size,omitempty"`
	Items []ApplicationRole `json:"items,omitempty" structs:"items,omitempty"`
}

// ApplicationRole represents a role assigned to a user
type ApplicationRole struct {
	Key                  string   `json:"key"`
	Groups               []string `json:"groups"`
	Name                 string   `json:"name"`
	DefaultGroups        []string `json:"defaultGroups"`
	SelectedByDefault    bool     `json:"selectedByDefault"`
	Defined              bool     `json:"defined"`
	NumberOfSeats        int      `json:"numberOfSeats"`
	RemainingSeats       int      `json:"remainingSeats"`
	UserCount            int      `json:"userCount"`
	UserCountDescription string   `json:"userCountDescription"`
	HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats"`
	Platform             bool     `json:"platform"`

	// Key `groupDetails` missing - https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
	// Key `defaultGroupsDetails` missing - https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-application-roles/#api-rest-api-3-applicationrole-key-get
}

type UserSearchParam struct {
	name  string
	value string
}

type UserSearch []UserSearchParam

type UserSearchF func(UserSearch) UserSearch

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

// GetCurrentUser returns details for the current user.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-myself-get
func (s *UserService) GetCurrentUser(ctx context.Context) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// WithMaxResults sets the max results to return
func WithMaxResults(maxResults int) UserSearchF {
	_ = "STUB: not implemented"
	return *new(UserSearchF)
}

// WithStartAt set the start pager
func WithStartAt(startAt int) UserSearchF { _ = "STUB: not implemented"; return *new(UserSearchF) }

// WithActive sets the active users lookup
func WithActive(active bool) UserSearchF { _ = "STUB: not implemented"; return *new(UserSearchF) }

// WithInactive sets the inactive users lookup
func WithInactive(inactive bool) UserSearchF { _ = "STUB: not implemented"; return *new(UserSearchF) }

// WithUsername sets the username to search
func WithUsername(username string) UserSearchF { _ = "STUB: not implemented"; return *new(UserSearchF) }

// WithAccountId sets the account id to search
func WithAccountId(accountId string) UserSearchF {
	_ = "STUB: not implemented"
	return *new(UserSearchF)
}

// WithProperty sets the property (Property keys are specified by path) to search
func WithProperty(property string) UserSearchF { _ = "STUB: not implemented"; return *new(UserSearchF) }

// Find searches for user info from Jira:
// It can find users by email or display name using the query parameter
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-user-search-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *UserService) Find(ctx context.Context, property string, tweaks ...UserSearchF) ([]User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
