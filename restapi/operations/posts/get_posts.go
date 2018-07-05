// Code generated by go-swagger; DO NOT EDIT.

package posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/VivaLaPanda/isle-api/models"
)

// GetPostsHandlerFunc turns a function with the right signature into a get posts handler
type GetPostsHandlerFunc func(GetPostsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPostsHandlerFunc) Handle(params GetPostsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetPostsHandler interface for that can handle valid get posts params
type GetPostsHandler interface {
	Handle(GetPostsParams, *models.Principal) middleware.Responder
}

// NewGetPosts creates a new http.Handler for the get posts operation
func NewGetPosts(ctx *middleware.Context, handler GetPostsHandler) *GetPosts {
	return &GetPosts{Context: ctx, Handler: handler}
}

/*GetPosts swagger:route GET /posts posts getPosts

Query posts

*/
type GetPosts struct {
	Context *middleware.Context
	Handler GetPostsHandler
}

func (o *GetPosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPostsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}