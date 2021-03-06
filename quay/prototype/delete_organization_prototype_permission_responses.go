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

// DeleteOrganizationPrototypePermissionReader is a Reader for the DeleteOrganizationPrototypePermission structure.
type DeleteOrganizationPrototypePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteOrganizationPrototypePermissionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOrganizationPrototypePermissionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOrganizationPrototypePermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteOrganizationPrototypePermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteOrganizationPrototypePermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrganizationPrototypePermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationPrototypePermissionNoContent creates a DeleteOrganizationPrototypePermissionNoContent with default headers values
func NewDeleteOrganizationPrototypePermissionNoContent() *DeleteOrganizationPrototypePermissionNoContent {
	return &DeleteOrganizationPrototypePermissionNoContent{}
}

/*DeleteOrganizationPrototypePermissionNoContent handles this case with default header values.

Deleted
*/
type DeleteOrganizationPrototypePermissionNoContent struct {
}

func (o *DeleteOrganizationPrototypePermissionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] deleteOrganizationPrototypePermissionNoContent ", 204)
}

func (o *DeleteOrganizationPrototypePermissionNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationPrototypePermissionBadRequest creates a DeleteOrganizationPrototypePermissionBadRequest with default headers values
func NewDeleteOrganizationPrototypePermissionBadRequest() *DeleteOrganizationPrototypePermissionBadRequest {
	return &DeleteOrganizationPrototypePermissionBadRequest{}
}

/*DeleteOrganizationPrototypePermissionBadRequest handles this case with default header values.

Bad Request
*/
type DeleteOrganizationPrototypePermissionBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteOrganizationPrototypePermissionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] deleteOrganizationPrototypePermissionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrganizationPrototypePermissionBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationPrototypePermissionUnauthorized creates a DeleteOrganizationPrototypePermissionUnauthorized with default headers values
func NewDeleteOrganizationPrototypePermissionUnauthorized() *DeleteOrganizationPrototypePermissionUnauthorized {
	return &DeleteOrganizationPrototypePermissionUnauthorized{}
}

/*DeleteOrganizationPrototypePermissionUnauthorized handles this case with default header values.

Session required
*/
type DeleteOrganizationPrototypePermissionUnauthorized struct {
}

func (o *DeleteOrganizationPrototypePermissionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] deleteOrganizationPrototypePermissionUnauthorized ", 401)
}

func (o *DeleteOrganizationPrototypePermissionUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationPrototypePermissionForbidden creates a DeleteOrganizationPrototypePermissionForbidden with default headers values
func NewDeleteOrganizationPrototypePermissionForbidden() *DeleteOrganizationPrototypePermissionForbidden {
	return &DeleteOrganizationPrototypePermissionForbidden{}
}

/*DeleteOrganizationPrototypePermissionForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteOrganizationPrototypePermissionForbidden struct {
}

func (o *DeleteOrganizationPrototypePermissionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] deleteOrganizationPrototypePermissionForbidden ", 403)
}

func (o *DeleteOrganizationPrototypePermissionForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationPrototypePermissionNotFound creates a DeleteOrganizationPrototypePermissionNotFound with default headers values
func NewDeleteOrganizationPrototypePermissionNotFound() *DeleteOrganizationPrototypePermissionNotFound {
	return &DeleteOrganizationPrototypePermissionNotFound{}
}

/*DeleteOrganizationPrototypePermissionNotFound handles this case with default header values.

Not found
*/
type DeleteOrganizationPrototypePermissionNotFound struct {
}

func (o *DeleteOrganizationPrototypePermissionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/prototypes/{prototypeid}][%d] deleteOrganizationPrototypePermissionNotFound ", 404)
}

func (o *DeleteOrganizationPrototypePermissionNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
