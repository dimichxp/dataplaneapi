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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceServerSwitchingRuleOKCode is the HTTP code returned for type ReplaceServerSwitchingRuleOK
const ReplaceServerSwitchingRuleOKCode int = 200

/*ReplaceServerSwitchingRuleOK Server Switching Rule replaced

swagger:response replaceServerSwitchingRuleOK
*/
type ReplaceServerSwitchingRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServerSwitchingRule `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRuleOK creates ReplaceServerSwitchingRuleOK with default headers values
func NewReplaceServerSwitchingRuleOK() *ReplaceServerSwitchingRuleOK {

	return &ReplaceServerSwitchingRuleOK{}
}

// WithPayload adds the payload to the replace server switching rule o k response
func (o *ReplaceServerSwitchingRuleOK) WithPayload(payload *models.ServerSwitchingRule) *ReplaceServerSwitchingRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rule o k response
func (o *ReplaceServerSwitchingRuleOK) SetPayload(payload *models.ServerSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerSwitchingRuleAcceptedCode is the HTTP code returned for type ReplaceServerSwitchingRuleAccepted
const ReplaceServerSwitchingRuleAcceptedCode int = 202

/*ReplaceServerSwitchingRuleAccepted Configuration change accepted and reload requested

swagger:response replaceServerSwitchingRuleAccepted
*/
type ReplaceServerSwitchingRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.ServerSwitchingRule `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRuleAccepted creates ReplaceServerSwitchingRuleAccepted with default headers values
func NewReplaceServerSwitchingRuleAccepted() *ReplaceServerSwitchingRuleAccepted {

	return &ReplaceServerSwitchingRuleAccepted{}
}

// WithReloadID adds the reloadId to the replace server switching rule accepted response
func (o *ReplaceServerSwitchingRuleAccepted) WithReloadID(reloadID string) *ReplaceServerSwitchingRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace server switching rule accepted response
func (o *ReplaceServerSwitchingRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace server switching rule accepted response
func (o *ReplaceServerSwitchingRuleAccepted) WithPayload(payload *models.ServerSwitchingRule) *ReplaceServerSwitchingRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rule accepted response
func (o *ReplaceServerSwitchingRuleAccepted) SetPayload(payload *models.ServerSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerSwitchingRuleBadRequestCode is the HTTP code returned for type ReplaceServerSwitchingRuleBadRequest
const ReplaceServerSwitchingRuleBadRequestCode int = 400

/*ReplaceServerSwitchingRuleBadRequest Bad request

swagger:response replaceServerSwitchingRuleBadRequest
*/
type ReplaceServerSwitchingRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRuleBadRequest creates ReplaceServerSwitchingRuleBadRequest with default headers values
func NewReplaceServerSwitchingRuleBadRequest() *ReplaceServerSwitchingRuleBadRequest {

	return &ReplaceServerSwitchingRuleBadRequest{}
}

// WithPayload adds the payload to the replace server switching rule bad request response
func (o *ReplaceServerSwitchingRuleBadRequest) WithPayload(payload *models.Error) *ReplaceServerSwitchingRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rule bad request response
func (o *ReplaceServerSwitchingRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerSwitchingRuleNotFoundCode is the HTTP code returned for type ReplaceServerSwitchingRuleNotFound
const ReplaceServerSwitchingRuleNotFoundCode int = 404

/*ReplaceServerSwitchingRuleNotFound The specified resource was not found

swagger:response replaceServerSwitchingRuleNotFound
*/
type ReplaceServerSwitchingRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRuleNotFound creates ReplaceServerSwitchingRuleNotFound with default headers values
func NewReplaceServerSwitchingRuleNotFound() *ReplaceServerSwitchingRuleNotFound {

	return &ReplaceServerSwitchingRuleNotFound{}
}

// WithPayload adds the payload to the replace server switching rule not found response
func (o *ReplaceServerSwitchingRuleNotFound) WithPayload(payload *models.Error) *ReplaceServerSwitchingRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rule not found response
func (o *ReplaceServerSwitchingRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceServerSwitchingRuleDefault General Error

swagger:response replaceServerSwitchingRuleDefault
*/
type ReplaceServerSwitchingRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerSwitchingRuleDefault creates ReplaceServerSwitchingRuleDefault with default headers values
func NewReplaceServerSwitchingRuleDefault(code int) *ReplaceServerSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace server switching rule default response
func (o *ReplaceServerSwitchingRuleDefault) WithStatusCode(code int) *ReplaceServerSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace server switching rule default response
func (o *ReplaceServerSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace server switching rule default response
func (o *ReplaceServerSwitchingRuleDefault) WithPayload(payload *models.Error) *ReplaceServerSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server switching rule default response
func (o *ReplaceServerSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
