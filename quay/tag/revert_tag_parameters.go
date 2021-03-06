package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewRevertTagParams creates a new RevertTagParams object
// with the default values initialized.
func NewRevertTagParams() *RevertTagParams {
	var ()
	return &RevertTagParams{}
}

/*RevertTagParams contains all the parameters to send to the API endpoint
for the revert tag operation typically these are written to a http.Request
*/
type RevertTagParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.RevertTag
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Tag
	  The name of the tag

	*/
	Tag string
}

// WithBody adds the body to the revert tag params
func (o *RevertTagParams) WithBody(body *models.RevertTag) *RevertTagParams {
	o.Body = body
	return o
}

// WithRepository adds the repository to the revert tag params
func (o *RevertTagParams) WithRepository(repository string) *RevertTagParams {
	o.Repository = repository
	return o
}

// WithTag adds the tag to the revert tag params
func (o *RevertTagParams) WithTag(tag string) *RevertTagParams {
	o.Tag = tag
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RevertTagParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.RevertTag)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
