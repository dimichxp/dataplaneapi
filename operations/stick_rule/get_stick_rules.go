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

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetStickRulesHandlerFunc turns a function with the right signature into a get stick rules handler
type GetStickRulesHandlerFunc func(GetStickRulesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStickRulesHandlerFunc) Handle(params GetStickRulesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetStickRulesHandler interface for that can handle valid get stick rules params
type GetStickRulesHandler interface {
	Handle(GetStickRulesParams, interface{}) middleware.Responder
}

// NewGetStickRules creates a new http.Handler for the get stick rules operation
func NewGetStickRules(ctx *middleware.Context, handler GetStickRulesHandler) *GetStickRules {
	return &GetStickRules{Context: ctx, Handler: handler}
}

/*GetStickRules swagger:route GET /services/haproxy/configuration/stick_rules StickRule getStickRules

Return an array of all Stick Rules

Returns all Stick Rules that are configured in specified backend.

*/
type GetStickRules struct {
	Context *middleware.Context
	Handler GetStickRulesHandler
}

func (o *GetStickRules) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStickRulesParams()

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

// GetStickRulesOKBody get stick rules o k body
// swagger:model GetStickRulesOKBody
type GetStickRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data models.StickRules `json:"data,omitempty"`
}

// Validate validates this get stick rules o k body
func (o *GetStickRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStickRulesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getStickRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStickRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStickRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetStickRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
