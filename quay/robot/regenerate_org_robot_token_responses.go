package robot

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

// RegenerateOrgRobotTokenReader is a Reader for the RegenerateOrgRobotToken structure.
type RegenerateOrgRobotTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *RegenerateOrgRobotTokenReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRegenerateOrgRobotTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRegenerateOrgRobotTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRegenerateOrgRobotTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRegenerateOrgRobotTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRegenerateOrgRobotTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegenerateOrgRobotTokenOK creates a RegenerateOrgRobotTokenOK with default headers values
func NewRegenerateOrgRobotTokenOK() *RegenerateOrgRobotTokenOK {
	return &RegenerateOrgRobotTokenOK{}
}

/*RegenerateOrgRobotTokenOK handles this case with default header values.

Successful invocation
*/
type RegenerateOrgRobotTokenOK struct {
}

func (o *RegenerateOrgRobotTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate][%d] regenerateOrgRobotTokenOK ", 200)
}

func (o *RegenerateOrgRobotTokenOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateOrgRobotTokenBadRequest creates a RegenerateOrgRobotTokenBadRequest with default headers values
func NewRegenerateOrgRobotTokenBadRequest() *RegenerateOrgRobotTokenBadRequest {
	return &RegenerateOrgRobotTokenBadRequest{}
}

/*RegenerateOrgRobotTokenBadRequest handles this case with default header values.

Bad Request
*/
type RegenerateOrgRobotTokenBadRequest struct {
	Payload *models.GeneralError
}

func (o *RegenerateOrgRobotTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate][%d] regenerateOrgRobotTokenBadRequest  %+v", 400, o.Payload)
}

func (o *RegenerateOrgRobotTokenBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateOrgRobotTokenUnauthorized creates a RegenerateOrgRobotTokenUnauthorized with default headers values
func NewRegenerateOrgRobotTokenUnauthorized() *RegenerateOrgRobotTokenUnauthorized {
	return &RegenerateOrgRobotTokenUnauthorized{}
}

/*RegenerateOrgRobotTokenUnauthorized handles this case with default header values.

Session required
*/
type RegenerateOrgRobotTokenUnauthorized struct {
}

func (o *RegenerateOrgRobotTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate][%d] regenerateOrgRobotTokenUnauthorized ", 401)
}

func (o *RegenerateOrgRobotTokenUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateOrgRobotTokenForbidden creates a RegenerateOrgRobotTokenForbidden with default headers values
func NewRegenerateOrgRobotTokenForbidden() *RegenerateOrgRobotTokenForbidden {
	return &RegenerateOrgRobotTokenForbidden{}
}

/*RegenerateOrgRobotTokenForbidden handles this case with default header values.

Unauthorized access
*/
type RegenerateOrgRobotTokenForbidden struct {
}

func (o *RegenerateOrgRobotTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate][%d] regenerateOrgRobotTokenForbidden ", 403)
}

func (o *RegenerateOrgRobotTokenForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateOrgRobotTokenNotFound creates a RegenerateOrgRobotTokenNotFound with default headers values
func NewRegenerateOrgRobotTokenNotFound() *RegenerateOrgRobotTokenNotFound {
	return &RegenerateOrgRobotTokenNotFound{}
}

/*RegenerateOrgRobotTokenNotFound handles this case with default header values.

Not found
*/
type RegenerateOrgRobotTokenNotFound struct {
}

func (o *RegenerateOrgRobotTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate][%d] regenerateOrgRobotTokenNotFound ", 404)
}

func (o *RegenerateOrgRobotTokenNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
