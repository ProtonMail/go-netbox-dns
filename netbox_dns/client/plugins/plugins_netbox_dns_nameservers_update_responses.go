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

// PluginsNetboxDNSNameserversUpdateReader is a Reader for the PluginsNetboxDNSNameserversUpdate structure.
type PluginsNetboxDNSNameserversUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSNameserversUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSNameserversUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSNameserversUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSNameserversUpdateOK creates a PluginsNetboxDNSNameserversUpdateOK with default headers values
func NewPluginsNetboxDNSNameserversUpdateOK() *PluginsNetboxDNSNameserversUpdateOK {
	return &PluginsNetboxDNSNameserversUpdateOK{}
}

/*
PluginsNetboxDNSNameserversUpdateOK describes a response with status code 200, with default header values.

PluginsNetboxDNSNameserversUpdateOK plugins netbox Dns nameservers update o k
*/
type PluginsNetboxDNSNameserversUpdateOK struct {
	Payload *models.NameServer
}

// IsSuccess returns true when this plugins netbox Dns nameservers update o k response has a 2xx status code
func (o *PluginsNetboxDNSNameserversUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns nameservers update o k response has a 3xx status code
func (o *PluginsNetboxDNSNameserversUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns nameservers update o k response has a 4xx status code
func (o *PluginsNetboxDNSNameserversUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns nameservers update o k response has a 5xx status code
func (o *PluginsNetboxDNSNameserversUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns nameservers update o k response a status code equal to that given
func (o *PluginsNetboxDNSNameserversUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns nameservers update o k response
func (o *PluginsNetboxDNSNameserversUpdateOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSNameserversUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/nameservers/{id}/][%d] pluginsNetboxDnsNameserversUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSNameserversUpdateOK) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/nameservers/{id}/][%d] pluginsNetboxDnsNameserversUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSNameserversUpdateOK) GetPayload() *models.NameServer {
	return o.Payload
}

func (o *PluginsNetboxDNSNameserversUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NameServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSNameserversUpdateDefault creates a PluginsNetboxDNSNameserversUpdateDefault with default headers values
func NewPluginsNetboxDNSNameserversUpdateDefault(code int) *PluginsNetboxDNSNameserversUpdateDefault {
	return &PluginsNetboxDNSNameserversUpdateDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSNameserversUpdateDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSNameserversUpdateDefault plugins netbox dns nameservers update default
*/
type PluginsNetboxDNSNameserversUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns nameservers update default response has a 2xx status code
func (o *PluginsNetboxDNSNameserversUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns nameservers update default response has a 3xx status code
func (o *PluginsNetboxDNSNameserversUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns nameservers update default response has a 4xx status code
func (o *PluginsNetboxDNSNameserversUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns nameservers update default response has a 5xx status code
func (o *PluginsNetboxDNSNameserversUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns nameservers update default response a status code equal to that given
func (o *PluginsNetboxDNSNameserversUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns nameservers update default response
func (o *PluginsNetboxDNSNameserversUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSNameserversUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/nameservers/{id}/][%d] plugins_netbox-dns_nameservers_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSNameserversUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/nameservers/{id}/][%d] plugins_netbox-dns_nameservers_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSNameserversUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSNameserversUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}