// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/VivaLaPanda/isle-api/models"
)

// GetRolesHandlerFunc turns a function with the right signature into a get roles handler
type GetRolesHandlerFunc func(GetRolesParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRolesHandlerFunc) Handle(params GetRolesParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetRolesHandler interface for that can handle valid get roles params
type GetRolesHandler interface {
	Handle(GetRolesParams, *models.User) middleware.Responder
}

// NewGetRoles creates a new http.Handler for the get roles operation
func NewGetRoles(ctx *middleware.Context, handler GetRolesHandler) *GetRoles {
	return &GetRoles{Context: ctx, Handler: handler}
}

/*GetRoles swagger:route GET /roles roles getRoles

Query roles

*/
type GetRoles struct {
	Context *middleware.Context
	Handler GetRolesHandler
}

func (o *GetRoles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRolesParams()

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
