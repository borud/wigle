// Code generated by go-swagger; DO NOT EDIT.

package statistics_and_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/borud/wigle/models"
)

// GroupStatsReader is a Reader for the GroupStats structure.
type GroupStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGroupStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupStatsOK creates a GroupStatsOK with default headers values
func NewGroupStatsOK() *GroupStatsOK {
	return &GroupStatsOK{}
}

/*
GroupStatsOK describes a response with status code 200, with default header values.

Request succeeded
*/
type GroupStatsOK struct {
	Payload *models.GroupStatResponse
}

// IsSuccess returns true when this group stats o k response has a 2xx status code
func (o *GroupStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group stats o k response has a 3xx status code
func (o *GroupStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group stats o k response has a 4xx status code
func (o *GroupStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group stats o k response has a 5xx status code
func (o *GroupStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group stats o k response a status code equal to that given
func (o *GroupStatsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GroupStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/stats/group][%d] groupStatsOK  %+v", 200, o.Payload)
}

func (o *GroupStatsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/stats/group][%d] groupStatsOK  %+v", 200, o.Payload)
}

func (o *GroupStatsOK) GetPayload() *models.GroupStatResponse {
	return o.Payload
}

func (o *GroupStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupStatResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupStatsInternalServerError creates a GroupStatsInternalServerError with default headers values
func NewGroupStatsInternalServerError() *GroupStatsInternalServerError {
	return &GroupStatsInternalServerError{}
}

/*
GroupStatsInternalServerError describes a response with status code 500, with default header values.

Error fetching groups list
*/
type GroupStatsInternalServerError struct {
}

// IsSuccess returns true when this group stats internal server error response has a 2xx status code
func (o *GroupStatsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this group stats internal server error response has a 3xx status code
func (o *GroupStatsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group stats internal server error response has a 4xx status code
func (o *GroupStatsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this group stats internal server error response has a 5xx status code
func (o *GroupStatsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this group stats internal server error response a status code equal to that given
func (o *GroupStatsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GroupStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/stats/group][%d] groupStatsInternalServerError ", 500)
}

func (o *GroupStatsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/stats/group][%d] groupStatsInternalServerError ", 500)
}

func (o *GroupStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}