package repository

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

// ChangeRepoVisibilityReader is a Reader for the ChangeRepoVisibility structure.
type ChangeRepoVisibilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ChangeRepoVisibilityReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeRepoVisibilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeRepoVisibilityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChangeRepoVisibilityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewChangeRepoVisibilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeRepoVisibilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeRepoVisibilityOK creates a ChangeRepoVisibilityOK with default headers values
func NewChangeRepoVisibilityOK() *ChangeRepoVisibilityOK {
	return &ChangeRepoVisibilityOK{}
}

/*ChangeRepoVisibilityOK handles this case with default header values.

Successful invocation
*/
type ChangeRepoVisibilityOK struct {
}

func (o *ChangeRepoVisibilityOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/changevisibility][%d] changeRepoVisibilityOK ", 200)
}

func (o *ChangeRepoVisibilityOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRepoVisibilityBadRequest creates a ChangeRepoVisibilityBadRequest with default headers values
func NewChangeRepoVisibilityBadRequest() *ChangeRepoVisibilityBadRequest {
	return &ChangeRepoVisibilityBadRequest{}
}

/*ChangeRepoVisibilityBadRequest handles this case with default header values.

Bad Request
*/
type ChangeRepoVisibilityBadRequest struct {
	Payload *models.GeneralError
}

func (o *ChangeRepoVisibilityBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/changevisibility][%d] changeRepoVisibilityBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeRepoVisibilityBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeRepoVisibilityUnauthorized creates a ChangeRepoVisibilityUnauthorized with default headers values
func NewChangeRepoVisibilityUnauthorized() *ChangeRepoVisibilityUnauthorized {
	return &ChangeRepoVisibilityUnauthorized{}
}

/*ChangeRepoVisibilityUnauthorized handles this case with default header values.

Session required
*/
type ChangeRepoVisibilityUnauthorized struct {
}

func (o *ChangeRepoVisibilityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/changevisibility][%d] changeRepoVisibilityUnauthorized ", 401)
}

func (o *ChangeRepoVisibilityUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRepoVisibilityForbidden creates a ChangeRepoVisibilityForbidden with default headers values
func NewChangeRepoVisibilityForbidden() *ChangeRepoVisibilityForbidden {
	return &ChangeRepoVisibilityForbidden{}
}

/*ChangeRepoVisibilityForbidden handles this case with default header values.

Unauthorized access
*/
type ChangeRepoVisibilityForbidden struct {
}

func (o *ChangeRepoVisibilityForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/changevisibility][%d] changeRepoVisibilityForbidden ", 403)
}

func (o *ChangeRepoVisibilityForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeRepoVisibilityNotFound creates a ChangeRepoVisibilityNotFound with default headers values
func NewChangeRepoVisibilityNotFound() *ChangeRepoVisibilityNotFound {
	return &ChangeRepoVisibilityNotFound{}
}

/*ChangeRepoVisibilityNotFound handles this case with default header values.

Not found
*/
type ChangeRepoVisibilityNotFound struct {
}

func (o *ChangeRepoVisibilityNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/changevisibility][%d] changeRepoVisibilityNotFound ", 404)
}

func (o *ChangeRepoVisibilityNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
