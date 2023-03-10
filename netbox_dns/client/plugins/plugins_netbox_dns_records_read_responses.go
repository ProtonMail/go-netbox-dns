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

// PluginsNetboxDNSRecordsReadReader is a Reader for the PluginsNetboxDNSRecordsRead structure.
type PluginsNetboxDNSRecordsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSRecordsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSRecordsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSRecordsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSRecordsReadOK creates a PluginsNetboxDNSRecordsReadOK with default headers values
func NewPluginsNetboxDNSRecordsReadOK() *PluginsNetboxDNSRecordsReadOK {
	return &PluginsNetboxDNSRecordsReadOK{}
}

/*
PluginsNetboxDNSRecordsReadOK describes a response with status code 200, with default header values.

PluginsNetboxDNSRecordsReadOK plugins netbox Dns records read o k
*/
type PluginsNetboxDNSRecordsReadOK struct {
	Payload *models.Record
}

// IsSuccess returns true when this plugins netbox Dns records read o k response has a 2xx status code
func (o *PluginsNetboxDNSRecordsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns records read o k response has a 3xx status code
func (o *PluginsNetboxDNSRecordsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns records read o k response has a 4xx status code
func (o *PluginsNetboxDNSRecordsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns records read o k response has a 5xx status code
func (o *PluginsNetboxDNSRecordsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns records read o k response a status code equal to that given
func (o *PluginsNetboxDNSRecordsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns records read o k response
func (o *PluginsNetboxDNSRecordsReadOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSRecordsReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/records/{id}/][%d] pluginsNetboxDnsRecordsReadOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSRecordsReadOK) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/records/{id}/][%d] pluginsNetboxDnsRecordsReadOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSRecordsReadOK) GetPayload() *models.Record {
	return o.Payload
}

func (o *PluginsNetboxDNSRecordsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Record)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSRecordsReadDefault creates a PluginsNetboxDNSRecordsReadDefault with default headers values
func NewPluginsNetboxDNSRecordsReadDefault(code int) *PluginsNetboxDNSRecordsReadDefault {
	return &PluginsNetboxDNSRecordsReadDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSRecordsReadDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSRecordsReadDefault plugins netbox dns records read default
*/
type PluginsNetboxDNSRecordsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns records read default response has a 2xx status code
func (o *PluginsNetboxDNSRecordsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns records read default response has a 3xx status code
func (o *PluginsNetboxDNSRecordsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns records read default response has a 4xx status code
func (o *PluginsNetboxDNSRecordsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns records read default response has a 5xx status code
func (o *PluginsNetboxDNSRecordsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns records read default response a status code equal to that given
func (o *PluginsNetboxDNSRecordsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns records read default response
func (o *PluginsNetboxDNSRecordsReadDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSRecordsReadDefault) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/records/{id}/][%d] plugins_netbox-dns_records_read default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSRecordsReadDefault) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/records/{id}/][%d] plugins_netbox-dns_records_read default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSRecordsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSRecordsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
