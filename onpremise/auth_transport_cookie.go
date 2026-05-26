package onpremise

import (
	"net/http"
)

// CookieAuthTransport is an http.RoundTripper that authenticates all requests
// using Jira's cookie-based authentication.
//
// Note that it is generally preferable to use HTTP BASIC authentication with the REST API.
// However, this resource may be used to mimic the behaviour of Jira's log-in page (e.g. to display log-in errors to a user).
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
type CookieAuthTransport struct {
	Username string
	Password string
	AuthURL  string

	// SessionObject is the authenticated cookie string.s
	// It's passed in each call to prove the client is authenticated.
	SessionObject []*http.Cookie

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip adds the session object to the request.
func (t *CookieAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// per RoundTripper contract

// Don't add an empty value cookie to the request

// Client returns an *http.Client that makes requests that are authenticated
// using cookie authentication
func (t *CookieAuthTransport) Client() *http.Client { _ = "STUB: not implemented"; return nil }

// setSessionObject attempts to authenticate the user and set
// the session object (e.g. cookie)
func (t *CookieAuthTransport) setSessionObject() error { _ = "STUB: not implemented"; return nil }

// getAuthRequest assembles the request to get the authenticated cookie
func (t *CookieAuthTransport) buildAuthRequest() (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO Use a context here

func (t *CookieAuthTransport) transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}
