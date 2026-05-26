package cloud

import (
	"context"
)

// OrganizationService handles Organizations for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/
type OrganizationService service

// OrganizationCreationDTO is DTO for creat organization API
type OrganizationCreationDTO struct {
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

// SelfLink Stores REST API URL to the organization.
type SelfLink struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
}

// Organization contains Organization data
type Organization struct {
	ID    string    `json:"id,omitempty" structs:"id,omitempty"`
	Name  string    `json:"name,omitempty" structs:"name,omitempty"`
	Links *SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// OrganizationUsersDTO contains organization user ids
type OrganizationUsersDTO struct {
	AccountIds []string `json:"accountIds,omitempty" structs:"accountIds,omitempty"`
}

// PagedDTO is response of a paged list
type PagedDTO struct {
	Size       int           `json:"size,omitempty" structs:"size,omitempty"`
	Start      int           `json:"start,omitempty" structs:"start,omitempty"`
	Limit      int           `limit:"size,omitempty" structs:"limit,omitempty"`
	IsLastPage bool          `json:"isLastPage,omitempty" structs:"isLastPage,omitempty"`
	Values     []interface{} `values:"isLastPage,omitempty" structs:"values,omitempty"`
	Expands    []string      `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// PropertyKey contains Property key details.
type PropertyKey struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Key  string `json:"key,omitempty" structs:"key,omitempty"`
}

// PropertyKeys contains an array of PropertyKey
type PropertyKeys struct {
	Keys []PropertyKey `json:"keys,omitempty" structs:"keys,omitempty"`
}

// GetAllOrganizations returns a list of organizations in
// the Jira Service Management instance.
// Use this method when you want to present a list
// of organizations or want to locate an organization
// by name.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-group-organization
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetAllOrganizations(ctx context.Context, start int, limit int, accountID string) (*PagedDTO, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateOrganization creates an organization by
// passing the name of the organization.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) CreateOrganization(ctx context.Context, name string) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetOrganization returns details of an
// organization. Use this method to get organization
// details whenever your application component is
// passed an organization ID but needs to display
// other organization details.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetOrganization(ctx context.Context, organizationID int) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DeleteOrganization deletes an organization. Note that
// the organization is deleted regardless
// of other associations it may have.
// For example, associations with service desks.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, organizationID int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPropertiesKeys returns the keys of
// all properties for an organization. Use this resource
// when you need to find out what additional properties
// items have been added to an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetPropertiesKeys(ctx context.Context, organizationID int) (*PropertyKeys, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetProperty returns the value of a property
// from an organization. Use this method to obtain the JSON
// content for an organization's property.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-propertykey-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetProperty(ctx context.Context, organizationID int, propertyKey string) (*EntityProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// SetProperty sets the value of a
// property for an organization. Use this
// resource to store custom data against an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-propertykey-put
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
// Caller must close resp.Body
func (s *OrganizationService) SetProperty(ctx context.Context, organizationID int, propertyKey string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteProperty removes a property from an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-propertykey-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) DeleteProperty(ctx context.Context, organizationID int, propertyKey string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUsers returns all the users
// associated with an organization. Use this
// method where you want to provide a list of
// users for an organization or determine if
// a user is associated with an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-user-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetUsers(ctx context.Context, organizationID int, start int, limit int) (*PagedDTO, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// AddUsers adds users to an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-user-post
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) AddUsers(ctx context.Context, organizationID int, users OrganizationUsersDTO) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveUsers removes users from an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-user-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) RemoveUsers(ctx context.Context, organizationID int, users OrganizationUsersDTO) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
