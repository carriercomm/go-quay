package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type DeleteStarReader struct {
	formats strfmt.Registry
}

func (o *DeleteStarReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		var result DeleteStarNoContent
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result DeleteStarBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteStarBadRequest", &result, response.Code())

	case 401:
		var result DeleteStarUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteStarUnauthorized", &result, response.Code())

	case 403:
		var result DeleteStarForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteStarForbidden", &result, response.Code())

	case 404:
		var result DeleteStarNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteStarNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*DeleteStarNoContent

Deleted
*/
type DeleteStarNoContent struct {
}

func (o *DeleteStarNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteStarBadRequest

Bad Request
*/
type DeleteStarBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteStarBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*DeleteStarUnauthorized

Session required
*/
type DeleteStarUnauthorized struct {
}

func (o *DeleteStarUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteStarForbidden

Unauthorized access
*/
type DeleteStarForbidden struct {
}

func (o *DeleteStarForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteStarNotFound

Not found
*/
type DeleteStarNotFound struct {
}

func (o *DeleteStarNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}