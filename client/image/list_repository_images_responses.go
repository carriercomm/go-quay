package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type ListRepositoryImagesReader struct {
	formats strfmt.Registry
}

func (o *ListRepositoryImagesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ListRepositoryImagesOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ListRepositoryImagesBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepositoryImagesBadRequest", &result, response.Code())

	case 401:
		var result ListRepositoryImagesUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepositoryImagesUnauthorized", &result, response.Code())

	case 403:
		var result ListRepositoryImagesForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepositoryImagesForbidden", &result, response.Code())

	case 404:
		var result ListRepositoryImagesNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepositoryImagesNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type ListRepositoryImagesOK struct {
}

func (o *ListRepositoryImagesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type ListRepositoryImagesBadRequest struct {
}

func (o *ListRepositoryImagesBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type ListRepositoryImagesUnauthorized struct {
}

func (o *ListRepositoryImagesUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type ListRepositoryImagesForbidden struct {
}

func (o *ListRepositoryImagesForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type ListRepositoryImagesNotFound struct {
}

func (o *ListRepositoryImagesNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}