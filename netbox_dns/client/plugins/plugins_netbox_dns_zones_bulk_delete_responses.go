// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PluginsNetboxDNSZonesBulkDeleteReader is a Reader for the PluginsNetboxDNSZonesBulkDelete structure.
type PluginsNetboxDNSZonesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSZonesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsNetboxDNSZonesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSZonesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSZonesBulkDeleteNoContent creates a PluginsNetboxDNSZonesBulkDeleteNoContent with default headers values
func NewPluginsNetboxDNSZonesBulkDeleteNoContent() *PluginsNetboxDNSZonesBulkDeleteNoContent {
	return &PluginsNetboxDNSZonesBulkDeleteNoContent{}
}

/*
PluginsNetboxDNSZonesBulkDeleteNoContent describes a response with status code 204, with default header values.

PluginsNetboxDNSZonesBulkDeleteNoContent plugins netbox Dns zones bulk delete no content
*/
type PluginsNetboxDNSZonesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this plugins netbox Dns zones bulk delete no content response has a 2xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns zones bulk delete no content response has a 3xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns zones bulk delete no content response has a 4xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns zones bulk delete no content response has a 5xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns zones bulk delete no content response a status code equal to that given
func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the plugins netbox Dns zones bulk delete no content response
func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) Code() int {
	return 204
}

func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/zones/][%d] pluginsNetboxDnsZonesBulkDeleteNoContent ", 204)
}

func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/zones/][%d] pluginsNetboxDnsZonesBulkDeleteNoContent ", 204)
}

func (o *PluginsNetboxDNSZonesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginsNetboxDNSZonesBulkDeleteDefault creates a PluginsNetboxDNSZonesBulkDeleteDefault with default headers values
func NewPluginsNetboxDNSZonesBulkDeleteDefault(code int) *PluginsNetboxDNSZonesBulkDeleteDefault {
	return &PluginsNetboxDNSZonesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSZonesBulkDeleteDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSZonesBulkDeleteDefault plugins netbox dns zones bulk delete default
*/
type PluginsNetboxDNSZonesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns zones bulk delete default response has a 2xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns zones bulk delete default response has a 3xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns zones bulk delete default response has a 4xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns zones bulk delete default response has a 5xx status code
func (o *PluginsNetboxDNSZonesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns zones bulk delete default response a status code equal to that given
func (o *PluginsNetboxDNSZonesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns zones bulk delete default response
func (o *PluginsNetboxDNSZonesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSZonesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/zones/][%d] plugins_netbox-dns_zones_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/zones/][%d] plugins_netbox-dns_zones_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
