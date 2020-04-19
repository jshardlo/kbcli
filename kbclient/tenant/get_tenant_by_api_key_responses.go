// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// GetTenantByAPIKeyReader is a Reader for the GetTenantByAPIKey structure.
type GetTenantByAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantByAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTenantByAPIKeyOK()
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

// NewGetTenantByAPIKeyOK creates a GetTenantByAPIKeyOK with default headers values
func NewGetTenantByAPIKeyOK() *GetTenantByAPIKeyOK {
	return &GetTenantByAPIKeyOK{}
}

/*GetTenantByAPIKeyOK handles this case with default header values.

successful operation
*/
type GetTenantByAPIKeyOK struct {
	Payload *kbmodel.Tenant

	HttpResponse runtime.ClientResponse
}

func (o *GetTenantByAPIKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants][%d] getTenantByApiKeyOK  %+v", 200, o.Payload)
}

func (o *GetTenantByAPIKeyOK) GetPayload() *kbmodel.Tenant {
	return o.Payload
}

func (o *GetTenantByAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantByAPIKeyNotFound creates a GetTenantByAPIKeyNotFound with default headers values
func NewGetTenantByAPIKeyNotFound() *GetTenantByAPIKeyNotFound {
	return &GetTenantByAPIKeyNotFound{}
}

/*GetTenantByAPIKeyNotFound handles this case with default header values.

Tenant not found
*/
type GetTenantByAPIKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetTenantByAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants][%d] getTenantByApiKeyNotFound ", 404)
}

func (o *GetTenantByAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
