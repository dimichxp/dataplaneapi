// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this files except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReplaceSiteHandlerFunc turns a function with the right signature into a replace site handler
type ReplaceSiteHandlerFunc func(ReplaceSiteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceSiteHandlerFunc) Handle(params ReplaceSiteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceSiteHandler interface for that can handle valid replace site params
type ReplaceSiteHandler interface {
	Handle(ReplaceSiteParams, interface{}) middleware.Responder
}

// NewReplaceSite creates a new http.Handler for the replace site operation
func NewReplaceSite(ctx *middleware.Context, handler ReplaceSiteHandler) *ReplaceSite {
	return &ReplaceSite{Context: ctx, Handler: handler}
}

/*ReplaceSite swagger:route PUT /services/haproxy/sites/{name} Sites replaceSite

Replace a site

Replaces a site configuration by it's name.

*/
type ReplaceSite struct {
	Context *middleware.Context
	Handler ReplaceSiteHandler
}

func (o *ReplaceSite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceSiteParams()

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
