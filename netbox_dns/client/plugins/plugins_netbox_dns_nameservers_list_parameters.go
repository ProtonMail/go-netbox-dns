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

// NewPluginsNetboxDNSNameserversListParams creates a new PluginsNetboxDNSNameserversListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSNameserversListParams() *PluginsNetboxDNSNameserversListParams {
	return &PluginsNetboxDNSNameserversListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSNameserversListParamsWithTimeout creates a new PluginsNetboxDNSNameserversListParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSNameserversListParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSNameserversListParams {
	return &PluginsNetboxDNSNameserversListParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSNameserversListParamsWithContext creates a new PluginsNetboxDNSNameserversListParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSNameserversListParamsWithContext(ctx context.Context) *PluginsNetboxDNSNameserversListParams {
	return &PluginsNetboxDNSNameserversListParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSNameserversListParamsWithHTTPClient creates a new PluginsNetboxDNSNameserversListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSNameserversListParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSNameserversListParams {
	return &PluginsNetboxDNSNameserversListParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSNameserversListParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns nameservers list operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSNameserversListParams struct {

	// Created.
	Created *string

	// CreatedGt.
	CreatedGt *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLt.
	CreatedLt *string

	// CreatedLte.
	CreatedLte *string

	// Createdn.
	Createdn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGt.
	LastUpdatedGt *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLt.
	LastUpdatedLt *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	// LastUpdatedn.
	LastUpdatedn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	/* Ordering.

	   Which field to use when ordering the results.
	*/
	Ordering *string

	// Q.
	Q *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns nameservers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSNameserversListParams) WithDefaults() *PluginsNetboxDNSNameserversListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns nameservers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSNameserversListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSNameserversListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithContext(ctx context.Context) *PluginsNetboxDNSNameserversListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSNameserversListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithCreated(created *string) *PluginsNetboxDNSNameserversListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGt adds the createdGt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithCreatedGt(createdGt *string) *PluginsNetboxDNSNameserversListParams {
	o.SetCreatedGt(createdGt)
	return o
}

// SetCreatedGt adds the createdGt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetCreatedGt(createdGt *string) {
	o.CreatedGt = createdGt
}

// WithCreatedGte adds the createdGte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithCreatedGte(createdGte *string) *PluginsNetboxDNSNameserversListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLt adds the createdLt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithCreatedLt(createdLt *string) *PluginsNetboxDNSNameserversListParams {
	o.SetCreatedLt(createdLt)
	return o
}

// SetCreatedLt adds the createdLt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetCreatedLt(createdLt *string) {
	o.CreatedLt = createdLt
}

// WithCreatedLte adds the createdLte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithCreatedLte(createdLte *string) *PluginsNetboxDNSNameserversListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithCreatedn adds the createdn to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithCreatedn(createdn *string) *PluginsNetboxDNSNameserversListParams {
	o.SetCreatedn(createdn)
	return o
}

// SetCreatedn adds the createdN to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetCreatedn(createdn *string) {
	o.Createdn = createdn
}

// WithLastUpdated adds the lastUpdated to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLastUpdated(lastUpdated *string) *PluginsNetboxDNSNameserversListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGt adds the lastUpdatedGt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLastUpdatedGt(lastUpdatedGt *string) *PluginsNetboxDNSNameserversListParams {
	o.SetLastUpdatedGt(lastUpdatedGt)
	return o
}

// SetLastUpdatedGt adds the lastUpdatedGt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLastUpdatedGt(lastUpdatedGt *string) {
	o.LastUpdatedGt = lastUpdatedGt
}

// WithLastUpdatedGte adds the lastUpdatedGte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLastUpdatedGte(lastUpdatedGte *string) *PluginsNetboxDNSNameserversListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLt adds the lastUpdatedLt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLastUpdatedLt(lastUpdatedLt *string) *PluginsNetboxDNSNameserversListParams {
	o.SetLastUpdatedLt(lastUpdatedLt)
	return o
}

// SetLastUpdatedLt adds the lastUpdatedLt to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLastUpdatedLt(lastUpdatedLt *string) {
	o.LastUpdatedLt = lastUpdatedLt
}

// WithLastUpdatedLte adds the lastUpdatedLte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLastUpdatedLte(lastUpdatedLte *string) *PluginsNetboxDNSNameserversListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLastUpdatedn adds the lastUpdatedn to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLastUpdatedn(lastUpdatedn *string) *PluginsNetboxDNSNameserversListParams {
	o.SetLastUpdatedn(lastUpdatedn)
	return o
}

// SetLastUpdatedn adds the lastUpdatedN to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLastUpdatedn(lastUpdatedn *string) {
	o.LastUpdatedn = lastUpdatedn
}

// WithLimit adds the limit to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithLimit(limit *int64) *PluginsNetboxDNSNameserversListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithName(name *string) *PluginsNetboxDNSNameserversListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameEmpty(nameEmpty *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameIc(nameIc *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameIe(nameIe *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameIew(nameIew *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameIsw(nameIsw *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNamen(namen *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameNic(nameNic *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameNie(nameNie *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameNiew(nameNiew *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithNameNisw(nameNisw *string) *PluginsNetboxDNSNameserversListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithOffset(offset *int64) *PluginsNetboxDNSNameserversListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithOrdering(ordering *string) *PluginsNetboxDNSNameserversListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithQ adds the q to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithQ(q *string) *PluginsNetboxDNSNameserversListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithTag(tag *string) *PluginsNetboxDNSNameserversListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) WithTagn(tagn *string) *PluginsNetboxDNSNameserversListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the plugins netbox dns nameservers list params
func (o *PluginsNetboxDNSNameserversListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSNameserversListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGt != nil {

		// query param created__gt
		var qrCreatedGt string

		if o.CreatedGt != nil {
			qrCreatedGt = *o.CreatedGt
		}
		qCreatedGt := qrCreatedGt
		if qCreatedGt != "" {

			if err := r.SetQueryParam("created__gt", qCreatedGt); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLt != nil {

		// query param created__lt
		var qrCreatedLt string

		if o.CreatedLt != nil {
			qrCreatedLt = *o.CreatedLt
		}
		qCreatedLt := qrCreatedLt
		if qCreatedLt != "" {

			if err := r.SetQueryParam("created__lt", qCreatedLt); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Createdn != nil {

		// query param created__n
		var qrCreatedn string

		if o.Createdn != nil {
			qrCreatedn = *o.Createdn
		}
		qCreatedn := qrCreatedn
		if qCreatedn != "" {

			if err := r.SetQueryParam("created__n", qCreatedn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGt != nil {

		// query param last_updated__gt
		var qrLastUpdatedGt string

		if o.LastUpdatedGt != nil {
			qrLastUpdatedGt = *o.LastUpdatedGt
		}
		qLastUpdatedGt := qrLastUpdatedGt
		if qLastUpdatedGt != "" {

			if err := r.SetQueryParam("last_updated__gt", qLastUpdatedGt); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLt != nil {

		// query param last_updated__lt
		var qrLastUpdatedLt string

		if o.LastUpdatedLt != nil {
			qrLastUpdatedLt = *o.LastUpdatedLt
		}
		qLastUpdatedLt := qrLastUpdatedLt
		if qLastUpdatedLt != "" {

			if err := r.SetQueryParam("last_updated__lt", qLastUpdatedLt); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedn != nil {

		// query param last_updated__n
		var qrLastUpdatedn string

		if o.LastUpdatedn != nil {
			qrLastUpdatedn = *o.LastUpdatedn
		}
		qLastUpdatedn := qrLastUpdatedn
		if qLastUpdatedn != "" {

			if err := r.SetQueryParam("last_updated__n", qLastUpdatedn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
