package cloud

// Error message from Jira
// See https://docs.atlassian.com/jira/REST/cloud/#error-responses
type Error struct {
	HTTPError     error
	ErrorMessages []string          `json:"errorMessages"`
	Errors        map[string]string `json:"errors"`
}

// NewJiraError creates a new jira Error
func NewJiraError(resp *Response, httpError error) error { _ = "STUB: not implemented"; return nil }

// Error is a short string representing the error
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// return fmt.Sprintf("%v", e.HTTPError)

// LongError is a full representation of the error as a string
func (e *Error) LongError() string { _ = "STUB: not implemented"; return "" }
