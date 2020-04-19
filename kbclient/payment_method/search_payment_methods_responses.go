// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// SearchPaymentMethodsReader is a Reader for the SearchPaymentMethods structure.
type SearchPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchPaymentMethodsOK()
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

// NewSearchPaymentMethodsOK creates a SearchPaymentMethodsOK with default headers values
func NewSearchPaymentMethodsOK() *SearchPaymentMethodsOK {
	return &SearchPaymentMethodsOK{}
}

/*SearchPaymentMethodsOK handles this case with default header values.

successful operation
*/
type SearchPaymentMethodsOK struct {
	Payload []*kbmodel.PaymentMethod

	HttpResponse runtime.ClientResponse
}

func (o *SearchPaymentMethodsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/search/{searchKey}][%d] searchPaymentMethodsOK  %+v", 200, o.Payload)
}

func (o *SearchPaymentMethodsOK) GetPayload() []*kbmodel.PaymentMethod {
	return o.Payload
}

func (o *SearchPaymentMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
