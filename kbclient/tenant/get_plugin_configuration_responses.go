// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetPluginConfigurationReader is a Reader for the GetPluginConfiguration structure.
type GetPluginConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPluginConfigurationOK()
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

// NewGetPluginConfigurationOK creates a GetPluginConfigurationOK with default headers values
func NewGetPluginConfigurationOK() *GetPluginConfigurationOK {
	return &GetPluginConfigurationOK{}
}

/*GetPluginConfigurationOK handles this case with default header values.

successful operation
*/
type GetPluginConfigurationOK struct {
	Payload *kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
}

func (o *GetPluginConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] getPluginConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetPluginConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginConfigurationBadRequest creates a GetPluginConfigurationBadRequest with default headers values
func NewGetPluginConfigurationBadRequest() *GetPluginConfigurationBadRequest {
	return &GetPluginConfigurationBadRequest{}
}

/*GetPluginConfigurationBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type GetPluginConfigurationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPluginConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] getPluginConfigurationBadRequest ", 400)
}

func (o *GetPluginConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}