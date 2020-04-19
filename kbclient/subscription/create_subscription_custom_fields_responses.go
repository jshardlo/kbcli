// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateSubscriptionCustomFieldsReader is a Reader for the CreateSubscriptionCustomFields structure.
type CreateSubscriptionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateSubscriptionCustomFieldsCreated()
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

// NewCreateSubscriptionCustomFieldsCreated creates a CreateSubscriptionCustomFieldsCreated with default headers values
func NewCreateSubscriptionCustomFieldsCreated() *CreateSubscriptionCustomFieldsCreated {
	return &CreateSubscriptionCustomFieldsCreated{}
}

/*CreateSubscriptionCustomFieldsCreated handles this case with default header values.

Custom field created successfully
*/
type CreateSubscriptionCustomFieldsCreated struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateSubscriptionCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] createSubscriptionCustomFieldsCreated ", 201)
}

func (o *CreateSubscriptionCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSubscriptionCustomFieldsBadRequest creates a CreateSubscriptionCustomFieldsBadRequest with default headers values
func NewCreateSubscriptionCustomFieldsBadRequest() *CreateSubscriptionCustomFieldsBadRequest {
	return &CreateSubscriptionCustomFieldsBadRequest{}
}

/*CreateSubscriptionCustomFieldsBadRequest handles this case with default header values.

Invalid subscription id supplied
*/
type CreateSubscriptionCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateSubscriptionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] createSubscriptionCustomFieldsBadRequest ", 400)
}

func (o *CreateSubscriptionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
