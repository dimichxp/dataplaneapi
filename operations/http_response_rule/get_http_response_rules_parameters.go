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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetHTTPResponseRulesParams creates a new GetHTTPResponseRulesParams object
// no default values defined in spec.
func NewGetHTTPResponseRulesParams() GetHTTPResponseRulesParams {

	return GetHTTPResponseRulesParams{}
}

// GetHTTPResponseRulesParams contains all the bound params for the get HTTP response rules operation
// typically these are obtained from a http.Request
//
// swagger:parameters getHTTPResponseRules
type GetHTTPResponseRulesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Parent name
	  Required: true
	  In: query
	*/
	ParentName string
	/*Parent type
	  Required: true
	  In: query
	*/
	ParentType string
	/*ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	  In: query
	*/
	TransactionID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetHTTPResponseRulesParams() beforehand.
func (o *GetHTTPResponseRulesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qParentName, qhkParentName, _ := qs.GetOK("parent_name")
	if err := o.bindParentName(qParentName, qhkParentName, route.Formats); err != nil {
		res = append(res, err)
	}

	qParentType, qhkParentType, _ := qs.GetOK("parent_type")
	if err := o.bindParentType(qParentType, qhkParentType, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transaction_id")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParentName binds and validates parameter ParentName from query.
func (o *GetHTTPResponseRulesParams) bindParentName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("parent_name", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("parent_name", "query", raw); err != nil {
		return err
	}

	o.ParentName = raw

	return nil
}

// bindParentType binds and validates parameter ParentType from query.
func (o *GetHTTPResponseRulesParams) bindParentType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("parent_type", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("parent_type", "query", raw); err != nil {
		return err
	}

	o.ParentType = raw

	if err := o.validateParentType(formats); err != nil {
		return err
	}

	return nil
}

// validateParentType carries on validations for parameter ParentType
func (o *GetHTTPResponseRulesParams) validateParentType(formats strfmt.Registry) error {

	if err := validate.Enum("parent_type", "query", o.ParentType, []interface{}{"frontend", "backend"}); err != nil {
		return err
	}

	return nil
}

// bindTransactionID binds and validates parameter TransactionID from query.
func (o *GetHTTPResponseRulesParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TransactionID = &raw

	return nil
}
