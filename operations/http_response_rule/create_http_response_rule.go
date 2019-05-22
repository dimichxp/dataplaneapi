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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateHTTPResponseRuleHandlerFunc turns a function with the right signature into a create HTTP response rule handler
type CreateHTTPResponseRuleHandlerFunc func(CreateHTTPResponseRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateHTTPResponseRuleHandlerFunc) Handle(params CreateHTTPResponseRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateHTTPResponseRuleHandler interface for that can handle valid create HTTP response rule params
type CreateHTTPResponseRuleHandler interface {
	Handle(CreateHTTPResponseRuleParams, interface{}) middleware.Responder
}

// NewCreateHTTPResponseRule creates a new http.Handler for the create HTTP response rule operation
func NewCreateHTTPResponseRule(ctx *middleware.Context, handler CreateHTTPResponseRuleHandler) *CreateHTTPResponseRule {
	return &CreateHTTPResponseRule{Context: ctx, Handler: handler}
}

/*CreateHTTPResponseRule swagger:route POST /services/haproxy/configuration/http_response_rules HTTPResponseRule createHttpResponseRule

Add a new HTTP Response Rule

Adds a new HTTP Response Rule of the specified type in the specified parent.

*/
type CreateHTTPResponseRule struct {
	Context *middleware.Context
	Handler CreateHTTPResponseRuleHandler
}

func (o *CreateHTTPResponseRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateHTTPResponseRuleParams()

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
