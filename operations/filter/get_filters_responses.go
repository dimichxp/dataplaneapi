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

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetFiltersOKCode is the HTTP code returned for type GetFiltersOK
const GetFiltersOKCode int = 200

/*GetFiltersOK Successful operation

swagger:response getFiltersOK
*/
type GetFiltersOK struct {

	/*
	  In: Body
	*/
	Payload *GetFiltersOKBody `json:"body,omitempty"`
}

// NewGetFiltersOK creates GetFiltersOK with default headers values
func NewGetFiltersOK() *GetFiltersOK {

	return &GetFiltersOK{}
}

// WithPayload adds the payload to the get filters o k response
func (o *GetFiltersOK) WithPayload(payload *GetFiltersOKBody) *GetFiltersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filters o k response
func (o *GetFiltersOK) SetPayload(payload *GetFiltersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFiltersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetFiltersDefault General Error

swagger:response getFiltersDefault
*/
type GetFiltersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFiltersDefault creates GetFiltersDefault with default headers values
func NewGetFiltersDefault(code int) *GetFiltersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFiltersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get filters default response
func (o *GetFiltersDefault) WithStatusCode(code int) *GetFiltersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get filters default response
func (o *GetFiltersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get filters default response
func (o *GetFiltersDefault) WithPayload(payload *models.Error) *GetFiltersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filters default response
func (o *GetFiltersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFiltersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
