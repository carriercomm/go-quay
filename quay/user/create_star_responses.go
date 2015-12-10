package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type CreateStarReader struct {
	formats strfmt.Registry
}

func (o *CreateStarReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result CreateStarOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result CreateStarBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("createStarBadRequest", &result, response.Code())

	case 401:
		var result CreateStarUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("createStarUnauthorized", &result, response.Code())

	case 403:
		var result CreateStarForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("createStarForbidden", &result, response.Code())

	case 404:
		var result CreateStarNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("createStarNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*CreateStarOK

Successful invocation
*/
type CreateStarOK struct {
}

func (o *CreateStarOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateStarBadRequest

Bad Request
*/
type CreateStarBadRequest struct {
	Payload *models.GeneralError
}

func (o *CreateStarBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*CreateStarUnauthorized

Session required
*/
type CreateStarUnauthorized struct {
}

func (o *CreateStarUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateStarForbidden

Unauthorized access
*/
type CreateStarForbidden struct {
}

func (o *CreateStarForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateStarNotFound

Not found
*/
type CreateStarNotFound struct {
}

func (o *CreateStarNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}