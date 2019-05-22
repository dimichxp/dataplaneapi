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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceStickRuleOKCode is the HTTP code returned for type ReplaceStickRuleOK
const ReplaceStickRuleOKCode int = 200

/*ReplaceStickRuleOK Stick Rule replaced

swagger:response replaceStickRuleOK
*/
type ReplaceStickRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.StickRule `json:"body,omitempty"`
}

// NewReplaceStickRuleOK creates ReplaceStickRuleOK with default headers values
func NewReplaceStickRuleOK() *ReplaceStickRuleOK {

	return &ReplaceStickRuleOK{}
}

// WithPayload adds the payload to the replace stick rule o k response
func (o *ReplaceStickRuleOK) WithPayload(payload *models.StickRule) *ReplaceStickRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace stick rule o k response
func (o *ReplaceStickRuleOK) SetPayload(payload *models.StickRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStickRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStickRuleAcceptedCode is the HTTP code returned for type ReplaceStickRuleAccepted
const ReplaceStickRuleAcceptedCode int = 202

/*ReplaceStickRuleAccepted Configuration change accepted and reload requested

swagger:response replaceStickRuleAccepted
*/
type ReplaceStickRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.StickRule `json:"body,omitempty"`
}

// NewReplaceStickRuleAccepted creates ReplaceStickRuleAccepted with default headers values
func NewReplaceStickRuleAccepted() *ReplaceStickRuleAccepted {

	return &ReplaceStickRuleAccepted{}
}

// WithReloadID adds the reloadId to the replace stick rule accepted response
func (o *ReplaceStickRuleAccepted) WithReloadID(reloadID string) *ReplaceStickRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace stick rule accepted response
func (o *ReplaceStickRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace stick rule accepted response
func (o *ReplaceStickRuleAccepted) WithPayload(payload *models.StickRule) *ReplaceStickRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace stick rule accepted response
func (o *ReplaceStickRuleAccepted) SetPayload(payload *models.StickRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStickRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceStickRuleBadRequestCode is the HTTP code returned for type ReplaceStickRuleBadRequest
const ReplaceStickRuleBadRequestCode int = 400

/*ReplaceStickRuleBadRequest Bad request

swagger:response replaceStickRuleBadRequest
*/
type ReplaceStickRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceStickRuleBadRequest creates ReplaceStickRuleBadRequest with default headers values
func NewReplaceStickRuleBadRequest() *ReplaceStickRuleBadRequest {

	return &ReplaceStickRuleBadRequest{}
}

// WithPayload adds the payload to the replace stick rule bad request response
func (o *ReplaceStickRuleBadRequest) WithPayload(payload *models.Error) *ReplaceStickRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace stick rule bad request response
func (o *ReplaceStickRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStickRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStickRuleNotFoundCode is the HTTP code returned for type ReplaceStickRuleNotFound
const ReplaceStickRuleNotFoundCode int = 404

/*ReplaceStickRuleNotFound The specified resource was not found

swagger:response replaceStickRuleNotFound
*/
type ReplaceStickRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceStickRuleNotFound creates ReplaceStickRuleNotFound with default headers values
func NewReplaceStickRuleNotFound() *ReplaceStickRuleNotFound {

	return &ReplaceStickRuleNotFound{}
}

// WithPayload adds the payload to the replace stick rule not found response
func (o *ReplaceStickRuleNotFound) WithPayload(payload *models.Error) *ReplaceStickRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace stick rule not found response
func (o *ReplaceStickRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStickRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceStickRuleDefault General Error

swagger:response replaceStickRuleDefault
*/
type ReplaceStickRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceStickRuleDefault creates ReplaceStickRuleDefault with default headers values
func NewReplaceStickRuleDefault(code int) *ReplaceStickRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceStickRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace stick rule default response
func (o *ReplaceStickRuleDefault) WithStatusCode(code int) *ReplaceStickRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace stick rule default response
func (o *ReplaceStickRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace stick rule default response
func (o *ReplaceStickRuleDefault) WithPayload(payload *models.Error) *ReplaceStickRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace stick rule default response
func (o *ReplaceStickRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStickRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
