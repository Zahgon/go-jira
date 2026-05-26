package onpremise

import (
	"net/http"
	"net/url"
)

// JWTAuthTransport is an http.RoundTripper that authenticates all requests
// using Jira's JWT based authentication.
//
// NOTE: this form of auth should be used by add-ons installed from the Atlassian marketplace.
//
// Jira docs: https://developer.atlassian.com/cloud/jira/platform/understanding-jwt
// Examples in other languages:
//
//	https://bitbucket.org/atlassian/atlassian-jwt-ruby/src/d44a8e7a4649e4f23edaa784402655fda7c816ea/lib/atlassian/jwt.rb
//	https://bitbucket.org/atlassian/atlassian-jwt-py/src/master/atlassian_jwt/url_utils.py
type JWTAuthTransport struct {
	Secret []byte
	Issuer string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func (t *JWTAuthTransport) Client() *http.Client { _ = "STUB: not implemented"; return nil }

func (t *JWTAuthTransport) transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

// RoundTrip adds the session object to the request.
func (t *JWTAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil,
		// per RoundTripper contract
		nil
}

func (t *JWTAuthTransport) createQueryStringHash(httpMethod string, jiraURL *url.URL) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *JWTAuthTransport) canonicalizeRequest(httpMethod string, jiraURL *url.URL) string {
	_ = "STUB: not implemented"
	return ""
}
