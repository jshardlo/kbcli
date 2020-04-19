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

// SearchInvoicesReader is a Reader for the SearchInvoices structure.
type SearchInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchInvoicesOK()
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

// NewSearchInvoicesOK creates a SearchInvoicesOK with default headers values
func NewSearchInvoicesOK() *SearchInvoicesOK {
	return &SearchInvoicesOK{}
}

/*SearchInvoicesOK handles this case with default header values.

successful operation
*/
type SearchInvoicesOK struct {
	Payload []*kbmodel.Invoice

	HttpResponse runtime.ClientResponse
}

func (o *SearchInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/search/{searchKey}][%d] searchInvoicesOK  %+v", 200, o.Payload)
}

func (o *SearchInvoicesOK) GetPayload() []*kbmodel.Invoice {
	return o.Payload
}

func (o *SearchInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
