// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/terraform-provider-nuon/internal/api/client/models"
)

// PostV1SandboxesReader is a Reader for the PostV1Sandboxes structure.
type PostV1SandboxesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1SandboxesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1SandboxesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/sandboxes] PostV1Sandboxes", response, response.Code())
	}
}

// NewPostV1SandboxesCreated creates a PostV1SandboxesCreated with default headers values
func NewPostV1SandboxesCreated() *PostV1SandboxesCreated {
	return &PostV1SandboxesCreated{}
}

/*
PostV1SandboxesCreated describes a response with status code 201, with default header values.

Created
*/
type PostV1SandboxesCreated struct {
	Payload *models.AppSandbox
}

// IsSuccess returns true when this post v1 sandboxes created response has a 2xx status code
func (o *PostV1SandboxesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 sandboxes created response has a 3xx status code
func (o *PostV1SandboxesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 sandboxes created response has a 4xx status code
func (o *PostV1SandboxesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 sandboxes created response has a 5xx status code
func (o *PostV1SandboxesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 sandboxes created response a status code equal to that given
func (o *PostV1SandboxesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 sandboxes created response
func (o *PostV1SandboxesCreated) Code() int {
	return 201
}

func (o *PostV1SandboxesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/sandboxes][%d] postV1SandboxesCreated  %+v", 201, o.Payload)
}

func (o *PostV1SandboxesCreated) String() string {
	return fmt.Sprintf("[POST /v1/sandboxes][%d] postV1SandboxesCreated  %+v", 201, o.Payload)
}

func (o *PostV1SandboxesCreated) GetPayload() *models.AppSandbox {
	return o.Payload
}

func (o *PostV1SandboxesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppSandbox)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}