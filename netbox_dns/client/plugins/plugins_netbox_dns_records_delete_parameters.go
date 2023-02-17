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

// NewPluginsNetboxDNSRecordsDeleteParams creates a new PluginsNetboxDNSRecordsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSRecordsDeleteParams() *PluginsNetboxDNSRecordsDeleteParams {
	return &PluginsNetboxDNSRecordsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSRecordsDeleteParamsWithTimeout creates a new PluginsNetboxDNSRecordsDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSRecordsDeleteParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSRecordsDeleteParams {
	return &PluginsNetboxDNSRecordsDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSRecordsDeleteParamsWithContext creates a new PluginsNetboxDNSRecordsDeleteParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSRecordsDeleteParamsWithContext(ctx context.Context) *PluginsNetboxDNSRecordsDeleteParams {
	return &PluginsNetboxDNSRecordsDeleteParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSRecordsDeleteParamsWithHTTPClient creates a new PluginsNetboxDNSRecordsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSRecordsDeleteParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSRecordsDeleteParams {
	return &PluginsNetboxDNSRecordsDeleteParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSRecordsDeleteParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns records delete operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSRecordsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this record.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns records delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSRecordsDeleteParams) WithDefaults() *PluginsNetboxDNSRecordsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns records delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSRecordsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSRecordsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) WithContext(ctx context.Context) *PluginsNetboxDNSRecordsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSRecordsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) WithID(id int64) *PluginsNetboxDNSRecordsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins netbox dns records delete params
func (o *PluginsNetboxDNSRecordsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSRecordsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
