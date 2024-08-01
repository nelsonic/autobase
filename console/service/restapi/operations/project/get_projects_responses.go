// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"postgesql-cluster-console/models"
)

// GetProjectsOKCode is the HTTP code returned for type GetProjectsOK
const GetProjectsOKCode int = 200

/*
GetProjectsOK OK

swagger:response getProjectsOK
*/
type GetProjectsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseProjectsList `json:"body,omitempty"`
}

// NewGetProjectsOK creates GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {

	return &GetProjectsOK{}
}

// WithPayload adds the payload to the get projects o k response
func (o *GetProjectsOK) WithPayload(payload *models.ResponseProjectsList) *GetProjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects o k response
func (o *GetProjectsOK) SetPayload(payload *models.ResponseProjectsList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectsBadRequestCode is the HTTP code returned for type GetProjectsBadRequest
const GetProjectsBadRequestCode int = 400

/*
GetProjectsBadRequest Error

swagger:response getProjectsBadRequest
*/
type GetProjectsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseError `json:"body,omitempty"`
}

// NewGetProjectsBadRequest creates GetProjectsBadRequest with default headers values
func NewGetProjectsBadRequest() *GetProjectsBadRequest {

	return &GetProjectsBadRequest{}
}

// WithPayload adds the payload to the get projects bad request response
func (o *GetProjectsBadRequest) WithPayload(payload *models.ResponseError) *GetProjectsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects bad request response
func (o *GetProjectsBadRequest) SetPayload(payload *models.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
