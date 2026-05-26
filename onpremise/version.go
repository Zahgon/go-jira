package onpremise

import (
	"context"
)

// VersionService handles Versions for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/version
type VersionService service

// Version represents a single release version of a project
type Version struct {
	Self            string `json:"self,omitempty" structs:"self,omitempty"`
	ID              string `json:"id,omitempty" structs:"id,omitempty"`
	Name            string `json:"name,omitempty" structs:"name,omitempty"`
	Description     string `json:"description,omitempty" structs:"description,omitempty"`
	Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
	Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
	ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"` // Unlike other IDs, this is returned as a number
	StartDate       string `json:"startDate,omitempty" structs:"startDate,omitempty"`
}

// Get gets version info from Jira
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-version-id-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *VersionService) Get(ctx context.Context, versionID int) (*Version, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create creates a version in Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-version-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *VersionService) Create(ctx context.Context, version *Version) (*Version, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update updates a version from a JSON representation.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-version-id-put
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *VersionService) Update(ctx context.Context, version *Version) (*Version, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// This is just to follow the rest of the API's convention of returning a version.
// Returning the same pointer here is pointless, so we return a copy instead.
