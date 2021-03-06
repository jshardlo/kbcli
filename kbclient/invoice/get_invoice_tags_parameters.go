// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewGetInvoiceTagsParams creates a new GetInvoiceTagsParams object
// with the default values initialized.
func NewGetInvoiceTagsParams() *GetInvoiceTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceTagsParamsWithTimeout creates a new GetInvoiceTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceTagsParamsWithTimeout(timeout time.Duration) *GetInvoiceTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetInvoiceTagsParamsWithContext creates a new GetInvoiceTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceTagsParamsWithContext(ctx context.Context) *GetInvoiceTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetInvoiceTagsParamsWithHTTPClient creates a new GetInvoiceTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceTagsParamsWithHTTPClient(client *http.Client) *GetInvoiceTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetInvoiceTagsParams contains all the parameters to send to the API endpoint
for the get invoice tags operation typically these are written to a http.Request
*/
type GetInvoiceTagsParams struct {

	/*Audit*/
	Audit *string
	/*IncludedDeleted*/
	IncludedDeleted *bool
	/*InvoiceID*/
	InvoiceID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice tags params
func (o *GetInvoiceTagsParams) WithTimeout(timeout time.Duration) *GetInvoiceTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice tags params
func (o *GetInvoiceTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice tags params
func (o *GetInvoiceTagsParams) WithContext(ctx context.Context) *GetInvoiceTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice tags params
func (o *GetInvoiceTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice tags params
func (o *GetInvoiceTagsParams) WithHTTPClient(client *http.Client) *GetInvoiceTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice tags params
func (o *GetInvoiceTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get invoice tags params
func (o *GetInvoiceTagsParams) WithAudit(audit *string) *GetInvoiceTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice tags params
func (o *GetInvoiceTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get invoice tags params
func (o *GetInvoiceTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetInvoiceTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get invoice tags params
func (o *GetInvoiceTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithInvoiceID adds the invoiceID to the get invoice tags params
func (o *GetInvoiceTagsParams) WithInvoiceID(invoiceID strfmt.UUID) *GetInvoiceTagsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the get invoice tags params
func (o *GetInvoiceTagsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
		}

	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
		return err
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
