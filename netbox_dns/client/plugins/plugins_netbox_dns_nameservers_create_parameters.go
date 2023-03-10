// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ProtonMail/go-netbox-dns/netbox_dns/models"
)

// NewPluginsNetboxDNSNameserversCreateParams creates a new PluginsNetboxDNSNameserversCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSNameserversCreateParams() *PluginsNetboxDNSNameserversCreateParams {
	return &PluginsNetboxDNSNameserversCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSNameserversCreateParamsWithTimeout creates a new PluginsNetboxDNSNameserversCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSNameserversCreateParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSNameserversCreateParams {
	return &PluginsNetboxDNSNameserversCreateParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSNameserversCreateParamsWithContext creates a new PluginsNetboxDNSNameserversCreateParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSNameserversCreateParamsWithContext(ctx context.Context) *PluginsNetboxDNSNameserversCreateParams {
	return &PluginsNetboxDNSNameserversCreateParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSNameserversCreateParamsWithHTTPClient creates a new PluginsNetboxDNSNameserversCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSNameserversCreateParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSNameserversCreateParams {
	return &PluginsNetboxDNSNameserversCreateParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSNameserversCreateParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns nameservers create operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSNameserversCreateParams struct {

	// Data.
	Data *models.NameServer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns nameservers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSNameserversCreateParams) WithDefaults() *PluginsNetboxDNSNameserversCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns nameservers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSNameserversCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSNameserversCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) WithContext(ctx context.Context) *PluginsNetboxDNSNameserversCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSNameserversCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) WithData(data *models.NameServer) *PluginsNetboxDNSNameserversCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins netbox dns nameservers create params
func (o *PluginsNetboxDNSNameserversCreateParams) SetData(data *models.NameServer) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSNameserversCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
