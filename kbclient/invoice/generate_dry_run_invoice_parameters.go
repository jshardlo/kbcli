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

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewGenerateDryRunInvoiceParams creates a new GenerateDryRunInvoiceParams object
// with the default values initialized.
func NewGenerateDryRunInvoiceParams() *GenerateDryRunInvoiceParams {
	var ()
	return &GenerateDryRunInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateDryRunInvoiceParamsWithTimeout creates a new GenerateDryRunInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenerateDryRunInvoiceParamsWithTimeout(timeout time.Duration) *GenerateDryRunInvoiceParams {
	var ()
	return &GenerateDryRunInvoiceParams{

		timeout: timeout,
	}
}

// NewGenerateDryRunInvoiceParamsWithContext creates a new GenerateDryRunInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenerateDryRunInvoiceParamsWithContext(ctx context.Context) *GenerateDryRunInvoiceParams {
	var ()
	return &GenerateDryRunInvoiceParams{

		Context: ctx,
	}
}

// NewGenerateDryRunInvoiceParamsWithHTTPClient creates a new GenerateDryRunInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenerateDryRunInvoiceParamsWithHTTPClient(client *http.Client) *GenerateDryRunInvoiceParams {
	var ()
	return &GenerateDryRunInvoiceParams{
		HTTPClient: client,
	}
}

/*GenerateDryRunInvoiceParams contains all the parameters to send to the API endpoint
for the generate dry run invoice operation typically these are written to a http.Request
*/
type GenerateDryRunInvoiceParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Body*/
	Body *kbmodel.InvoiceDryRun
	/*TargetDate*/
	TargetDate *strfmt.Date

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithTimeout(timeout time.Duration) *GenerateDryRunInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithContext(ctx context.Context) *GenerateDryRunInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithHTTPClient(client *http.Client) *GenerateDryRunInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithXKillbillComment(xKillbillComment *string) *GenerateDryRunInvoiceParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *GenerateDryRunInvoiceParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithXKillbillReason(xKillbillReason *string) *GenerateDryRunInvoiceParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithAccountID(accountID strfmt.UUID) *GenerateDryRunInvoiceParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithBody(body *kbmodel.InvoiceDryRun) *GenerateDryRunInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetBody(body *kbmodel.InvoiceDryRun) {
	o.Body = body
}

// WithTargetDate adds the targetDate to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) WithTargetDate(targetDate *strfmt.Date) *GenerateDryRunInvoiceParams {
	o.SetTargetDate(targetDate)
	return o
}

// SetTargetDate adds the targetDate to the generate dry run invoice params
func (o *GenerateDryRunInvoiceParams) SetTargetDate(targetDate *strfmt.Date) {
	o.TargetDate = targetDate
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateDryRunInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param accountId
	qrAccountID := o.AccountID
	qAccountID := qrAccountID.String()
	if qAccountID != "" {
		if err := r.SetQueryParam("accountId", qAccountID); err != nil {
			return err
		}
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.TargetDate != nil {

		// query param targetDate
		var qrTargetDate strfmt.Date
		if o.TargetDate != nil {
			qrTargetDate = *o.TargetDate
		}
		qTargetDate := qrTargetDate.String()
		if qTargetDate != "" {
			if err := r.SetQueryParam("targetDate", qTargetDate); err != nil {
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
