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

// PatchV1AppsAppIDSandboxReader is a Reader for the PatchV1AppsAppIDSandbox structure.
type PatchV1AppsAppIDSandboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1AppsAppIDSandboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1AppsAppIDSandboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/apps/{app_id}/sandbox] PatchV1AppsAppIDSandbox", response, response.Code())
	}
}

// NewPatchV1AppsAppIDSandboxOK creates a PatchV1AppsAppIDSandboxOK with default headers values
func NewPatchV1AppsAppIDSandboxOK() *PatchV1AppsAppIDSandboxOK {
	return &PatchV1AppsAppIDSandboxOK{}
}

/*
PatchV1AppsAppIDSandboxOK describes a response with status code 200, with default header values.

OK
*/
type PatchV1AppsAppIDSandboxOK struct {
	Payload *models.AppApp
}

// IsSuccess returns true when this patch v1 apps app Id sandbox o k response has a 2xx status code
func (o *PatchV1AppsAppIDSandboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 apps app Id sandbox o k response has a 3xx status code
func (o *PatchV1AppsAppIDSandboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 apps app Id sandbox o k response has a 4xx status code
func (o *PatchV1AppsAppIDSandboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 apps app Id sandbox o k response has a 5xx status code
func (o *PatchV1AppsAppIDSandboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 apps app Id sandbox o k response a status code equal to that given
func (o *PatchV1AppsAppIDSandboxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 apps app Id sandbox o k response
func (o *PatchV1AppsAppIDSandboxOK) Code() int {
	return 200
}

func (o *PatchV1AppsAppIDSandboxOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}/sandbox][%d] patchV1AppsAppIdSandboxOK  %+v", 200, o.Payload)
}

func (o *PatchV1AppsAppIDSandboxOK) String() string {
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}/sandbox][%d] patchV1AppsAppIdSandboxOK  %+v", 200, o.Payload)
}

func (o *PatchV1AppsAppIDSandboxOK) GetPayload() *models.AppApp {
	return o.Payload
}

func (o *PatchV1AppsAppIDSandboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}