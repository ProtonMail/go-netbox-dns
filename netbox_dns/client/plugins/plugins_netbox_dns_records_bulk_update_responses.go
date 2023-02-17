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

// PluginsNetboxDNSRecordsBulkUpdateReader is a Reader for the PluginsNetboxDNSRecordsBulkUpdate structure.
type PluginsNetboxDNSRecordsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSRecordsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSRecordsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSRecordsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSRecordsBulkUpdateOK creates a PluginsNetboxDNSRecordsBulkUpdateOK with default headers values
func NewPluginsNetboxDNSRecordsBulkUpdateOK() *PluginsNetboxDNSRecordsBulkUpdateOK {
	return &PluginsNetboxDNSRecordsBulkUpdateOK{}
}

/*
PluginsNetboxDNSRecordsBulkUpdateOK describes a response with status code 200, with default header values.

PluginsNetboxDNSRecordsBulkUpdateOK plugins netbox Dns records bulk update o k
*/
type PluginsNetboxDNSRecordsBulkUpdateOK struct {
	Payload *models.Record
}

// IsSuccess returns true when this plugins netbox Dns records bulk update o k response has a 2xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns records bulk update o k response has a 3xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns records bulk update o k response has a 4xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns records bulk update o k response has a 5xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns records bulk update o k response a status code equal to that given
func (o *PluginsNetboxDNSRecordsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns records bulk update o k response
func (o *PluginsNetboxDNSRecordsBulkUpdateOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSRecordsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/records/][%d] pluginsNetboxDnsRecordsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSRecordsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/records/][%d] pluginsNetboxDnsRecordsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSRecordsBulkUpdateOK) GetPayload() *models.Record {
	return o.Payload
}

func (o *PluginsNetboxDNSRecordsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Record)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSRecordsBulkUpdateDefault creates a PluginsNetboxDNSRecordsBulkUpdateDefault with default headers values
func NewPluginsNetboxDNSRecordsBulkUpdateDefault(code int) *PluginsNetboxDNSRecordsBulkUpdateDefault {
	return &PluginsNetboxDNSRecordsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSRecordsBulkUpdateDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSRecordsBulkUpdateDefault plugins netbox dns records bulk update default
*/
type PluginsNetboxDNSRecordsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns records bulk update default response has a 2xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns records bulk update default response has a 3xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns records bulk update default response has a 4xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns records bulk update default response has a 5xx status code
func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns records bulk update default response a status code equal to that given
func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns records bulk update default response
func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/records/][%d] plugins_netbox-dns_records_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /plugins/netbox-dns/records/][%d] plugins_netbox-dns_records_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSRecordsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
