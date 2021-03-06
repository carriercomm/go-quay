package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewRegenerateOrgRobotTokenParams creates a new RegenerateOrgRobotTokenParams object
// with the default values initialized.
func NewRegenerateOrgRobotTokenParams() *RegenerateOrgRobotTokenParams {
	var ()
	return &RegenerateOrgRobotTokenParams{}
}

/*RegenerateOrgRobotTokenParams contains all the parameters to send to the API endpoint
for the regenerate org robot token operation typically these are written to a http.Request
*/
type RegenerateOrgRobotTokenParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*RobotShortname
	  The short name for the robot, without any user or organization prefix

	*/
	RobotShortname string
}

// WithOrgname adds the orgname to the regenerate org robot token params
func (o *RegenerateOrgRobotTokenParams) WithOrgname(orgname string) *RegenerateOrgRobotTokenParams {
	o.Orgname = orgname
	return o
}

// WithRobotShortname adds the robotShortname to the regenerate org robot token params
func (o *RegenerateOrgRobotTokenParams) WithRobotShortname(robotShortname string) *RegenerateOrgRobotTokenParams {
	o.RobotShortname = robotShortname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RegenerateOrgRobotTokenParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param robot_shortname
	if err := r.SetPathParam("robot_shortname", o.RobotShortname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
