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
	"github.com/go-openapi/swag"
)

// NewPluginsNetboxDNSNameserversZonesParams creates a new PluginsNetboxDNSNameserversZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSNameserversZonesParams() *PluginsNetboxDNSNameserversZonesParams {
	return &PluginsNetboxDNSNameserversZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSNameserversZonesParamsWithTimeout creates a new PluginsNetboxDNSNameserversZonesParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSNameserversZonesParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSNameserversZonesParams {
	return &PluginsNetboxDNSNameserversZonesParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSNameserversZonesParamsWithContext creates a new PluginsNetboxDNSNameserversZonesParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSNameserversZonesParamsWithContext(ctx context.Context) *PluginsNetboxDNSNameserversZonesParams {
	return &PluginsNetboxDNSNameserversZonesParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSNameserversZonesParamsWithHTTPClient creates a new PluginsNetboxDNSNameserversZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSNameserversZonesParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSNameserversZonesParams {
	return &PluginsNetboxDNSNameserversZonesParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSNameserversZonesParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns nameservers zones operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSNameserversZonesParams struct {

	/* ID.

	   A unique integer value identifying this name server.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns nameservers zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSNameserversZonesParams) WithDefaults() *PluginsNetboxDNSNameserversZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns nameservers zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSNameserversZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSNameserversZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) WithContext(ctx context.Context) *PluginsNetboxDNSNameserversZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSNameserversZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) WithID(id int64) *PluginsNetboxDNSNameserversZonesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins netbox dns nameservers zones params
func (o *PluginsNetboxDNSNameserversZonesParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSNameserversZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
