package cloud

import (
	"context"
)

// PriorityService handles priorities for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-Priority
type PriorityService service

// Priority represents a priority of a Jira issue.
// Typical types are "Normal", "Moderate", "Urgent", ...
type Priority struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	StatusColor string `json:"statusColor,omitempty" structs:"statusColor,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
}

// GetList gets all priorities from Jira
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-priority-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *PriorityService) GetList(ctx context.Context) ([]Priority, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
