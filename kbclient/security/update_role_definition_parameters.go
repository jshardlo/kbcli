// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewUpdateRoleDefinitionParams creates a new UpdateRoleDefinitionParams object
// with the default values initialized.
func NewUpdateRoleDefinitionParams() *UpdateRoleDefinitionParams {
	var ()
	return &UpdateRoleDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoleDefinitionParamsWithTimeout creates a new UpdateRoleDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRoleDefinitionParamsWithTimeout(timeout time.Duration) *UpdateRoleDefinitionParams {
	var ()
	return &UpdateRoleDefinitionParams{

		timeout: timeout,
	}
}

// NewUpdateRoleDefinitionParamsWithContext creates a new UpdateRoleDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRoleDefinitionParamsWithContext(ctx context.Context) *UpdateRoleDefinitionParams {
	var ()
	return &UpdateRoleDefinitionParams{

		Context: ctx,
	}
}

// NewUpdateRoleDefinitionParamsWithHTTPClient creates a new UpdateRoleDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRoleDefinitionParamsWithHTTPClient(client *http.Client) *UpdateRoleDefinitionParams {
	var ()
	return &UpdateRoleDefinitionParams{
		HTTPClient: client,
	}
}

/*UpdateRoleDefinitionParams contains all the parameters to send to the API endpoint
for the update role definition operation typically these are written to a http.Request
*/
type UpdateRoleDefinitionParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.RoleDefinition

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the update role definition params
func (o *UpdateRoleDefinitionParams) WithTimeout(timeout time.Duration) *UpdateRoleDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update role definition params
func (o *UpdateRoleDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update role definition params
func (o *UpdateRoleDefinitionParams) WithContext(ctx context.Context) *UpdateRoleDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update role definition params
func (o *UpdateRoleDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update role definition params
func (o *UpdateRoleDefinitionParams) WithHTTPClient(client *http.Client) *UpdateRoleDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update role definition params
func (o *UpdateRoleDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the update role definition params
func (o *UpdateRoleDefinitionParams) WithXKillbillComment(xKillbillComment *string) *UpdateRoleDefinitionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the update role definition params
func (o *UpdateRoleDefinitionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the update role definition params
func (o *UpdateRoleDefinitionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UpdateRoleDefinitionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the update role definition params
func (o *UpdateRoleDefinitionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the update role definition params
func (o *UpdateRoleDefinitionParams) WithXKillbillReason(xKillbillReason *string) *UpdateRoleDefinitionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the update role definition params
func (o *UpdateRoleDefinitionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the update role definition params
func (o *UpdateRoleDefinitionParams) WithBody(body *kbmodel.RoleDefinition) *UpdateRoleDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update role definition params
func (o *UpdateRoleDefinitionParams) SetBody(body *kbmodel.RoleDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoleDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
