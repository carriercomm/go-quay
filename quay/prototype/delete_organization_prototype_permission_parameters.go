package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteOrganizationPrototypePermissionParams creates a new DeleteOrganizationPrototypePermissionParams object
// with the default values initialized.
func NewDeleteOrganizationPrototypePermissionParams() *DeleteOrganizationPrototypePermissionParams {
	var ()
	return &DeleteOrganizationPrototypePermissionParams{}
}

/*DeleteOrganizationPrototypePermissionParams contains all the parameters to send to the API endpoint
for the delete organization prototype permission operation typically these are written to a http.Request
*/
type DeleteOrganizationPrototypePermissionParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Prototypeid
	  The ID of the prototype

	*/
	Prototypeid string
}

// WithOrgname adds the orgname to the delete organization prototype permission params
func (o *DeleteOrganizationPrototypePermissionParams) WithOrgname(orgname string) *DeleteOrganizationPrototypePermissionParams {
	o.Orgname = orgname
	return o
}

// WithPrototypeid adds the prototypeid to the delete organization prototype permission params
func (o *DeleteOrganizationPrototypePermissionParams) WithPrototypeid(prototypeid string) *DeleteOrganizationPrototypePermissionParams {
	o.Prototypeid = prototypeid
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationPrototypePermissionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param prototypeid
	if err := r.SetPathParam("prototypeid", o.Prototypeid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
