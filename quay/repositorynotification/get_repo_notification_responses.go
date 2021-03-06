package repositorynotification

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

// GetRepoNotificationReader is a Reader for the GetRepoNotification structure.
type GetRepoNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetRepoNotificationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepoNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRepoNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRepoNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRepoNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRepoNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepoNotificationOK creates a GetRepoNotificationOK with default headers values
func NewGetRepoNotificationOK() *GetRepoNotificationOK {
	return &GetRepoNotificationOK{}
}

/*GetRepoNotificationOK handles this case with default header values.

Successful invocation
*/
type GetRepoNotificationOK struct {
}

func (o *GetRepoNotificationOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/notification/{uuid}][%d] getRepoNotificationOK ", 200)
}

func (o *GetRepoNotificationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoNotificationBadRequest creates a GetRepoNotificationBadRequest with default headers values
func NewGetRepoNotificationBadRequest() *GetRepoNotificationBadRequest {
	return &GetRepoNotificationBadRequest{}
}

/*GetRepoNotificationBadRequest handles this case with default header values.

Bad Request
*/
type GetRepoNotificationBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetRepoNotificationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/notification/{uuid}][%d] getRepoNotificationBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepoNotificationBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoNotificationUnauthorized creates a GetRepoNotificationUnauthorized with default headers values
func NewGetRepoNotificationUnauthorized() *GetRepoNotificationUnauthorized {
	return &GetRepoNotificationUnauthorized{}
}

/*GetRepoNotificationUnauthorized handles this case with default header values.

Session required
*/
type GetRepoNotificationUnauthorized struct {
}

func (o *GetRepoNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/notification/{uuid}][%d] getRepoNotificationUnauthorized ", 401)
}

func (o *GetRepoNotificationUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoNotificationForbidden creates a GetRepoNotificationForbidden with default headers values
func NewGetRepoNotificationForbidden() *GetRepoNotificationForbidden {
	return &GetRepoNotificationForbidden{}
}

/*GetRepoNotificationForbidden handles this case with default header values.

Unauthorized access
*/
type GetRepoNotificationForbidden struct {
}

func (o *GetRepoNotificationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/notification/{uuid}][%d] getRepoNotificationForbidden ", 403)
}

func (o *GetRepoNotificationForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoNotificationNotFound creates a GetRepoNotificationNotFound with default headers values
func NewGetRepoNotificationNotFound() *GetRepoNotificationNotFound {
	return &GetRepoNotificationNotFound{}
}

/*GetRepoNotificationNotFound handles this case with default header values.

Not found
*/
type GetRepoNotificationNotFound struct {
}

func (o *GetRepoNotificationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/notification/{uuid}][%d] getRepoNotificationNotFound ", 404)
}

func (o *GetRepoNotificationNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
