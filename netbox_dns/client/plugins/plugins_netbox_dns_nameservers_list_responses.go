// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ProtonMail/go-netbox-dns/netbox_dns/models"
)

// PluginsNetboxDNSNameserversListReader is a Reader for the PluginsNetboxDNSNameserversList structure.
type PluginsNetboxDNSNameserversListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSNameserversListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSNameserversListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSNameserversListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSNameserversListOK creates a PluginsNetboxDNSNameserversListOK with default headers values
func NewPluginsNetboxDNSNameserversListOK() *PluginsNetboxDNSNameserversListOK {
	return &PluginsNetboxDNSNameserversListOK{}
}

/*
PluginsNetboxDNSNameserversListOK describes a response with status code 200, with default header values.

PluginsNetboxDNSNameserversListOK plugins netbox Dns nameservers list o k
*/
type PluginsNetboxDNSNameserversListOK struct {
	Payload *PluginsNetboxDNSNameserversListOKBody
}

// IsSuccess returns true when this plugins netbox Dns nameservers list o k response has a 2xx status code
func (o *PluginsNetboxDNSNameserversListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns nameservers list o k response has a 3xx status code
func (o *PluginsNetboxDNSNameserversListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns nameservers list o k response has a 4xx status code
func (o *PluginsNetboxDNSNameserversListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns nameservers list o k response has a 5xx status code
func (o *PluginsNetboxDNSNameserversListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns nameservers list o k response a status code equal to that given
func (o *PluginsNetboxDNSNameserversListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns nameservers list o k response
func (o *PluginsNetboxDNSNameserversListOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSNameserversListOK) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/][%d] pluginsNetboxDnsNameserversListOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSNameserversListOK) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/][%d] pluginsNetboxDnsNameserversListOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSNameserversListOK) GetPayload() *PluginsNetboxDNSNameserversListOKBody {
	return o.Payload
}

func (o *PluginsNetboxDNSNameserversListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PluginsNetboxDNSNameserversListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSNameserversListDefault creates a PluginsNetboxDNSNameserversListDefault with default headers values
func NewPluginsNetboxDNSNameserversListDefault(code int) *PluginsNetboxDNSNameserversListDefault {
	return &PluginsNetboxDNSNameserversListDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSNameserversListDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSNameserversListDefault plugins netbox dns nameservers list default
*/
type PluginsNetboxDNSNameserversListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns nameservers list default response has a 2xx status code
func (o *PluginsNetboxDNSNameserversListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns nameservers list default response has a 3xx status code
func (o *PluginsNetboxDNSNameserversListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns nameservers list default response has a 4xx status code
func (o *PluginsNetboxDNSNameserversListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns nameservers list default response has a 5xx status code
func (o *PluginsNetboxDNSNameserversListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns nameservers list default response a status code equal to that given
func (o *PluginsNetboxDNSNameserversListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns nameservers list default response
func (o *PluginsNetboxDNSNameserversListDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSNameserversListDefault) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/][%d] plugins_netbox-dns_nameservers_list default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSNameserversListDefault) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/nameservers/][%d] plugins_netbox-dns_nameservers_list default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSNameserversListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSNameserversListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PluginsNetboxDNSNameserversListOKBody plugins netbox DNS nameservers list o k body
swagger:model PluginsNetboxDNSNameserversListOKBody
*/
type PluginsNetboxDNSNameserversListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.NameServer `json:"results"`
}

// Validate validates this plugins netbox DNS nameservers list o k body
func (o *PluginsNetboxDNSNameserversListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsNetboxDNSNameserversListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("pluginsNetboxDnsNameserversListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNetboxDNSNameserversListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsNetboxDnsNameserversListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNetboxDNSNameserversListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsNetboxDnsNameserversListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNetboxDNSNameserversListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("pluginsNetboxDnsNameserversListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsNetboxDnsNameserversListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pluginsNetboxDnsNameserversListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins netbox DNS nameservers list o k body based on the context it is used
func (o *PluginsNetboxDNSNameserversListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsNetboxDNSNameserversListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsNetboxDnsNameserversListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pluginsNetboxDnsNameserversListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PluginsNetboxDNSNameserversListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginsNetboxDNSNameserversListOKBody) UnmarshalBinary(b []byte) error {
	var res PluginsNetboxDNSNameserversListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
