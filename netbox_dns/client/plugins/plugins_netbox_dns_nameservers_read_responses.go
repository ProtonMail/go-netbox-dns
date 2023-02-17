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

// PluginsNetboxDNSNameserversReadReader is a Reader for the PluginsNetboxDNSNameserversRead structure.
type PluginsNetboxDNSNameserversReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSNameserversReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSNameserversReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSNameserversReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSNameserversReadOK creates a PluginsNetboxDNSNameserversReadOK with default headers values
func NewPluginsNetboxDNSNameserversReadOK() *PluginsNetboxDNSNameserversReadOK {
	return &PluginsNetboxDNSNameserversReadOK{}
}

/*
PluginsNetboxDNSNameserversReadOK describes a response with status code 200, with default header values.

PluginsNetboxDNSNameserversReadOK plugins netbox Dns nameservers read o k
*/
type PluginsNetboxDNSNameserversReadOK struct {
	Payload *models.NameServer
}

// IsSuccess returns true when this plugins netbox Dns nameservers read o k response has a 2xx status code
func (o *PluginsNetboxDNSNameserversReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns nameservers read o k response has a 3xx status code
func (o *PluginsNetboxDNSNameserversReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns nameservers read o k response has a 4xx status code
func (o *PluginsNetboxDNSNameserversReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns nameservers read o k response has a 5xx status code
func (o *PluginsNetboxDNSNameserversReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns nameservers read o k response a status code equal to that given
func (o *PluginsNetboxDNSNameserversReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns nameservers read o k response
func (o *PluginsNetboxDNSNameserversReadOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSNameserversReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/{id}/][%d] pluginsNetboxDnsNameserversReadOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSNameserversReadOK) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/{id}/][%d] pluginsNetboxDnsNameserversReadOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSNameserversReadOK) GetPayload() *models.NameServer {
	return o.Payload
}

func (o *PluginsNetboxDNSNameserversReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NameServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSNameserversReadDefault creates a PluginsNetboxDNSNameserversReadDefault with default headers values
func NewPluginsNetboxDNSNameserversReadDefault(code int) *PluginsNetboxDNSNameserversReadDefault {
	return &PluginsNetboxDNSNameserversReadDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSNameserversReadDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSNameserversReadDefault plugins netbox dns nameservers read default
*/
type PluginsNetboxDNSNameserversReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns nameservers read default response has a 2xx status code
func (o *PluginsNetboxDNSNameserversReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns nameservers read default response has a 3xx status code
func (o *PluginsNetboxDNSNameserversReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns nameservers read default response has a 4xx status code
func (o *PluginsNetboxDNSNameserversReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns nameservers read default response has a 5xx status code
func (o *PluginsNetboxDNSNameserversReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns nameservers read default response a status code equal to that given
func (o *PluginsNetboxDNSNameserversReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns nameservers read default response
func (o *PluginsNetboxDNSNameserversReadDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSNameserversReadDefault) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/{id}/][%d] plugins_netbox-dns_nameservers_read default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSNameserversReadDefault) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/{id}/][%d] plugins_netbox-dns_nameservers_read default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSNameserversReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSNameserversReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
