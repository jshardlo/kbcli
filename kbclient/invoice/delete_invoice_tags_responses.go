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
)

// DeleteInvoiceTagsReader is a Reader for the DeleteInvoiceTags structure.
type DeleteInvoiceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInvoiceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteInvoiceTagsNoContent()
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

// NewDeleteInvoiceTagsNoContent creates a DeleteInvoiceTagsNoContent with default headers values
func NewDeleteInvoiceTagsNoContent() *DeleteInvoiceTagsNoContent {
	return &DeleteInvoiceTagsNoContent{}
}

/*DeleteInvoiceTagsNoContent handles this case with default header values.

Successful operation
*/
type DeleteInvoiceTagsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteInvoiceTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/tags][%d] deleteInvoiceTagsNoContent ", 204)
}

func (o *DeleteInvoiceTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInvoiceTagsBadRequest creates a DeleteInvoiceTagsBadRequest with default headers values
func NewDeleteInvoiceTagsBadRequest() *DeleteInvoiceTagsBadRequest {
	return &DeleteInvoiceTagsBadRequest{}
}

/*DeleteInvoiceTagsBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type DeleteInvoiceTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteInvoiceTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/tags][%d] deleteInvoiceTagsBadRequest ", 400)
}

func (o *DeleteInvoiceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
