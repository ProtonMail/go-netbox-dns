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

// NewPluginsNetboxDNSViewsViewsParams creates a new PluginsNetboxDNSViewsViewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSViewsViewsParams() *PluginsNetboxDNSViewsViewsParams {
	return &PluginsNetboxDNSViewsViewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSViewsViewsParamsWithTimeout creates a new PluginsNetboxDNSViewsViewsParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSViewsViewsParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSViewsViewsParams {
	return &PluginsNetboxDNSViewsViewsParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSViewsViewsParamsWithContext creates a new PluginsNetboxDNSViewsViewsParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSViewsViewsParamsWithContext(ctx context.Context) *PluginsNetboxDNSViewsViewsParams {
	return &PluginsNetboxDNSViewsViewsParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSViewsViewsParamsWithHTTPClient creates a new PluginsNetboxDNSViewsViewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSViewsViewsParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSViewsViewsParams {
	return &PluginsNetboxDNSViewsViewsParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSViewsViewsParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns views views operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSViewsViewsParams struct {

	/* ID.

	   A unique integer value identifying this view.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns views views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSViewsViewsParams) WithDefaults() *PluginsNetboxDNSViewsViewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns views views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSViewsViewsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSViewsViewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) WithContext(ctx context.Context) *PluginsNetboxDNSViewsViewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSViewsViewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) WithID(id int64) *PluginsNetboxDNSViewsViewsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins netbox dns views views params
func (o *PluginsNetboxDNSViewsViewsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSViewsViewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
