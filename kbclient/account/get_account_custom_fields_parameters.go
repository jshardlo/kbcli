// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountCustomFieldsParams creates a new GetAccountCustomFieldsParams object
// with the default values initialized.
func NewGetAccountCustomFieldsParams() *GetAccountCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAccountCustomFieldsParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountCustomFieldsParamsWithTimeout creates a new GetAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountCustomFieldsParamsWithTimeout(timeout time.Duration) *GetAccountCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAccountCustomFieldsParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetAccountCustomFieldsParamsWithContext creates a new GetAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountCustomFieldsParamsWithContext(ctx context.Context) *GetAccountCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAccountCustomFieldsParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetAccountCustomFieldsParamsWithHTTPClient creates a new GetAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountCustomFieldsParamsWithHTTPClient(client *http.Client) *GetAccountCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAccountCustomFieldsParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetAccountCustomFieldsParams contains all the parameters to send to the API endpoint
for the get account custom fields operation typically these are written to a http.Request
*/
type GetAccountCustomFieldsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithTimeout(timeout time.Duration) *GetAccountCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithContext(ctx context.Context) *GetAccountCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithHTTPClient(client *http.Client) *GetAccountCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetAccountCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetAccountCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithAccountID(accountID strfmt.UUID) *GetAccountCustomFieldsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithAudit(audit *string) *GetAccountCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

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