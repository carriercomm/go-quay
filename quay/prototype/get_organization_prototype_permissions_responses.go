package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// GetOrganizationPrototypePermissionsReader is a Reader for the GetOrganizationPrototypePermissions structure.
type GetOrganizationPrototypePermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetOrganizationPrototypePermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationPrototypePermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganizationPrototypePermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrganizationPrototypePermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganizationPrototypePermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationPrototypePermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationPrototypePermissionsOK creates a GetOrganizationPrototypePermissionsOK with default headers values
func NewGetOrganizationPrototypePermissionsOK() *GetOrganizationPrototypePermissionsOK {
	return &GetOrganizationPrototypePermissionsOK{}
}

/*GetOrganizationPrototypePermissionsOK handles this case with default header values.

Successful invocation
*/
type GetOrganizationPrototypePermissionsOK struct {
}

func (o *GetOrganizationPrototypePermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/prototypes][%d] getOrganizationPrototypePermissionsOK ", 200)
}

func (o *GetOrganizationPrototypePermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPrototypePermissionsBadRequest creates a GetOrganizationPrototypePermissionsBadRequest with default headers values
func NewGetOrganizationPrototypePermissionsBadRequest() *GetOrganizationPrototypePermissionsBadRequest {
	return &GetOrganizationPrototypePermissionsBadRequest{}
}

/*GetOrganizationPrototypePermissionsBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganizationPrototypePermissionsBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetOrganizationPrototypePermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/prototypes][%d] getOrganizationPrototypePermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationPrototypePermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationPrototypePermissionsUnauthorized creates a GetOrganizationPrototypePermissionsUnauthorized with default headers values
func NewGetOrganizationPrototypePermissionsUnauthorized() *GetOrganizationPrototypePermissionsUnauthorized {
	return &GetOrganizationPrototypePermissionsUnauthorized{}
}

/*GetOrganizationPrototypePermissionsUnauthorized handles this case with default header values.

Session required
*/
type GetOrganizationPrototypePermissionsUnauthorized struct {
}

func (o *GetOrganizationPrototypePermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/prototypes][%d] getOrganizationPrototypePermissionsUnauthorized ", 401)
}

func (o *GetOrganizationPrototypePermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPrototypePermissionsForbidden creates a GetOrganizationPrototypePermissionsForbidden with default headers values
func NewGetOrganizationPrototypePermissionsForbidden() *GetOrganizationPrototypePermissionsForbidden {
	return &GetOrganizationPrototypePermissionsForbidden{}
}

/*GetOrganizationPrototypePermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrganizationPrototypePermissionsForbidden struct {
}

func (o *GetOrganizationPrototypePermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/prototypes][%d] getOrganizationPrototypePermissionsForbidden ", 403)
}

func (o *GetOrganizationPrototypePermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationPrototypePermissionsNotFound creates a GetOrganizationPrototypePermissionsNotFound with default headers values
func NewGetOrganizationPrototypePermissionsNotFound() *GetOrganizationPrototypePermissionsNotFound {
	return &GetOrganizationPrototypePermissionsNotFound{}
}

/*GetOrganizationPrototypePermissionsNotFound handles this case with default header values.

Not found
*/
type GetOrganizationPrototypePermissionsNotFound struct {
}

func (o *GetOrganizationPrototypePermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/prototypes][%d] getOrganizationPrototypePermissionsNotFound ", 404)
}

func (o *GetOrganizationPrototypePermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
