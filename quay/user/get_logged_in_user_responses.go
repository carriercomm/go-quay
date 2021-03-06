package user

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

// GetLoggedInUserReader is a Reader for the GetLoggedInUser structure.
type GetLoggedInUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetLoggedInUserReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoggedInUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLoggedInUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetLoggedInUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLoggedInUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLoggedInUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoggedInUserOK creates a GetLoggedInUserOK with default headers values
func NewGetLoggedInUserOK() *GetLoggedInUserOK {
	return &GetLoggedInUserOK{}
}

/*GetLoggedInUserOK handles this case with default header values.

Successful invocation
*/
type GetLoggedInUserOK struct {
	Payload *models.UserView
}

func (o *GetLoggedInUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/][%d] getLoggedInUserOK  %+v", 200, o.Payload)
}

func (o *GetLoggedInUserOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedInUserBadRequest creates a GetLoggedInUserBadRequest with default headers values
func NewGetLoggedInUserBadRequest() *GetLoggedInUserBadRequest {
	return &GetLoggedInUserBadRequest{}
}

/*GetLoggedInUserBadRequest handles this case with default header values.

Bad Request
*/
type GetLoggedInUserBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetLoggedInUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/][%d] getLoggedInUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetLoggedInUserBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedInUserUnauthorized creates a GetLoggedInUserUnauthorized with default headers values
func NewGetLoggedInUserUnauthorized() *GetLoggedInUserUnauthorized {
	return &GetLoggedInUserUnauthorized{}
}

/*GetLoggedInUserUnauthorized handles this case with default header values.

Session required
*/
type GetLoggedInUserUnauthorized struct {
}

func (o *GetLoggedInUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/][%d] getLoggedInUserUnauthorized ", 401)
}

func (o *GetLoggedInUserUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoggedInUserForbidden creates a GetLoggedInUserForbidden with default headers values
func NewGetLoggedInUserForbidden() *GetLoggedInUserForbidden {
	return &GetLoggedInUserForbidden{}
}

/*GetLoggedInUserForbidden handles this case with default header values.

Unauthorized access
*/
type GetLoggedInUserForbidden struct {
}

func (o *GetLoggedInUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/][%d] getLoggedInUserForbidden ", 403)
}

func (o *GetLoggedInUserForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoggedInUserNotFound creates a GetLoggedInUserNotFound with default headers values
func NewGetLoggedInUserNotFound() *GetLoggedInUserNotFound {
	return &GetLoggedInUserNotFound{}
}

/*GetLoggedInUserNotFound handles this case with default header values.

Not found
*/
type GetLoggedInUserNotFound struct {
}

func (o *GetLoggedInUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/][%d] getLoggedInUserNotFound ", 404)
}

func (o *GetLoggedInUserNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
