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

// CreateTokenReader is a Reader for the CreateToken structure.
type CreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateTokenReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTokenOK creates a CreateTokenOK with default headers values
func NewCreateTokenOK() *CreateTokenOK {
	return &CreateTokenOK{}
}

/*CreateTokenOK handles this case with default header values.

Successful invocation
*/
type CreateTokenOK struct {
}

func (o *CreateTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenOK ", 200)
}

func (o *CreateTokenOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTokenBadRequest creates a CreateTokenBadRequest with default headers values
func NewCreateTokenBadRequest() *CreateTokenBadRequest {
	return &CreateTokenBadRequest{}
}

/*CreateTokenBadRequest handles this case with default header values.

Bad Request
*/
type CreateTokenBadRequest struct {
	Payload *models.GeneralError
}

func (o *CreateTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTokenBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenUnauthorized creates a CreateTokenUnauthorized with default headers values
func NewCreateTokenUnauthorized() *CreateTokenUnauthorized {
	return &CreateTokenUnauthorized{}
}

/*CreateTokenUnauthorized handles this case with default header values.

Session required
*/
type CreateTokenUnauthorized struct {
}

func (o *CreateTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenUnauthorized ", 401)
}

func (o *CreateTokenUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTokenForbidden creates a CreateTokenForbidden with default headers values
func NewCreateTokenForbidden() *CreateTokenForbidden {
	return &CreateTokenForbidden{}
}

/*CreateTokenForbidden handles this case with default header values.

Unauthorized access
*/
type CreateTokenForbidden struct {
}

func (o *CreateTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenForbidden ", 403)
}

func (o *CreateTokenForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTokenNotFound creates a CreateTokenNotFound with default headers values
func NewCreateTokenNotFound() *CreateTokenNotFound {
	return &CreateTokenNotFound{}
}

/*CreateTokenNotFound handles this case with default header values.

Not found
*/
type CreateTokenNotFound struct {
}

func (o *CreateTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenNotFound ", 404)
}

func (o *CreateTokenNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
