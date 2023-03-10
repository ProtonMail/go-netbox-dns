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

// PluginsNetboxDNSZonesListReader is a Reader for the PluginsNetboxDNSZonesList structure.
type PluginsNetboxDNSZonesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSZonesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSZonesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSZonesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSZonesListOK creates a PluginsNetboxDNSZonesListOK with default headers values
func NewPluginsNetboxDNSZonesListOK() *PluginsNetboxDNSZonesListOK {
	return &PluginsNetboxDNSZonesListOK{}
}

/*
PluginsNetboxDNSZonesListOK describes a response with status code 200, with default header values.

PluginsNetboxDNSZonesListOK plugins netbox Dns zones list o k
*/
type PluginsNetboxDNSZonesListOK struct {
	Payload *PluginsNetboxDNSZonesListOKBody
}

// IsSuccess returns true when this plugins netbox Dns zones list o k response has a 2xx status code
func (o *PluginsNetboxDNSZonesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns zones list o k response has a 3xx status code
func (o *PluginsNetboxDNSZonesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns zones list o k response has a 4xx status code
func (o *PluginsNetboxDNSZonesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns zones list o k response has a 5xx status code
func (o *PluginsNetboxDNSZonesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns zones list o k response a status code equal to that given
func (o *PluginsNetboxDNSZonesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns zones list o k response
func (o *PluginsNetboxDNSZonesListOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSZonesListOK) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/zones/][%d] pluginsNetboxDnsZonesListOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSZonesListOK) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/zones/][%d] pluginsNetboxDnsZonesListOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSZonesListOK) GetPayload() *PluginsNetboxDNSZonesListOKBody {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PluginsNetboxDNSZonesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSZonesListDefault creates a PluginsNetboxDNSZonesListDefault with default headers values
func NewPluginsNetboxDNSZonesListDefault(code int) *PluginsNetboxDNSZonesListDefault {
	return &PluginsNetboxDNSZonesListDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSZonesListDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSZonesListDefault plugins netbox dns zones list default
*/
type PluginsNetboxDNSZonesListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns zones list default response has a 2xx status code
func (o *PluginsNetboxDNSZonesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns zones list default response has a 3xx status code
func (o *PluginsNetboxDNSZonesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns zones list default response has a 4xx status code
func (o *PluginsNetboxDNSZonesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns zones list default response has a 5xx status code
func (o *PluginsNetboxDNSZonesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns zones list default response a status code equal to that given
func (o *PluginsNetboxDNSZonesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns zones list default response
func (o *PluginsNetboxDNSZonesListDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSZonesListDefault) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/zones/][%d] plugins_netbox-dns_zones_list default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesListDefault) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/zones/][%d] plugins_netbox-dns_zones_list default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSZonesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSZonesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PluginsNetboxDNSZonesListOKBody plugins netbox DNS zones list o k body
swagger:model PluginsNetboxDNSZonesListOKBody
*/
type PluginsNetboxDNSZonesListOKBody struct {

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
	Results []*models.Zone `json:"results"`
}

// Validate validates this plugins netbox DNS zones list o k body
func (o *PluginsNetboxDNSZonesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PluginsNetboxDNSZonesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("pluginsNetboxDnsZonesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNetboxDNSZonesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsNetboxDnsZonesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNetboxDNSZonesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsNetboxDnsZonesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNetboxDNSZonesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("pluginsNetboxDnsZonesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsNetboxDnsZonesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pluginsNetboxDnsZonesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins netbox DNS zones list o k body based on the context it is used
func (o *PluginsNetboxDNSZonesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsNetboxDNSZonesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsNetboxDnsZonesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pluginsNetboxDnsZonesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PluginsNetboxDNSZonesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginsNetboxDNSZonesListOKBody) UnmarshalBinary(b []byte) error {
	var res PluginsNetboxDNSZonesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
