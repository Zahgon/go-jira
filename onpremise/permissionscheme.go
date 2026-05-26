package onpremise

import (
	"context"
)

// PermissionSchemeService handles permissionschemes for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-group-Permissionscheme
type PermissionSchemeService service

type PermissionSchemes struct {
	PermissionSchemes []PermissionScheme `json:"permissionSchemes" structs:"permissionSchemes"`
}

type Permission struct {
	ID     int    `json:"id" structs:"id"`
	Self   string `json:"expand" structs:"expand"`
	Holder Holder `json:"holder" structs:"holder"`
	Name   string `json:"permission" structs:"permission"`
}

type Holder struct {
	Type      string `json:"type" structs:"type"`
	Parameter string `json:"parameter" structs:"parameter"`
	Expand    string `json:"expand" structs:"expand"`
}

// GetList returns a list of all permission schemes
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-api-3-permissionscheme-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *PermissionSchemeService) GetList(ctx context.Context) (*PermissionSchemes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get returns a full representation of the permission scheme for the schemeID
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-api-3-permissionscheme-schemeId-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *PermissionSchemeService) Get(ctx context.Context, schemeID int) (*PermissionScheme, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
