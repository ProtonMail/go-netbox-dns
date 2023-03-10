// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ProtonMail/go-netbox-dns/netbox_dns/models"
)

// PluginsNetboxDNSViewsPartialUpdateReader is a Reader for the PluginsNetboxDNSViewsPartialUpdate structure.
type PluginsNetboxDNSViewsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSViewsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSViewsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSViewsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSViewsPartialUpdateOK creates a PluginsNetboxDNSViewsPartialUpdateOK with default headers values
func NewPluginsNetboxDNSViewsPartialUpdateOK() *PluginsNetboxDNSViewsPartialUpdateOK {
	return &PluginsNetboxDNSViewsPartialUpdateOK{}
}

/*
PluginsNetboxDNSViewsPartialUpdateOK describes a response with status code 200, with default header values.

PluginsNetboxDNSViewsPartialUpdateOK plugins netbox Dns views partial update o k
*/
type PluginsNetboxDNSViewsPartialUpdateOK struct {
	Payload *models.View
}

// IsSuccess returns true when this plugins netbox Dns views partial update o k response has a 2xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns views partial update o k response has a 3xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns views partial update o k response has a 4xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns views partial update o k response has a 5xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns views partial update o k response a status code equal to that given
func (o *PluginsNetboxDNSViewsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns views partial update o k response
func (o *PluginsNetboxDNSViewsPartialUpdateOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSViewsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/netbox-dns/views/{id}/][%d] pluginsNetboxDnsViewsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSViewsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /plugins/netbox-dns/views/{id}/][%d] pluginsNetboxDnsViewsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSViewsPartialUpdateOK) GetPayload() *models.View {
	return o.Payload
}

func (o *PluginsNetboxDNSViewsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.View)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSViewsPartialUpdateDefault creates a PluginsNetboxDNSViewsPartialUpdateDefault with default headers values
func NewPluginsNetboxDNSViewsPartialUpdateDefault(code int) *PluginsNetboxDNSViewsPartialUpdateDefault {
	return &PluginsNetboxDNSViewsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSViewsPartialUpdateDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSViewsPartialUpdateDefault plugins netbox dns views partial update default
*/
type PluginsNetboxDNSViewsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns views partial update default response has a 2xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns views partial update default response has a 3xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns views partial update default response has a 4xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns views partial update default response has a 5xx status code
func (o *PluginsNetboxDNSViewsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns views partial update default response a status code equal to that given
func (o *PluginsNetboxDNSViewsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns views partial update default response
func (o *PluginsNetboxDNSViewsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSViewsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /plugins/netbox-dns/views/{id}/][%d] plugins_netbox-dns_views_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSViewsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /plugins/netbox-dns/views/{id}/][%d] plugins_netbox-dns_views_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSViewsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSViewsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
