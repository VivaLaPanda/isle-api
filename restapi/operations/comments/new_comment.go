// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/VivaLaPanda/isle-api/models"
)

// NewCommentHandlerFunc turns a function with the right signature into a new comment handler
type NewCommentHandlerFunc func(NewCommentParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn NewCommentHandlerFunc) Handle(params NewCommentParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// NewCommentHandler interface for that can handle valid new comment params
type NewCommentHandler interface {
	Handle(NewCommentParams, *models.User) middleware.Responder
}

// NewNewComment creates a new http.Handler for the new comment operation
func NewNewComment(ctx *middleware.Context, handler NewCommentHandler) *NewComment {
	return &NewComment{Context: ctx, Handler: handler}
}

/*NewComment swagger:route POST /comments comments newComment

Create a comment

*/
type NewComment struct {
	Context *middleware.Context
	Handler NewCommentHandler
}

func (o *NewComment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNewCommentParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
