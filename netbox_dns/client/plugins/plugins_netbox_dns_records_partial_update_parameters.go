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

	"github.com/ProtonMail/go-netbox-dns/netbox_dns/models"
)

// NewPluginsNetboxDNSRecordsPartialUpdateParams creates a new PluginsNetboxDNSRecordsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSRecordsPartialUpdateParams() *PluginsNetboxDNSRecordsPartialUpdateParams {
	return &PluginsNetboxDNSRecordsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSRecordsPartialUpdateParamsWithTimeout creates a new PluginsNetboxDNSRecordsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSRecordsPartialUpdateParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSRecordsPartialUpdateParams {
	return &PluginsNetboxDNSRecordsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSRecordsPartialUpdateParamsWithContext creates a new PluginsNetboxDNSRecordsPartialUpdateParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSRecordsPartialUpdateParamsWithContext(ctx context.Context) *PluginsNetboxDNSRecordsPartialUpdateParams {
	return &PluginsNetboxDNSRecordsPartialUpdateParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSRecordsPartialUpdateParamsWithHTTPClient creates a new PluginsNetboxDNSRecordsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSRecordsPartialUpdateParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSRecordsPartialUpdateParams {
	return &PluginsNetboxDNSRecordsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSRecordsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns records partial update operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSRecordsPartialUpdateParams struct {

	// Data.
	Data *models.WritableRecord

	/* ID.

	   A unique integer value identifying this record.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns records partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WithDefaults() *PluginsNetboxDNSRecordsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns records partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSRecordsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WithContext(ctx context.Context) *PluginsNetboxDNSRecordsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSRecordsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WithData(data *models.WritableRecord) *PluginsNetboxDNSRecordsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) SetData(data *models.WritableRecord) {
	o.Data = data
}

// WithID adds the id to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WithID(id int64) *PluginsNetboxDNSRecordsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins netbox dns records partial update params
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSRecordsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}