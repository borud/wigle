// Code generated by go-swagger; DO NOT EDIT.

package network_observation_file_upload_and_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/borud/wigle/models"
)

// GetTransLogsForUserReader is a Reader for the GetTransLogsForUser structure.
type GetTransLogsForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransLogsForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransLogsForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTransLogsForUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransLogsForUserOK creates a GetTransLogsForUserOK with default headers values
func NewGetTransLogsForUserOK() *GetTransLogsForUserOK {
	return &GetTransLogsForUserOK{}
}

/*
GetTransLogsForUserOK describes a response with status code 200, with default header values.

Request successful
*/
type GetTransLogsForUserOK struct {
	Payload *models.TranslogResponse
}

// IsSuccess returns true when this get trans logs for user o k response has a 2xx status code
func (o *GetTransLogsForUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get trans logs for user o k response has a 3xx status code
func (o *GetTransLogsForUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trans logs for user o k response has a 4xx status code
func (o *GetTransLogsForUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get trans logs for user o k response has a 5xx status code
func (o *GetTransLogsForUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get trans logs for user o k response a status code equal to that given
func (o *GetTransLogsForUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTransLogsForUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/file/transactions][%d] getTransLogsForUserOK  %+v", 200, o.Payload)
}

func (o *GetTransLogsForUserOK) String() string {
	return fmt.Sprintf("[GET /api/v2/file/transactions][%d] getTransLogsForUserOK  %+v", 200, o.Payload)
}

func (o *GetTransLogsForUserOK) GetPayload() *models.TranslogResponse {
	return o.Payload
}

func (o *GetTransLogsForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TranslogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransLogsForUserUnauthorized creates a GetTransLogsForUserUnauthorized with default headers values
func NewGetTransLogsForUserUnauthorized() *GetTransLogsForUserUnauthorized {
	return &GetTransLogsForUserUnauthorized{}
}

/*
GetTransLogsForUserUnauthorized describes a response with status code 401, with default header values.

User not authenticated
*/
type GetTransLogsForUserUnauthorized struct {
}

// IsSuccess returns true when this get trans logs for user unauthorized response has a 2xx status code
func (o *GetTransLogsForUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trans logs for user unauthorized response has a 3xx status code
func (o *GetTransLogsForUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trans logs for user unauthorized response has a 4xx status code
func (o *GetTransLogsForUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trans logs for user unauthorized response has a 5xx status code
func (o *GetTransLogsForUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get trans logs for user unauthorized response a status code equal to that given
func (o *GetTransLogsForUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTransLogsForUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/file/transactions][%d] getTransLogsForUserUnauthorized ", 401)
}

func (o *GetTransLogsForUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/file/transactions][%d] getTransLogsForUserUnauthorized ", 401)
}

func (o *GetTransLogsForUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
