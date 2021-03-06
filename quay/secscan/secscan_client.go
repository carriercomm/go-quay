package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new secscan API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secscan API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
Fetches the packages added/removed in the given repo image.
*/
func (a *Client) GetRepoImagePackages(params *GetRepoImagePackagesParams, authInfo client.AuthInfoWriter) (*GetRepoImagePackagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoImagePackagesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getRepoImagePackages",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/image/{imageid}/packages",
		ProducesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoImagePackagesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoImagePackagesOK), nil
}

/*
Fetches the vulnerabilities (if any) for a repository tag.
*/
func (a *Client) GetRepoTagVulnerabilities(params *GetRepoTagVulnerabilitiesParams, authInfo client.AuthInfoWriter) (*GetRepoTagVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoTagVulnerabilitiesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getRepoTagVulnerabilities",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/tag/{tag}/vulnerabilities",
		ProducesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoTagVulnerabilitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoTagVulnerabilitiesOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
