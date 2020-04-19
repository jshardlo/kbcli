// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// GetInvoicePaymentReader is a Reader for the GetInvoicePayment structure.
type GetInvoicePaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicePaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicePaymentOK()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewGetInvoicePaymentOK creates a GetInvoicePaymentOK with default headers values
func NewGetInvoicePaymentOK() *GetInvoicePaymentOK {
	return &GetInvoicePaymentOK{}
}

/*GetInvoicePaymentOK handles this case with default header values.

successful operation
*/
type GetInvoicePaymentOK struct {
	Payload *kbmodel.InvoicePayment

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoicePaymentOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{paymentId}][%d] getInvoicePaymentOK  %+v", 200, o.Payload)
}

func (o *GetInvoicePaymentOK) GetPayload() *kbmodel.InvoicePayment {
	return o.Payload
}

func (o *GetInvoicePaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoicePayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicePaymentBadRequest creates a GetInvoicePaymentBadRequest with default headers values
func NewGetInvoicePaymentBadRequest() *GetInvoicePaymentBadRequest {
	return &GetInvoicePaymentBadRequest{}
}

/*GetInvoicePaymentBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type GetInvoicePaymentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoicePaymentBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{paymentId}][%d] getInvoicePaymentBadRequest ", 400)
}

func (o *GetInvoicePaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInvoicePaymentNotFound creates a GetInvoicePaymentNotFound with default headers values
func NewGetInvoicePaymentNotFound() *GetInvoicePaymentNotFound {
	return &GetInvoicePaymentNotFound{}
}

/*GetInvoicePaymentNotFound handles this case with default header values.

Payment not found
*/
type GetInvoicePaymentNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoicePaymentNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{paymentId}][%d] getInvoicePaymentNotFound ", 404)
}

func (o *GetInvoicePaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
