// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/SpenserCai/sd-webui-discord/api/gen/models"
)

// SetUserPrivateOKCode is the HTTP code returned for type SetUserPrivateOK
const SetUserPrivateOKCode int = 200

/*
SetUserPrivateOK Success

swagger:response setUserPrivateOK
*/
type SetUserPrivateOK struct {

	/*
	  In: Body
	*/
	Payload *models.BaseResponse `json:"body,omitempty"`
}

// NewSetUserPrivateOK creates SetUserPrivateOK with default headers values
func NewSetUserPrivateOK() *SetUserPrivateOK {

	return &SetUserPrivateOK{}
}

// WithPayload adds the payload to the set user private o k response
func (o *SetUserPrivateOK) WithPayload(payload *models.BaseResponse) *SetUserPrivateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set user private o k response
func (o *SetUserPrivateOK) SetPayload(payload *models.BaseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetUserPrivateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
