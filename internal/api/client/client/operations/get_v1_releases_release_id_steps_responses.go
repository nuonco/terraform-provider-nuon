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

// GetV1ReleasesReleaseIDStepsReader is a Reader for the GetV1ReleasesReleaseIDSteps structure.
type GetV1ReleasesReleaseIDStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ReleasesReleaseIDStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ReleasesReleaseIDStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases/{release_id}/steps] GetV1ReleasesReleaseIDSteps", response, response.Code())
	}
}

// NewGetV1ReleasesReleaseIDStepsOK creates a GetV1ReleasesReleaseIDStepsOK with default headers values
func NewGetV1ReleasesReleaseIDStepsOK() *GetV1ReleasesReleaseIDStepsOK {
	return &GetV1ReleasesReleaseIDStepsOK{}
}

/*
GetV1ReleasesReleaseIDStepsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1ReleasesReleaseIDStepsOK struct {
	Payload []*models.AppComponentReleaseStep
}

// IsSuccess returns true when this get v1 releases release Id steps o k response has a 2xx status code
func (o *GetV1ReleasesReleaseIDStepsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 releases release Id steps o k response has a 3xx status code
func (o *GetV1ReleasesReleaseIDStepsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 releases release Id steps o k response has a 4xx status code
func (o *GetV1ReleasesReleaseIDStepsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 releases release Id steps o k response has a 5xx status code
func (o *GetV1ReleasesReleaseIDStepsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 releases release Id steps o k response a status code equal to that given
func (o *GetV1ReleasesReleaseIDStepsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 releases release Id steps o k response
func (o *GetV1ReleasesReleaseIDStepsOK) Code() int {
	return 200
}

func (o *GetV1ReleasesReleaseIDStepsOK) Error() string {
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getV1ReleasesReleaseIdStepsOK  %+v", 200, o.Payload)
}

func (o *GetV1ReleasesReleaseIDStepsOK) String() string {
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getV1ReleasesReleaseIdStepsOK  %+v", 200, o.Payload)
}

func (o *GetV1ReleasesReleaseIDStepsOK) GetPayload() []*models.AppComponentReleaseStep {
	return o.Payload
}

func (o *GetV1ReleasesReleaseIDStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}