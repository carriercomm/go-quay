package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetUserPermissionsReader struct {
	formats strfmt.Registry
}

func (o *GetUserPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetUserPermissionsOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetUserPermissionsBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserPermissionsBadRequest", &result, response.Code())

	case 401:
		var result GetUserPermissionsUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserPermissionsUnauthorized", &result, response.Code())

	case 403:
		var result GetUserPermissionsForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserPermissionsForbidden", &result, response.Code())

	case 404:
		var result GetUserPermissionsNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserPermissionsNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetUserPermissionsOK struct {
}

func (o *GetUserPermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetUserPermissionsBadRequest struct {
}

func (o *GetUserPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetUserPermissionsUnauthorized struct {
}

func (o *GetUserPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetUserPermissionsForbidden struct {
}

func (o *GetUserPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetUserPermissionsNotFound struct {
}

func (o *GetUserPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}