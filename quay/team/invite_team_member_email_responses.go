package team

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

// InviteTeamMemberEmailReader is a Reader for the InviteTeamMemberEmail structure.
type InviteTeamMemberEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *InviteTeamMemberEmailReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInviteTeamMemberEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewInviteTeamMemberEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewInviteTeamMemberEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewInviteTeamMemberEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewInviteTeamMemberEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewInviteTeamMemberEmailOK creates a InviteTeamMemberEmailOK with default headers values
func NewInviteTeamMemberEmailOK() *InviteTeamMemberEmailOK {
	return &InviteTeamMemberEmailOK{}
}

/*InviteTeamMemberEmailOK handles this case with default header values.

Successful invocation
*/
type InviteTeamMemberEmailOK struct {
}

func (o *InviteTeamMemberEmailOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/invite/{email}][%d] inviteTeamMemberEmailOK ", 200)
}

func (o *InviteTeamMemberEmailOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInviteTeamMemberEmailBadRequest creates a InviteTeamMemberEmailBadRequest with default headers values
func NewInviteTeamMemberEmailBadRequest() *InviteTeamMemberEmailBadRequest {
	return &InviteTeamMemberEmailBadRequest{}
}

/*InviteTeamMemberEmailBadRequest handles this case with default header values.

Bad Request
*/
type InviteTeamMemberEmailBadRequest struct {
	Payload *models.GeneralError
}

func (o *InviteTeamMemberEmailBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/invite/{email}][%d] inviteTeamMemberEmailBadRequest  %+v", 400, o.Payload)
}

func (o *InviteTeamMemberEmailBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInviteTeamMemberEmailUnauthorized creates a InviteTeamMemberEmailUnauthorized with default headers values
func NewInviteTeamMemberEmailUnauthorized() *InviteTeamMemberEmailUnauthorized {
	return &InviteTeamMemberEmailUnauthorized{}
}

/*InviteTeamMemberEmailUnauthorized handles this case with default header values.

Session required
*/
type InviteTeamMemberEmailUnauthorized struct {
}

func (o *InviteTeamMemberEmailUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/invite/{email}][%d] inviteTeamMemberEmailUnauthorized ", 401)
}

func (o *InviteTeamMemberEmailUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInviteTeamMemberEmailForbidden creates a InviteTeamMemberEmailForbidden with default headers values
func NewInviteTeamMemberEmailForbidden() *InviteTeamMemberEmailForbidden {
	return &InviteTeamMemberEmailForbidden{}
}

/*InviteTeamMemberEmailForbidden handles this case with default header values.

Unauthorized access
*/
type InviteTeamMemberEmailForbidden struct {
}

func (o *InviteTeamMemberEmailForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/invite/{email}][%d] inviteTeamMemberEmailForbidden ", 403)
}

func (o *InviteTeamMemberEmailForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInviteTeamMemberEmailNotFound creates a InviteTeamMemberEmailNotFound with default headers values
func NewInviteTeamMemberEmailNotFound() *InviteTeamMemberEmailNotFound {
	return &InviteTeamMemberEmailNotFound{}
}

/*InviteTeamMemberEmailNotFound handles this case with default header values.

Not found
*/
type InviteTeamMemberEmailNotFound struct {
}

func (o *InviteTeamMemberEmailNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/team/{teamname}/invite/{email}][%d] inviteTeamMemberEmailNotFound ", 404)
}

func (o *InviteTeamMemberEmailNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
