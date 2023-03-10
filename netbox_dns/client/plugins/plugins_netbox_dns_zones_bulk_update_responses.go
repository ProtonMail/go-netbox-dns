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

// PluginsNetboxDNSZonesBulkUpdateReader is a Reader for the PluginsNetboxDNSZonesBulkUpdate structure.
type PluginsNetboxDNSZonesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSZonesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSZonesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSZonesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSZonesBulkUpdateOK creates a PluginsNetboxDNSZonesBulkUpdateOK with default headers values
func NewPluginsNetboxDNSZonesBulkUpdateOK() *PluginsNetboxDNSZonesBulkUpdateOK {
	return &PluginsNetboxDNSZonesBulkUpdateOK{}
}

/*
PluginsNetboxDNSZonesBulkUpdateOK describes a response with status code 200, with default header values.

PluginsNetboxDNSZonesBulkUpdateOK plugins netbox Dns zones bulk update o k
*/
type PluginsNetboxDNSZonesBulkUpdateOK struct {
	Payload *models.Zone
}

// IsSuccess returns true when this plugins netbox Dns zones bulk update o k response has a 2xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns zones bulk update o k response has a 3xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns zones bulk update o k response has a 4xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns zones bulk update o k response has a 5xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns zones bulk update o k response a status code equal to that given
func (o *PluginsNetboxDNSZonesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns zones bulk update o k response
func (o *PluginsNetboxDNSZonesBulkUpdateOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSZonesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/][%d] pluginsNetboxDnsZonesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSZonesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/][%d] pluginsNetboxDnsZonesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSZonesBulkUpdateOK) GetPayload() *models.Zone {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSZonesBulkUpdateDefault creates a PluginsNetboxDNSZonesBulkUpdateDefault with default headers values
func NewPluginsNetboxDNSZonesBulkUpdateDefault(code int) *PluginsNetboxDNSZonesBulkUpdateDefault {
	return &PluginsNetboxDNSZonesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSZonesBulkUpdateDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSZonesBulkUpdateDefault plugins netbox dns zones bulk update default
*/
type PluginsNetboxDNSZonesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns zones bulk update default response has a 2xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns zones bulk update default response has a 3xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns zones bulk update default response has a 4xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns zones bulk update default response has a 5xx status code
func (o *PluginsNetboxDNSZonesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns zones bulk update default response a status code equal to that given
func (o *PluginsNetboxDNSZonesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns zones bulk update default response
func (o *PluginsNetboxDNSZonesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSZonesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/][%d] plugins_netbox-dns_zones_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/zones/][%d] plugins_netbox-dns_zones_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
