// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/VivaLaPanda/isle-api/models"
)

// GetUserByIDOKCode is the HTTP code returned for type GetUserByIDOK
const GetUserByIDOKCode int = 200

/*GetUserByIDOK Expected response to a valid request

swagger:response getUserByIdOK
*/
type GetUserByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserByIDOK creates GetUserByIDOK with default headers values
func NewGetUserByIDOK() *GetUserByIDOK {

	return &GetUserByIDOK{}
}

// WithPayload adds the payload to the get user by Id o k response
func (o *GetUserByIDOK) WithPayload(payload *models.User) *GetUserByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id o k response
func (o *GetUserByIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUserByIDDefault unexpected error

swagger:response getUserByIdDefault
*/
type GetUserByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserByIDDefault creates GetUserByIDDefault with default headers values
func NewGetUserByIDDefault(code int) *GetUserByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user by Id default response
func (o *GetUserByIDDefault) WithStatusCode(code int) *GetUserByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user by Id default response
func (o *GetUserByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user by Id default response
func (o *GetUserByIDDefault) WithPayload(payload *models.Error) *GetUserByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id default response
func (o *GetUserByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
