// Code generated by go-swagger; DO NOT EDIT.

package cats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"rest_api/swagger/models"
)

// DeleteIDNoContentCode is the HTTP code returned for type DeleteIDNoContent
const DeleteIDNoContentCode int = 204

/*DeleteIDNoContent Cat deleted

swagger:response deleteIdNoContent
*/
type DeleteIDNoContent struct {
}

// NewDeleteIDNoContent creates DeleteIDNoContent with default headers values
func NewDeleteIDNoContent() *DeleteIDNoContent {

	return &DeleteIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteIDNotFoundCode is the HTTP code returned for type DeleteIDNotFound
const DeleteIDNotFoundCode int = 404

/*DeleteIDNotFound not found

swagger:response deleteIdNotFound
*/
type DeleteIDNotFound struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteIDNotFound creates DeleteIDNotFound with default headers values
func NewDeleteIDNotFound() *DeleteIDNotFound {

	return &DeleteIDNotFound{}
}

// WithPayload adds the payload to the delete Id not found response
func (o *DeleteIDNotFound) WithPayload(payload models.Error) *DeleteIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Id not found response
func (o *DeleteIDNotFound) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteIDInternalServerErrorCode is the HTTP code returned for type DeleteIDInternalServerError
const DeleteIDInternalServerErrorCode int = 500

/*DeleteIDInternalServerError internal error

swagger:response deleteIdInternalServerError
*/
type DeleteIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteIDInternalServerError creates DeleteIDInternalServerError with default headers values
func NewDeleteIDInternalServerError() *DeleteIDInternalServerError {

	return &DeleteIDInternalServerError{}
}

// WithPayload adds the payload to the delete Id internal server error response
func (o *DeleteIDInternalServerError) WithPayload(payload models.Error) *DeleteIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Id internal server error response
func (o *DeleteIDInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
