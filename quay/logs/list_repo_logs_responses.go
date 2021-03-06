package logs

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

// ListRepoLogsReader is a Reader for the ListRepoLogs structure.
type ListRepoLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ListRepoLogsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRepoLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListRepoLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListRepoLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListRepoLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListRepoLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRepoLogsOK creates a ListRepoLogsOK with default headers values
func NewListRepoLogsOK() *ListRepoLogsOK {
	return &ListRepoLogsOK{}
}

/*ListRepoLogsOK handles this case with default header values.

Successful invocation
*/
type ListRepoLogsOK struct {
}

func (o *ListRepoLogsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/logs][%d] listRepoLogsOK ", 200)
}

func (o *ListRepoLogsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoLogsBadRequest creates a ListRepoLogsBadRequest with default headers values
func NewListRepoLogsBadRequest() *ListRepoLogsBadRequest {
	return &ListRepoLogsBadRequest{}
}

/*ListRepoLogsBadRequest handles this case with default header values.

Bad Request
*/
type ListRepoLogsBadRequest struct {
	Payload *models.GeneralError
}

func (o *ListRepoLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/logs][%d] listRepoLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepoLogsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoLogsUnauthorized creates a ListRepoLogsUnauthorized with default headers values
func NewListRepoLogsUnauthorized() *ListRepoLogsUnauthorized {
	return &ListRepoLogsUnauthorized{}
}

/*ListRepoLogsUnauthorized handles this case with default header values.

Session required
*/
type ListRepoLogsUnauthorized struct {
}

func (o *ListRepoLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/logs][%d] listRepoLogsUnauthorized ", 401)
}

func (o *ListRepoLogsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoLogsForbidden creates a ListRepoLogsForbidden with default headers values
func NewListRepoLogsForbidden() *ListRepoLogsForbidden {
	return &ListRepoLogsForbidden{}
}

/*ListRepoLogsForbidden handles this case with default header values.

Unauthorized access
*/
type ListRepoLogsForbidden struct {
}

func (o *ListRepoLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/logs][%d] listRepoLogsForbidden ", 403)
}

func (o *ListRepoLogsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoLogsNotFound creates a ListRepoLogsNotFound with default headers values
func NewListRepoLogsNotFound() *ListRepoLogsNotFound {
	return &ListRepoLogsNotFound{}
}

/*ListRepoLogsNotFound handles this case with default header values.

Not found
*/
type ListRepoLogsNotFound struct {
}

func (o *ListRepoLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/logs][%d] listRepoLogsNotFound ", 404)
}

func (o *ListRepoLogsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
