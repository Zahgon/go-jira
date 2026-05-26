package cloud

import "net/http"

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and a Personal API Token.
//
// Jira docs: https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/
// Create a token: https://id.atlassian.com/manage-profile/security/api-tokens
type BasicAuthTransport struct {
	Username string
	APIToken string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.  We just add the
// basic auth information and return the RoundTripper for this transport type.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil,
		// per RoundTripper contract
		nil
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.  This is a nice little bit of sugar
// so we can just get the client instead of creating the client in the calling code.
// If it's necessary to send more information on client init, the calling code can
// always skip this and set the transport itself.
func (t *BasicAuthTransport) Client() *http.Client { _ = "STUB: not implemented"; return nil }

func (t *BasicAuthTransport) transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}
