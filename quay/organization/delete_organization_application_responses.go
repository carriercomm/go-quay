package organization

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

// DeleteOrganizationApplicationReader is a Reader for the DeleteOrganizationApplication structure.
type DeleteOrganizationApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteOrganizationApplicationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOrganizationApplicationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOrganizationApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteOrganizationApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteOrganizationApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrganizationApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationApplicationNoContent creates a DeleteOrganizationApplicationNoContent with default headers values
func NewDeleteOrganizationApplicationNoContent() *DeleteOrganizationApplicationNoContent {
	return &DeleteOrganizationApplicationNoContent{}
}

/*DeleteOrganizationApplicationNoContent handles this case with default header values.

Deleted
*/
type DeleteOrganizationApplicationNoContent struct {
}

func (o *DeleteOrganizationApplicationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/applications/{client_id}][%d] deleteOrganizationApplicationNoContent ", 204)
}

func (o *DeleteOrganizationApplicationNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationApplicationBadRequest creates a DeleteOrganizationApplicationBadRequest with default headers values
func NewDeleteOrganizationApplicationBadRequest() *DeleteOrganizationApplicationBadRequest {
	return &DeleteOrganizationApplicationBadRequest{}
}

/*DeleteOrganizationApplicationBadRequest handles this case with default header values.

Bad Request
*/
type DeleteOrganizationApplicationBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteOrganizationApplicationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/applications/{client_id}][%d] deleteOrganizationApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrganizationApplicationBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationApplicationUnauthorized creates a DeleteOrganizationApplicationUnauthorized with default headers values
func NewDeleteOrganizationApplicationUnauthorized() *DeleteOrganizationApplicationUnauthorized {
	return &DeleteOrganizationApplicationUnauthorized{}
}

/*DeleteOrganizationApplicationUnauthorized handles this case with default header values.

Session required
*/
type DeleteOrganizationApplicationUnauthorized struct {
}

func (o *DeleteOrganizationApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/applications/{client_id}][%d] deleteOrganizationApplicationUnauthorized ", 401)
}

func (o *DeleteOrganizationApplicationUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationApplicationForbidden creates a DeleteOrganizationApplicationForbidden with default headers values
func NewDeleteOrganizationApplicationForbidden() *DeleteOrganizationApplicationForbidden {
	return &DeleteOrganizationApplicationForbidden{}
}

/*DeleteOrganizationApplicationForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteOrganizationApplicationForbidden struct {
}

func (o *DeleteOrganizationApplicationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/applications/{client_id}][%d] deleteOrganizationApplicationForbidden ", 403)
}

func (o *DeleteOrganizationApplicationForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationApplicationNotFound creates a DeleteOrganizationApplicationNotFound with default headers values
func NewDeleteOrganizationApplicationNotFound() *DeleteOrganizationApplicationNotFound {
	return &DeleteOrganizationApplicationNotFound{}
}

/*DeleteOrganizationApplicationNotFound handles this case with default header values.

Not found
*/
type DeleteOrganizationApplicationNotFound struct {
}

func (o *DeleteOrganizationApplicationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organization/{orgname}/applications/{client_id}][%d] deleteOrganizationApplicationNotFound ", 404)
}

func (o *DeleteOrganizationApplicationNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
