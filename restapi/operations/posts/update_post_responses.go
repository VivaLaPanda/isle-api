// Code generated by go-swagger; DO NOT EDIT.

package posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/VivaLaPanda/isle-api/models"
)

// UpdatePostCreatedCode is the HTTP code returned for type UpdatePostCreated
const UpdatePostCreatedCode int = 201

/*UpdatePostCreated Null response

swagger:response updatePostCreated
*/
type UpdatePostCreated struct {
}

// NewUpdatePostCreated creates UpdatePostCreated with default headers values
func NewUpdatePostCreated() *UpdatePostCreated {

	return &UpdatePostCreated{}
}

// WriteResponse to the client
func (o *UpdatePostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*UpdatePostDefault unexpected error

swagger:response updatePostDefault
*/
type UpdatePostDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePostDefault creates UpdatePostDefault with default headers values
func NewUpdatePostDefault(code int) *UpdatePostDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdatePostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update post default response
func (o *UpdatePostDefault) WithStatusCode(code int) *UpdatePostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update post default response
func (o *UpdatePostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update post default response
func (o *UpdatePostDefault) WithPayload(payload *models.Error) *UpdatePostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update post default response
func (o *UpdatePostDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
