package onpremise

import (
	"context"
)

// IssueLinkTypeService handles issue link types for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-group-Issue-link-types
type IssueLinkTypeService service

// GetList gets all of the issue link types from Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueLinkTypeService) GetList(ctx context.Context) ([]IssueLinkType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get gets info of a specific issue link type from Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueLinkTypeService) Get(ctx context.Context, ID string) (*IssueLinkType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create creates an issue link type in Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueLinkTypeService) Create(ctx context.Context, linkType *IssueLinkType) (*IssueLinkType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update updates an issue link type.  The issue is found by key.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-put
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueLinkTypeService) Update(ctx context.Context, linkType *IssueLinkType) (*IssueLinkType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes an issue link type based on provided ID.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *IssueLinkTypeService) Delete(ctx context.Context, ID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
