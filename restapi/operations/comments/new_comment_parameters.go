// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/VivaLaPanda/isle-api/models"
)

// NewNewCommentParams creates a new NewCommentParams object
// no default values defined in spec.
func NewNewCommentParams() NewCommentParams {

	return NewCommentParams{}
}

// NewCommentParams contains all the bound params for the new comment operation
// typically these are obtained from a http.Request
//
// swagger:parameters newComment
type NewCommentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The comment to create
	  Required: true
	  In: body
	*/
	Comment *models.ContentNode
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNewCommentParams() beforehand.
func (o *NewCommentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ContentNode
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("comment", "body"))
			} else {
				res = append(res, errors.NewParseError("comment", "body", "", err))
			}
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Comment = &body
			}
		}
	} else {
		res = append(res, errors.Required("comment", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}