package repotoken

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

// ChangeTokenReader is a Reader for the ChangeToken structure.
type ChangeTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ChangeTokenReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChangeTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewChangeTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeTokenOK creates a ChangeTokenOK with default headers values
func NewChangeTokenOK() *ChangeTokenOK {
	return &ChangeTokenOK{}
}

/*ChangeTokenOK handles this case with default header values.

Successful invocation
*/
type ChangeTokenOK struct {
}

func (o *ChangeTokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/tokens/{code}][%d] changeTokenOK ", 200)
}

func (o *ChangeTokenOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeTokenBadRequest creates a ChangeTokenBadRequest with default headers values
func NewChangeTokenBadRequest() *ChangeTokenBadRequest {
	return &ChangeTokenBadRequest{}
}

/*ChangeTokenBadRequest handles this case with default header values.

Bad Request
*/
type ChangeTokenBadRequest struct {
	Payload *models.GeneralError
}

func (o *ChangeTokenBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/tokens/{code}][%d] changeTokenBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeTokenBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeTokenUnauthorized creates a ChangeTokenUnauthorized with default headers values
func NewChangeTokenUnauthorized() *ChangeTokenUnauthorized {
	return &ChangeTokenUnauthorized{}
}

/*ChangeTokenUnauthorized handles this case with default header values.

Session required
*/
type ChangeTokenUnauthorized struct {
}

func (o *ChangeTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/tokens/{code}][%d] changeTokenUnauthorized ", 401)
}

func (o *ChangeTokenUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeTokenForbidden creates a ChangeTokenForbidden with default headers values
func NewChangeTokenForbidden() *ChangeTokenForbidden {
	return &ChangeTokenForbidden{}
}

/*ChangeTokenForbidden handles this case with default header values.

Unauthorized access
*/
type ChangeTokenForbidden struct {
}

func (o *ChangeTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/tokens/{code}][%d] changeTokenForbidden ", 403)
}

func (o *ChangeTokenForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeTokenNotFound creates a ChangeTokenNotFound with default headers values
func NewChangeTokenNotFound() *ChangeTokenNotFound {
	return &ChangeTokenNotFound{}
}

/*ChangeTokenNotFound handles this case with default header values.

Not found
*/
type ChangeTokenNotFound struct {
}

func (o *ChangeTokenNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/repository/{repository}/tokens/{code}][%d] changeTokenNotFound ", 404)
}

func (o *ChangeTokenNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
