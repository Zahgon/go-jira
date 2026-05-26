package onpremise

import (
	"context"

	"github.com/trivago/tgo/tcontainer"
)

// CreateMetaInfo contains information about fields and their attributed to create a ticket.
type CreateMetaInfo struct {
	Expand   string         `json:"expand,omitempty"`
	Projects []*MetaProject `json:"projects,omitempty"`
}

// EditMetaInfo contains information about fields and their attributed to edit a ticket.
type EditMetaInfo struct {
	Fields tcontainer.MarshalMap `json:"fields,omitempty"`
}

// MetaProject is the meta information about a project returned from createmeta api
type MetaProject struct {
	Expand string `json:"expand,omitempty"`
	Self   string `json:"self,omitempty"`
	Id     string `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	// omitted avatarUrls
	IssueTypes []*MetaIssueType `json:"issuetypes,omitempty"`
}

// MetaIssueType represents the different issue types a project has.
//
// Note: Fields is interface because this is an object which can
// have arbitraty keys related to customfields. It is not possible to
// expect these for a general way. This will be returning a map.
// Further processing must be done depending on what is required.
type MetaIssueType struct {
	Self        string                `json:"self,omitempty"`
	Id          string                `json:"id,omitempty"`
	Description string                `json:"description,omitempty"`
	IconUrl     string                `json:"iconurl,omitempty"`
	Name        string                `json:"name,omitempty"`
	Subtasks    bool                  `json:"subtask,omitempty"`
	Expand      string                `json:"expand,omitempty"`
	Fields      tcontainer.MarshalMap `json:"fields,omitempty"`
}

// GetCreateMeta makes the api call to get the meta information without requiring to have a projectKey
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueService) GetCreateMeta(ctx context.Context, options *GetQueryOptions) (*CreateMetaInfo, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetEditMeta makes the api call to get the edit meta information for an issue
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueService) GetEditMeta(ctx context.Context, issue *Issue) (*EditMetaInfo, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetProjectWithName returns a project with "name" from the meta information received. If not found, this returns nil.
// The comparison of the name is case insensitive.
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (m *CreateMetaInfo) GetProjectWithName(name string) *MetaProject {
	_ = "STUB: not implemented"
	return nil
}

// GetProjectWithKey returns a project with "name" from the meta information received. If not found, this returns nil.
// The comparison of the name is case insensitive.
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (m *CreateMetaInfo) GetProjectWithKey(key string) *MetaProject {
	_ = "STUB: not implemented"
	return nil
}

// GetIssueTypeWithName returns an IssueType with name from a given MetaProject. If not found, this returns nil.
// The comparison of the name is case insensitive
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (p *MetaProject) GetIssueTypeWithName(name string) *MetaIssueType {
	_ = "STUB: not implemented"
	return nil
}

// GetMandatoryFields returns a map of all the required fields from the MetaIssueTypes.
// if a field returned by the api was:
//
//	"customfield_10806": {
//						"required": true,
//						"schema": {
//							"type": "any",
//							"custom": "com.pyxis.greenhopper.jira:gh-epic-link",
//							"customId": 10806
//						},
//						"name": "Epic Link",
//						"hasDefaultValue": false,
//						"operations": [
//							"set"
//						]
//					}
//
// the returned map would have "Epic Link" as the key and "customfield_10806" as value.
// This choice has been made so that the it is easier to generate the create api request later.
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (t *MetaIssueType) GetMandatoryFields() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllFields returns a map of all the fields for an IssueType. This includes all required and not required.
// The key of the returned map is what you see in the form and the value is how it is representated in the jira schema.
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (t *MetaIssueType) GetAllFields() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckCompleteAndAvailable checks if the given fields satisfies the mandatory field required to create a issue for the given type
// And also if the given fields are available.
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (t *MetaIssueType) CheckCompleteAndAvailable(config map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check templateconfig against mandatory fields

// check templateConfig against all fields to verify they are available
