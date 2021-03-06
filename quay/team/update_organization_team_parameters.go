package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewUpdateOrganizationTeamParams creates a new UpdateOrganizationTeamParams object
// with the default values initialized.
func NewUpdateOrganizationTeamParams() *UpdateOrganizationTeamParams {
	var ()
	return &UpdateOrganizationTeamParams{}
}

/*UpdateOrganizationTeamParams contains all the parameters to send to the API endpoint
for the update organization team operation typically these are written to a http.Request
*/
type UpdateOrganizationTeamParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.TeamDescription
	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Teamname
	  The name of the team

	*/
	Teamname string
}

// WithBody adds the body to the update organization team params
func (o *UpdateOrganizationTeamParams) WithBody(body *models.TeamDescription) *UpdateOrganizationTeamParams {
	o.Body = body
	return o
}

// WithOrgname adds the orgname to the update organization team params
func (o *UpdateOrganizationTeamParams) WithOrgname(orgname string) *UpdateOrganizationTeamParams {
	o.Orgname = orgname
	return o
}

// WithTeamname adds the teamname to the update organization team params
func (o *UpdateOrganizationTeamParams) WithTeamname(teamname string) *UpdateOrganizationTeamParams {
	o.Teamname = teamname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationTeamParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.TeamDescription)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param teamname
	if err := r.SetPathParam("teamname", o.Teamname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
