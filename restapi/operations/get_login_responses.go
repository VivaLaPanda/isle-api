// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/VivaLaPanda/isle-api/models"
)

// GetLoginCreatedCode is the HTTP code returned for type GetLoginCreated
const GetLoginCreatedCode int = 201

/*GetLoginCreated A successful login

swagger:response getLoginCreated
*/
type GetLoginCreated struct {
}

// NewGetLoginCreated creates GetLoginCreated with default headers values
func NewGetLoginCreated() *GetLoginCreated {

	return &GetLoginCreated{}
}

// WriteResponse to the client
func (o *GetLoginCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*GetLoginDefault unexpected error

swagger:response getLoginDefault
*/
type GetLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLoginDefault creates GetLoginDefault with default headers values
func NewGetLoginDefault(code int) *GetLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get login default response
func (o *GetLoginDefault) WithStatusCode(code int) *GetLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get login default response
func (o *GetLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get login default response
func (o *GetLoginDefault) WithPayload(payload *models.Error) *GetLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get login default response
func (o *GetLoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}