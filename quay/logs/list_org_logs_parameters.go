package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewListOrgLogsParams creates a new ListOrgLogsParams object
// with the default values initialized.
func NewListOrgLogsParams() *ListOrgLogsParams {
	var ()
	return &ListOrgLogsParams{}
}

/*ListOrgLogsParams contains all the parameters to send to the API endpoint
for the list org logs operation typically these are written to a http.Request
*/
type ListOrgLogsParams struct {

	/*Endtime
	  Latest time to which to get logs. (%m/%d/%Y %Z)

	*/
	Endtime *string
	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Page
	  The page number for the logs

	*/
	Page *int64
	/*Performer
	  Username for which to filter logs.

	*/
	Performer *string
	/*Starttime
	  Earliest time from which to get logs. (%m/%d/%Y %Z)

	*/
	Starttime *string
}

// WithEndtime adds the endtime to the list org logs params
func (o *ListOrgLogsParams) WithEndtime(endtime *string) *ListOrgLogsParams {
	o.Endtime = endtime
	return o
}

// WithOrgname adds the orgname to the list org logs params
func (o *ListOrgLogsParams) WithOrgname(orgname string) *ListOrgLogsParams {
	o.Orgname = orgname
	return o
}

// WithPage adds the page to the list org logs params
func (o *ListOrgLogsParams) WithPage(page *int64) *ListOrgLogsParams {
	o.Page = page
	return o
}

// WithPerformer adds the performer to the list org logs params
func (o *ListOrgLogsParams) WithPerformer(performer *string) *ListOrgLogsParams {
	o.Performer = performer
	return o
}

// WithStarttime adds the starttime to the list org logs params
func (o *ListOrgLogsParams) WithStarttime(starttime *string) *ListOrgLogsParams {
	o.Starttime = starttime
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrgLogsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Endtime != nil {

		// query param endtime
		var qrEndtime string
		if o.Endtime != nil {
			qrEndtime = *o.Endtime
		}
		qEndtime := qrEndtime
		if qEndtime != "" {
			if err := r.SetQueryParam("endtime", qEndtime); err != nil {
				return err
			}
		}

	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Performer != nil {

		// query param performer
		var qrPerformer string
		if o.Performer != nil {
			qrPerformer = *o.Performer
		}
		qPerformer := qrPerformer
		if qPerformer != "" {
			if err := r.SetQueryParam("performer", qPerformer); err != nil {
				return err
			}
		}

	}

	if o.Starttime != nil {

		// query param starttime
		var qrStarttime string
		if o.Starttime != nil {
			qrStarttime = *o.Starttime
		}
		qStarttime := qrStarttime
		if qStarttime != "" {
			if err := r.SetQueryParam("starttime", qStarttime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
