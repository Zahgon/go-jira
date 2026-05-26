package onpremise

import (
	"context"
)

// ServiceDeskService handles ServiceDesk for the Jira instance / API.
type ServiceDeskService service

// ServiceDeskOrganizationDTO is a DTO for ServiceDesk organizations
type ServiceDeskOrganizationDTO struct {
	OrganizationID int `json:"organizationId,omitempty" structs:"organizationId,omitempty"`
}

// GetOrganizations returns a list of
// all organizations associated with a service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) GetOrganizations(ctx context.Context, serviceDeskID interface{}, start int, limit int, accountID string) (*PagedDTO, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// AddOrganization adds an organization to
// a service desk. If the organization ID is already
// associated with the service desk, no change is made
// and the resource returns a 204 success code.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-post
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) AddOrganization(ctx context.Context, serviceDeskID interface{}, organizationID int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveOrganization removes an organization
// from a service desk. If the organization ID does not
// match an organization associated with the service desk,
// no change is made and the resource returns a 204 success code.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) RemoveOrganization(ctx context.Context, serviceDeskID interface{}, organizationID int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddCustomers adds customers to the given service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) AddCustomers(ctx context.Context, serviceDeskID interface{}, acountIDs ...string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveCustomers removes customers to the given service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-delete
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) RemoveCustomers(ctx context.Context, serviceDeskID interface{}, acountIDs ...string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListCustomers lists customers for a ServiceDesk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) ListCustomers(ctx context.Context, serviceDeskID interface{}, options *CustomerListOptions) (*CustomerList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// this is an experiemntal endpoint
