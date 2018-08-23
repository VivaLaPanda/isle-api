// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/VivaLaPanda/isle-api/models"
)

// UpdateTagCreatedCode is the HTTP code returned for type UpdateTagCreated
const UpdateTagCreatedCode int = 201

/*UpdateTagCreated Null response

swagger:response updateTagCreated
*/
type UpdateTagCreated struct {
}

// NewUpdateTagCreated creates UpdateTagCreated with default headers values
func NewUpdateTagCreated() *UpdateTagCreated {

	return &UpdateTagCreated{}
}

// WriteResponse to the client
func (o *UpdateTagCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*UpdateTagDefault unexpected error

swagger:response updateTagDefault
*/
type UpdateTagDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateTagDefault creates UpdateTagDefault with default headers values
func NewUpdateTagDefault(code int) *UpdateTagDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateTagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update tag default response
func (o *UpdateTagDefault) WithStatusCode(code int) *UpdateTagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update tag default response
func (o *UpdateTagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update tag default response
func (o *UpdateTagDefault) WithPayload(payload *models.Error) *UpdateTagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update tag default response
func (o *UpdateTagDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}