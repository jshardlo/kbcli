// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewAddBundleBlockingStateParams creates a new AddBundleBlockingStateParams object
// with the default values initialized.
func NewAddBundleBlockingStateParams() *AddBundleBlockingStateParams {
	var ()
	return &AddBundleBlockingStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddBundleBlockingStateParamsWithTimeout creates a new AddBundleBlockingStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddBundleBlockingStateParamsWithTimeout(timeout time.Duration) *AddBundleBlockingStateParams {
	var ()
	return &AddBundleBlockingStateParams{

		timeout: timeout,
	}
}

// NewAddBundleBlockingStateParamsWithContext creates a new AddBundleBlockingStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddBundleBlockingStateParamsWithContext(ctx context.Context) *AddBundleBlockingStateParams {
	var ()
	return &AddBundleBlockingStateParams{

		Context: ctx,
	}
}

// NewAddBundleBlockingStateParamsWithHTTPClient creates a new AddBundleBlockingStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddBundleBlockingStateParamsWithHTTPClient(client *http.Client) *AddBundleBlockingStateParams {
	var ()
	return &AddBundleBlockingStateParams{
		HTTPClient: client,
	}
}

/*AddBundleBlockingStateParams contains all the parameters to send to the API endpoint
for the add bundle blocking state operation typically these are written to a http.Request
*/
type AddBundleBlockingStateParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.BlockingState
	/*BundleID*/
	BundleID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string
	/*RequestedDate*/
	RequestedDate *strfmt.Date

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithTimeout(timeout time.Duration) *AddBundleBlockingStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithContext(ctx context.Context) *AddBundleBlockingStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithHTTPClient(client *http.Client) *AddBundleBlockingStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithXKillbillComment(xKillbillComment *string) *AddBundleBlockingStateParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddBundleBlockingStateParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithXKillbillReason(xKillbillReason *string) *AddBundleBlockingStateParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithBody(body *kbmodel.BlockingState) *AddBundleBlockingStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetBody(body *kbmodel.BlockingState) {
	o.Body = body
}

// WithBundleID adds the bundleID to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithBundleID(bundleID strfmt.UUID) *AddBundleBlockingStateParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithPluginProperty adds the pluginProperty to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithPluginProperty(pluginProperty []string) *AddBundleBlockingStateParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) WithRequestedDate(requestedDate *strfmt.Date) *AddBundleBlockingStateParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the add bundle blocking state params
func (o *AddBundleBlockingStateParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *AddBundleBlockingStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date
		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {
			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
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
