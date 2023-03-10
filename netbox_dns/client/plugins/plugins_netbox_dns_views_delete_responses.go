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

// PluginsNetboxDNSViewsDeleteReader is a Reader for the PluginsNetboxDNSViewsDelete structure.
type PluginsNetboxDNSViewsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSViewsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPluginsNetboxDNSViewsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSViewsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSViewsDeleteNoContent creates a PluginsNetboxDNSViewsDeleteNoContent with default headers values
func NewPluginsNetboxDNSViewsDeleteNoContent() *PluginsNetboxDNSViewsDeleteNoContent {
	return &PluginsNetboxDNSViewsDeleteNoContent{}
}

/*
PluginsNetboxDNSViewsDeleteNoContent describes a response with status code 204, with default header values.

PluginsNetboxDNSViewsDeleteNoContent plugins netbox Dns views delete no content
*/
type PluginsNetboxDNSViewsDeleteNoContent struct {
}

// IsSuccess returns true when this plugins netbox Dns views delete no content response has a 2xx status code
func (o *PluginsNetboxDNSViewsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns views delete no content response has a 3xx status code
func (o *PluginsNetboxDNSViewsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns views delete no content response has a 4xx status code
func (o *PluginsNetboxDNSViewsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns views delete no content response has a 5xx status code
func (o *PluginsNetboxDNSViewsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns views delete no content response a status code equal to that given
func (o *PluginsNetboxDNSViewsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the plugins netbox Dns views delete no content response
func (o *PluginsNetboxDNSViewsDeleteNoContent) Code() int {
	return 204
}

func (o *PluginsNetboxDNSViewsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/views/{id}/][%d] pluginsNetboxDnsViewsDeleteNoContent ", 204)
}

func (o *PluginsNetboxDNSViewsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/views/{id}/][%d] pluginsNetboxDnsViewsDeleteNoContent ", 204)
}

func (o *PluginsNetboxDNSViewsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginsNetboxDNSViewsDeleteDefault creates a PluginsNetboxDNSViewsDeleteDefault with default headers values
func NewPluginsNetboxDNSViewsDeleteDefault(code int) *PluginsNetboxDNSViewsDeleteDefault {
	return &PluginsNetboxDNSViewsDeleteDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSViewsDeleteDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSViewsDeleteDefault plugins netbox dns views delete default
*/
type PluginsNetboxDNSViewsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns views delete default response has a 2xx status code
func (o *PluginsNetboxDNSViewsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns views delete default response has a 3xx status code
func (o *PluginsNetboxDNSViewsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns views delete default response has a 4xx status code
func (o *PluginsNetboxDNSViewsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns views delete default response has a 5xx status code
func (o *PluginsNetboxDNSViewsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns views delete default response a status code equal to that given
func (o *PluginsNetboxDNSViewsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns views delete default response
func (o *PluginsNetboxDNSViewsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSViewsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/views/{id}/][%d] plugins_netbox-dns_views_delete default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSViewsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /plugins/netbox-dns/views/{id}/][%d] plugins_netbox-dns_views_delete default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSViewsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSViewsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
