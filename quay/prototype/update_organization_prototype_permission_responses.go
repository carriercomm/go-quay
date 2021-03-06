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

// UpdateOrganizationPrototypePermissionReader is a Reader for the UpdateOrganizationPrototypePermission structure.
type UpdateOrganizationPrototypePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateOrganizationPrototypePermissionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOrganizationPrototypePermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateOrganizationPrototypePermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateOrganizationPrototypePermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateOrganizationPrototypePermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateOrganizationPrototypePermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOrganizationPrototypePermissionOK creates a UpdateOrganizationPrototypePermissionOK with default headers values
func NewUpdateOrganizationPrototypePermissionOK() *UpdateOrganizationPrototypePermissionOK {
	return &UpdateOrganizationPrototypePermissionOK{}
}

/*UpdateOrganizationPrototypePermissionOK handles this case with default header values.

Successful invocation
*/
type UpdateOrganizationPrototypePermissionOK struct {
}

func (o *UpdateOrganizationPrototypePermissionOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] updateOrganizationPrototypePermissionOK ", 200)
}

func (o *UpdateOrganizationPrototypePermissionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationPrototypePermissionBadRequest creates a UpdateOrganizationPrototypePermissionBadRequest with default headers values
func NewUpdateOrganizationPrototypePermissionBadRequest() *UpdateOrganizationPrototypePermissionBadRequest {
	return &UpdateOrganizationPrototypePermissionBadRequest{}
}

/*UpdateOrganizationPrototypePermissionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateOrganizationPrototypePermissionBadRequest struct {
	Payload *models.GeneralError
}

func (o *UpdateOrganizationPrototypePermissionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] updateOrganizationPrototypePermissionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationPrototypePermissionBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationPrototypePermissionUnauthorized creates a UpdateOrganizationPrototypePermissionUnauthorized with default headers values
func NewUpdateOrganizationPrototypePermissionUnauthorized() *UpdateOrganizationPrototypePermissionUnauthorized {
	return &UpdateOrganizationPrototypePermissionUnauthorized{}
}

/*UpdateOrganizationPrototypePermissionUnauthorized handles this case with default header values.

Session required
*/
type UpdateOrganizationPrototypePermissionUnauthorized struct {
}

func (o *UpdateOrganizationPrototypePermissionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] updateOrganizationPrototypePermissionUnauthorized ", 401)
}

func (o *UpdateOrganizationPrototypePermissionUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationPrototypePermissionForbidden creates a UpdateOrganizationPrototypePermissionForbidden with default headers values
func NewUpdateOrganizationPrototypePermissionForbidden() *UpdateOrganizationPrototypePermissionForbidden {
	return &UpdateOrganizationPrototypePermissionForbidden{}
}

/*UpdateOrganizationPrototypePermissionForbidden handles this case with default header values.

Unauthorized access
*/
type UpdateOrganizationPrototypePermissionForbidden struct {
}

func (o *UpdateOrganizationPrototypePermissionForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] updateOrganizationPrototypePermissionForbidden ", 403)
}

func (o *UpdateOrganizationPrototypePermissionForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationPrototypePermissionNotFound creates a UpdateOrganizationPrototypePermissionNotFound with default headers values
func NewUpdateOrganizationPrototypePermissionNotFound() *UpdateOrganizationPrototypePermissionNotFound {
	return &UpdateOrganizationPrototypePermissionNotFound{}
}

/*UpdateOrganizationPrototypePermissionNotFound handles this case with default header values.

Not found
*/
type UpdateOrganizationPrototypePermissionNotFound struct {
}

func (o *UpdateOrganizationPrototypePermissionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] updateOrganizationPrototypePermissionNotFound ", 404)
}

func (o *UpdateOrganizationPrototypePermissionNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
