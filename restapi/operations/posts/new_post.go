// Code generated by go-swagger; DO NOT EDIT.

package posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NewPostHandlerFunc turns a function with the right signature into a new post handler
type NewPostHandlerFunc func(NewPostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn NewPostHandlerFunc) Handle(params NewPostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// NewPostHandler interface for that can handle valid new post params
type NewPostHandler interface {
	Handle(NewPostParams, interface{}) middleware.Responder
}

// NewNewPost creates a new http.Handler for the new post operation
func NewNewPost(ctx *middleware.Context, handler NewPostHandler) *NewPost {
	return &NewPost{Context: ctx, Handler: handler}
}

/*NewPost swagger:route POST /posts posts newPost

Create a post

*/
type NewPost struct {
	Context *middleware.Context
	Handler NewPostHandler
}

func (o *NewPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNewPostParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
