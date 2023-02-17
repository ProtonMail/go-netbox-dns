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

// NewPluginsNetboxDNSRecordsCreateParams creates a new PluginsNetboxDNSRecordsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSRecordsCreateParams() *PluginsNetboxDNSRecordsCreateParams {
	return &PluginsNetboxDNSRecordsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSRecordsCreateParamsWithTimeout creates a new PluginsNetboxDNSRecordsCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSRecordsCreateParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSRecordsCreateParams {
	return &PluginsNetboxDNSRecordsCreateParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSRecordsCreateParamsWithContext creates a new PluginsNetboxDNSRecordsCreateParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSRecordsCreateParamsWithContext(ctx context.Context) *PluginsNetboxDNSRecordsCreateParams {
	return &PluginsNetboxDNSRecordsCreateParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSRecordsCreateParamsWithHTTPClient creates a new PluginsNetboxDNSRecordsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSRecordsCreateParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSRecordsCreateParams {
	return &PluginsNetboxDNSRecordsCreateParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSRecordsCreateParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns records create operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSRecordsCreateParams struct {

	// Data.
	Data *models.WritableRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns records create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSRecordsCreateParams) WithDefaults() *PluginsNetboxDNSRecordsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns records create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSRecordsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSRecordsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) WithContext(ctx context.Context) *PluginsNetboxDNSRecordsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSRecordsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) WithData(data *models.WritableRecord) *PluginsNetboxDNSRecordsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins netbox dns records create params
func (o *PluginsNetboxDNSRecordsCreateParams) SetData(data *models.WritableRecord) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSRecordsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
