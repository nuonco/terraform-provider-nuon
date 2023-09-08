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

// PostV1AppsAppIDInstallsReader is a Reader for the PostV1AppsAppIDInstalls structure.
type PostV1AppsAppIDInstallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1AppsAppIDInstallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1AppsAppIDInstallsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/installs/] PostV1AppsAppIDInstalls", response, response.Code())
	}
}

// NewPostV1AppsAppIDInstallsCreated creates a PostV1AppsAppIDInstallsCreated with default headers values
func NewPostV1AppsAppIDInstallsCreated() *PostV1AppsAppIDInstallsCreated {
	return &PostV1AppsAppIDInstallsCreated{}
}

/*
PostV1AppsAppIDInstallsCreated describes a response with status code 201, with default header values.

Created
*/
type PostV1AppsAppIDInstallsCreated struct {
	Payload *models.AppInstall
}

// IsSuccess returns true when this post v1 apps app Id installs created response has a 2xx status code
func (o *PostV1AppsAppIDInstallsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 apps app Id installs created response has a 3xx status code
func (o *PostV1AppsAppIDInstallsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 apps app Id installs created response has a 4xx status code
func (o *PostV1AppsAppIDInstallsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 apps app Id installs created response has a 5xx status code
func (o *PostV1AppsAppIDInstallsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 apps app Id installs created response a status code equal to that given
func (o *PostV1AppsAppIDInstallsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 apps app Id installs created response
func (o *PostV1AppsAppIDInstallsCreated) Code() int {
	return 201
}

func (o *PostV1AppsAppIDInstallsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/installs/][%d] postV1AppsAppIdInstallsCreated  %+v", 201, o.Payload)
}

func (o *PostV1AppsAppIDInstallsCreated) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/installs/][%d] postV1AppsAppIdInstallsCreated  %+v", 201, o.Payload)
}

func (o *PostV1AppsAppIDInstallsCreated) GetPayload() *models.AppInstall {
	return o.Payload
}

func (o *PostV1AppsAppIDInstallsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}