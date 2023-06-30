// Code generated by go-swagger; DO NOT EDIT.

package v3_a_l_p_h_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/borud/wigle/models"
)

// CellChannelReader is a Reader for the CellChannel structure.
type CellChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CellChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCellChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCellChannelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCellChannelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCellChannelTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCellChannelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCellChannelOK creates a CellChannelOK with default headers values
func NewCellChannelOK() *CellChannelOK {
	return &CellChannelOK{}
}

/*
CellChannelOK describes a response with status code 200, with default header values.

Request succeeded
*/
type CellChannelOK struct {
	Payload *models.ChannelDetailResponse
}

// IsSuccess returns true when this cell channel o k response has a 2xx status code
func (o *CellChannelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cell channel o k response has a 3xx status code
func (o *CellChannelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cell channel o k response has a 4xx status code
func (o *CellChannelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cell channel o k response has a 5xx status code
func (o *CellChannelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cell channel o k response a status code equal to that given
func (o *CellChannelOK) IsCode(code int) bool {
	return code == 200
}

func (o *CellChannelOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelOK  %+v", 200, o.Payload)
}

func (o *CellChannelOK) String() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelOK  %+v", 200, o.Payload)
}

func (o *CellChannelOK) GetPayload() *models.ChannelDetailResponse {
	return o.Payload
}

func (o *CellChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelDetailResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCellChannelUnauthorized creates a CellChannelUnauthorized with default headers values
func NewCellChannelUnauthorized() *CellChannelUnauthorized {
	return &CellChannelUnauthorized{}
}

/*
CellChannelUnauthorized describes a response with status code 401, with default header values.

Error authenticating user
*/
type CellChannelUnauthorized struct {
}

// IsSuccess returns true when this cell channel unauthorized response has a 2xx status code
func (o *CellChannelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cell channel unauthorized response has a 3xx status code
func (o *CellChannelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cell channel unauthorized response has a 4xx status code
func (o *CellChannelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cell channel unauthorized response has a 5xx status code
func (o *CellChannelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cell channel unauthorized response a status code equal to that given
func (o *CellChannelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CellChannelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelUnauthorized ", 401)
}

func (o *CellChannelUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelUnauthorized ", 401)
}

func (o *CellChannelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCellChannelNotFound creates a CellChannelNotFound with default headers values
func NewCellChannelNotFound() *CellChannelNotFound {
	return &CellChannelNotFound{}
}

/*
CellChannelNotFound describes a response with status code 404, with default header values.

No matching cell sites found in region.
*/
type CellChannelNotFound struct {
}

// IsSuccess returns true when this cell channel not found response has a 2xx status code
func (o *CellChannelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cell channel not found response has a 3xx status code
func (o *CellChannelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cell channel not found response has a 4xx status code
func (o *CellChannelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cell channel not found response has a 5xx status code
func (o *CellChannelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cell channel not found response a status code equal to that given
func (o *CellChannelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CellChannelNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelNotFound ", 404)
}

func (o *CellChannelNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelNotFound ", 404)
}

func (o *CellChannelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCellChannelTooManyRequests creates a CellChannelTooManyRequests with default headers values
func NewCellChannelTooManyRequests() *CellChannelTooManyRequests {
	return &CellChannelTooManyRequests{}
}

/*
CellChannelTooManyRequests describes a response with status code 429, with default header values.

too many DETAIL queries today.
*/
type CellChannelTooManyRequests struct {
}

// IsSuccess returns true when this cell channel too many requests response has a 2xx status code
func (o *CellChannelTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cell channel too many requests response has a 3xx status code
func (o *CellChannelTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cell channel too many requests response has a 4xx status code
func (o *CellChannelTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cell channel too many requests response has a 5xx status code
func (o *CellChannelTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cell channel too many requests response a status code equal to that given
func (o *CellChannelTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CellChannelTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelTooManyRequests ", 429)
}

func (o *CellChannelTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelTooManyRequests ", 429)
}

func (o *CellChannelTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCellChannelInternalServerError creates a CellChannelInternalServerError with default headers values
func NewCellChannelInternalServerError() *CellChannelInternalServerError {
	return &CellChannelInternalServerError{}
}

/*
CellChannelInternalServerError describes a response with status code 500, with default header values.

Failed to retrieve due to error.
*/
type CellChannelInternalServerError struct {
}

// IsSuccess returns true when this cell channel internal server error response has a 2xx status code
func (o *CellChannelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cell channel internal server error response has a 3xx status code
func (o *CellChannelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cell channel internal server error response has a 4xx status code
func (o *CellChannelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cell channel internal server error response has a 5xx status code
func (o *CellChannelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cell channel internal server error response a status code equal to that given
func (o *CellChannelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CellChannelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelInternalServerError ", 500)
}

func (o *CellChannelInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v3/cellChannel/{type}][%d] cellChannelInternalServerError ", 500)
}

func (o *CellChannelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}