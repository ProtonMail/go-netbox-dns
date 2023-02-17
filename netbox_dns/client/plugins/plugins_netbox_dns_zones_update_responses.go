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

// PluginsNetboxDNSZonesUpdateReader is a Reader for the PluginsNetboxDNSZonesUpdate structure.
type PluginsNetboxDNSZonesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSZonesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSZonesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSZonesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSZonesUpdateOK creates a PluginsNetboxDNSZonesUpdateOK with default headers values
func NewPluginsNetboxDNSZonesUpdateOK() *PluginsNetboxDNSZonesUpdateOK {
	return &PluginsNetboxDNSZonesUpdateOK{}
}

/*
PluginsNetboxDNSZonesUpdateOK describes a response with status code 200, with default header values.

PluginsNetboxDNSZonesUpdateOK plugins netbox Dns zones update o k
*/
type PluginsNetboxDNSZonesUpdateOK struct {
	Payload *models.Zone
}

// IsSuccess returns true when this plugins netbox Dns zones update o k response has a 2xx status code
func (o *PluginsNetboxDNSZonesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns zones update o k response has a 3xx status code
func (o *PluginsNetboxDNSZonesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns zones update o k response has a 4xx status code
func (o *PluginsNetboxDNSZonesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns zones update o k response has a 5xx status code
func (o *PluginsNetboxDNSZonesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns zones update o k response a status code equal to that given
func (o *PluginsNetboxDNSZonesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns zones update o k response
func (o *PluginsNetboxDNSZonesUpdateOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSZonesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/{id}/][%d] pluginsNetboxDnsZonesUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSZonesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/{id}/][%d] pluginsNetboxDnsZonesUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSZonesUpdateOK) GetPayload() *models.Zone {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSZonesUpdateDefault creates a PluginsNetboxDNSZonesUpdateDefault with default headers values
func NewPluginsNetboxDNSZonesUpdateDefault(code int) *PluginsNetboxDNSZonesUpdateDefault {
	return &PluginsNetboxDNSZonesUpdateDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSZonesUpdateDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSZonesUpdateDefault plugins netbox dns zones update default
*/
type PluginsNetboxDNSZonesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns zones update default response has a 2xx status code
func (o *PluginsNetboxDNSZonesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns zones update default response has a 3xx status code
func (o *PluginsNetboxDNSZonesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns zones update default response has a 4xx status code
func (o *PluginsNetboxDNSZonesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns zones update default response has a 5xx status code
func (o *PluginsNetboxDNSZonesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns zones update default response a status code equal to that given
func (o *PluginsNetboxDNSZonesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns zones update default response
func (o *PluginsNetboxDNSZonesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSZonesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/{id}/][%d] plugins_netbox-dns_zones_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/{id}/][%d] plugins_netbox-dns_zones_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
