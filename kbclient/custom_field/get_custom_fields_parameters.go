// Code generated by go-swagger; DO NOT EDIT.

package custom_field

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

// NewGetCustomFieldsParams creates a new GetCustomFieldsParams object
// with the default values initialized.
func NewGetCustomFieldsParams() *GetCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetCustomFieldsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomFieldsParamsWithTimeout creates a new GetCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomFieldsParamsWithTimeout(timeout time.Duration) *GetCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetCustomFieldsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetCustomFieldsParamsWithContext creates a new GetCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomFieldsParamsWithContext(ctx context.Context) *GetCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetCustomFieldsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetCustomFieldsParamsWithHTTPClient creates a new GetCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomFieldsParamsWithHTTPClient(client *http.Client) *GetCustomFieldsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetCustomFieldsParams{
		Audit:      &auditDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetCustomFieldsParams contains all the parameters to send to the API endpoint
for the get custom fields operation typically these are written to a http.Request
*/
type GetCustomFieldsParams struct {

	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get custom fields params
func (o *GetCustomFieldsParams) WithTimeout(timeout time.Duration) *GetCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom fields params
func (o *GetCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom fields params
func (o *GetCustomFieldsParams) WithContext(ctx context.Context) *GetCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom fields params
func (o *GetCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom fields params
func (o *GetCustomFieldsParams) WithHTTPClient(client *http.Client) *GetCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom fields params
func (o *GetCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get custom fields params
func (o *GetCustomFieldsParams) WithAudit(audit *string) *GetCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get custom fields params
func (o *GetCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the get custom fields params
func (o *GetCustomFieldsParams) WithLimit(limit *int64) *GetCustomFieldsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get custom fields params
func (o *GetCustomFieldsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get custom fields params
func (o *GetCustomFieldsParams) WithOffset(offset *int64) *GetCustomFieldsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get custom fields params
func (o *GetCustomFieldsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
