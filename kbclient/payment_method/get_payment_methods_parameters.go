// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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
)

// NewGetPaymentMethodsParams creates a new GetPaymentMethodsParams object
// with the default values initialized.
func NewGetPaymentMethodsParams() *GetPaymentMethodsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentMethodsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentMethodsParamsWithTimeout creates a new GetPaymentMethodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentMethodsParamsWithTimeout(timeout time.Duration) *GetPaymentMethodsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentMethodsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetPaymentMethodsParamsWithContext creates a new GetPaymentMethodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentMethodsParamsWithContext(ctx context.Context) *GetPaymentMethodsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentMethodsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetPaymentMethodsParamsWithHTTPClient creates a new GetPaymentMethodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentMethodsParamsWithHTTPClient(client *http.Client) *GetPaymentMethodsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentMethodsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*GetPaymentMethodsParams contains all the parameters to send to the API endpoint
for the get payment methods operation typically these are written to a http.Request
*/
type GetPaymentMethodsParams struct {

	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*PluginName*/
	PluginName *string
	/*PluginProperty*/
	PluginProperty []string
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get payment methods params
func (o *GetPaymentMethodsParams) WithTimeout(timeout time.Duration) *GetPaymentMethodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment methods params
func (o *GetPaymentMethodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment methods params
func (o *GetPaymentMethodsParams) WithContext(ctx context.Context) *GetPaymentMethodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment methods params
func (o *GetPaymentMethodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment methods params
func (o *GetPaymentMethodsParams) WithHTTPClient(client *http.Client) *GetPaymentMethodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment methods params
func (o *GetPaymentMethodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get payment methods params
func (o *GetPaymentMethodsParams) WithAudit(audit *string) *GetPaymentMethodsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment methods params
func (o *GetPaymentMethodsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the get payment methods params
func (o *GetPaymentMethodsParams) WithLimit(limit *int64) *GetPaymentMethodsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get payment methods params
func (o *GetPaymentMethodsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get payment methods params
func (o *GetPaymentMethodsParams) WithOffset(offset *int64) *GetPaymentMethodsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get payment methods params
func (o *GetPaymentMethodsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPluginName adds the pluginName to the get payment methods params
func (o *GetPaymentMethodsParams) WithPluginName(pluginName *string) *GetPaymentMethodsParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the get payment methods params
func (o *GetPaymentMethodsParams) SetPluginName(pluginName *string) {
	o.PluginName = pluginName
}

// WithPluginProperty adds the pluginProperty to the get payment methods params
func (o *GetPaymentMethodsParams) WithPluginProperty(pluginProperty []string) *GetPaymentMethodsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payment methods params
func (o *GetPaymentMethodsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithPluginInfo adds the withPluginInfo to the get payment methods params
func (o *GetPaymentMethodsParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentMethodsParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payment methods params
func (o *GetPaymentMethodsParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentMethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.PluginName != nil {

		// query param pluginName
		var qrPluginName string
		if o.PluginName != nil {
			qrPluginName = *o.PluginName
		}
		qPluginName := qrPluginName
		if qPluginName != "" {
			if err := r.SetQueryParam("pluginName", qPluginName); err != nil {
				return err
			}
		}

	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool
		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {
			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
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
