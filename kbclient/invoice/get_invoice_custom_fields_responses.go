// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoiceCustomFieldsReader is a Reader for the GetInvoiceCustomFields structure.
type GetInvoiceCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceCustomFieldsOK()
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

// NewGetInvoiceCustomFieldsOK creates a GetInvoiceCustomFieldsOK with default headers values
func NewGetInvoiceCustomFieldsOK() *GetInvoiceCustomFieldsOK {
	return &GetInvoiceCustomFieldsOK{}
}

/*GetInvoiceCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetInvoiceCustomFieldsOK struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/customFields][%d] getInvoiceCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetInvoiceCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceCustomFieldsBadRequest creates a GetInvoiceCustomFieldsBadRequest with default headers values
func NewGetInvoiceCustomFieldsBadRequest() *GetInvoiceCustomFieldsBadRequest {
	return &GetInvoiceCustomFieldsBadRequest{}
}

/*GetInvoiceCustomFieldsBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type GetInvoiceCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/customFields][%d] getInvoiceCustomFieldsBadRequest ", 400)
}

func (o *GetInvoiceCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
