package onpremise

import (
	"context"
	"net/http"
)

const (
	// HTTP Basic Authentication
	authTypeBasic = 1
	// HTTP Session Authentication
	authTypeSession = 2
)

// AuthenticationService handles authentication for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#authentication
type AuthenticationService struct {
	client *Client

	// Authentication type
	authType int

	// Basic auth username
	username string

	// Basic auth password
	password string
}

// Session represents a Session JSON response by the Jira API.
type Session struct {
	Self    string `json:"self,omitempty"`
	Name    string `json:"name,omitempty"`
	Session struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"session,omitempty"`
	LoginInfo struct {
		FailedLoginCount    int    `json:"failedLoginCount"`
		LoginCount          int    `json:"loginCount"`
		LastFailedLoginTime string `json:"lastFailedLoginTime"`
		PreviousLoginTime   string `json:"previousLoginTime"`
	} `json:"loginInfo"`
	Cookies []*http.Cookie
}

// AcquireSessionCookie creates a new session for a user in Jira.
// Once a session has been successfully created it can be used to access any of Jira's remote APIs and also the web UI by passing the appropriate HTTP Cookie header.
// The header will by automatically applied to every API request.
// Note that it is generally preferrable to use HTTP BASIC authentication with the REST API.
// However, this resource may be used to mimic the behaviour of Jira's log-in page (e.g. to display log-in errors to a user).
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
//
// Deprecated: Use CookieAuthTransport instead
func (s *AuthenticationService) AcquireSessionCookie(ctx context.Context, username, password string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SetBasicAuth sets username and password for the basic auth against the Jira instance.
//
// Deprecated: Use BasicAuthTransport instead
func (s *AuthenticationService) SetBasicAuth(username, password string) {
	_ = "STUB: not implemented"
	return
}

// Authenticated reports if the current Client has authentication details for Jira
func (s *AuthenticationService) Authenticated() bool { _ = "STUB: not implemented"; return false }

// Logout logs out the current user that has been authenticated and the session in the client is destroyed.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
//
// Deprecated: Use CookieAuthTransport to create base client.  Logging out is as simple as not using the
// client anymore
func (s *AuthenticationService) Logout(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If logout successful, delete session

// GetCurrentUser gets the details of the current user.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
func (s *AuthenticationService) GetCurrentUser(ctx context.Context) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
