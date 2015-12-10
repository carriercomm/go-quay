package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type ChangeRepoVisibilityReader struct {
	formats strfmt.Registry
}

func (o *ChangeRepoVisibilityReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ChangeRepoVisibilityOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ChangeRepoVisibilityBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeRepoVisibilityBadRequest", &result, response.Code())

	case 401:
		var result ChangeRepoVisibilityUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeRepoVisibilityUnauthorized", &result, response.Code())

	case 403:
		var result ChangeRepoVisibilityForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeRepoVisibilityForbidden", &result, response.Code())

	case 404:
		var result ChangeRepoVisibilityNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("changeRepoVisibilityNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*ChangeRepoVisibilityOK

Successful invocation
*/
type ChangeRepoVisibilityOK struct {
}

func (o *ChangeRepoVisibilityOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ChangeRepoVisibilityBadRequest

Bad Request
*/
type ChangeRepoVisibilityBadRequest struct {
	Payload *models.GeneralError
}

func (o *ChangeRepoVisibilityBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*ChangeRepoVisibilityUnauthorized

Session required
*/
type ChangeRepoVisibilityUnauthorized struct {
}

func (o *ChangeRepoVisibilityUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ChangeRepoVisibilityForbidden

Unauthorized access
*/
type ChangeRepoVisibilityForbidden struct {
}

func (o *ChangeRepoVisibilityForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ChangeRepoVisibilityNotFound

Not found
*/
type ChangeRepoVisibilityNotFound struct {
}

func (o *ChangeRepoVisibilityNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}