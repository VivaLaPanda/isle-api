// Code generated by go-swagger; DO NOT EDIT.

package posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/VivaLaPanda/isle-api/models"
)

// GetPostsOKCode is the HTTP code returned for type GetPostsOK
const GetPostsOKCode int = 200

/*GetPostsOK An paged array of posts

swagger:response getPostsOK
*/
type GetPostsOK struct {
	/*A link to the next page of responses

	 */
	XNext string `json:"x-next"`

	/*
	  In: Body
	*/
	Payload models.ContentNodes `json:"body,omitempty"`
}

// NewGetPostsOK creates GetPostsOK with default headers values
func NewGetPostsOK() *GetPostsOK {

	return &GetPostsOK{}
}

// WithXNext adds the xNext to the get posts o k response
func (o *GetPostsOK) WithXNext(xNext string) *GetPostsOK {
	o.XNext = xNext
	return o
}

// SetXNext sets the xNext to the get posts o k response
func (o *GetPostsOK) SetXNext(xNext string) {
	o.XNext = xNext
}

// WithPayload adds the payload to the get posts o k response
func (o *GetPostsOK) WithPayload(payload models.ContentNodes) *GetPostsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get posts o k response
func (o *GetPostsOK) SetPayload(payload models.ContentNodes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header x-next

	xNext := o.XNext
	if xNext != "" {
		rw.Header().Set("x-next", xNext)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.ContentNodes, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetPostsDefault unexpected error

swagger:response getPostsDefault
*/
type GetPostsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPostsDefault creates GetPostsDefault with default headers values
func NewGetPostsDefault(code int) *GetPostsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPostsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get posts default response
func (o *GetPostsDefault) WithStatusCode(code int) *GetPostsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get posts default response
func (o *GetPostsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get posts default response
func (o *GetPostsDefault) WithPayload(payload *models.Error) *GetPostsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get posts default response
func (o *GetPostsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
