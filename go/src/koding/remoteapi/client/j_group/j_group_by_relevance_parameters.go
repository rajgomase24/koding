package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJGroupByRelevanceParams creates a new JGroupByRelevanceParams object
// with the default values initialized.
func NewJGroupByRelevanceParams() *JGroupByRelevanceParams {
	var ()
	return &JGroupByRelevanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupByRelevanceParamsWithTimeout creates a new JGroupByRelevanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupByRelevanceParamsWithTimeout(timeout time.Duration) *JGroupByRelevanceParams {
	var ()
	return &JGroupByRelevanceParams{

		timeout: timeout,
	}
}

// NewJGroupByRelevanceParamsWithContext creates a new JGroupByRelevanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupByRelevanceParamsWithContext(ctx context.Context) *JGroupByRelevanceParams {
	var ()
	return &JGroupByRelevanceParams{

		Context: ctx,
	}
}

/*JGroupByRelevanceParams contains all the parameters to send to the API endpoint
for the j group by relevance operation typically these are written to a http.Request
*/
type JGroupByRelevanceParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j group by relevance params
func (o *JGroupByRelevanceParams) WithTimeout(timeout time.Duration) *JGroupByRelevanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group by relevance params
func (o *JGroupByRelevanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group by relevance params
func (o *JGroupByRelevanceParams) WithContext(ctx context.Context) *JGroupByRelevanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group by relevance params
func (o *JGroupByRelevanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group by relevance params
func (o *JGroupByRelevanceParams) WithBody(body models.DefaultSelector) *JGroupByRelevanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group by relevance params
func (o *JGroupByRelevanceParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupByRelevanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
