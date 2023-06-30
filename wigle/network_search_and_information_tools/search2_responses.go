// Code generated by go-swagger; DO NOT EDIT.

package network_search_and_information_tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/borud/wigle/models"
)

// Search2Reader is a Reader for the Search2 structure.
type Search2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Search2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearch2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearch2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewSearch2PaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewSearch2Gone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSearch2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearch2OK creates a Search2OK with default headers values
func NewSearch2OK() *Search2OK {
	return &Search2OK{}
}

/*
Search2OK describes a response with status code 200, with default header values.

Request succeeded
*/
type Search2OK struct {
	Payload *models.NetSearchResponse
}

// IsSuccess returns true when this search2 o k response has a 2xx status code
func (o *Search2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search2 o k response has a 3xx status code
func (o *Search2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search2 o k response has a 4xx status code
func (o *Search2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search2 o k response has a 5xx status code
func (o *Search2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this search2 o k response a status code equal to that given
func (o *Search2OK) IsCode(code int) bool {
	return code == 200
}

func (o *Search2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2OK  %+v", 200, o.Payload)
}

func (o *Search2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2OK  %+v", 200, o.Payload)
}

func (o *Search2OK) GetPayload() *models.NetSearchResponse {
	return o.Payload
}

func (o *Search2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearch2BadRequest creates a Search2BadRequest with default headers values
func NewSearch2BadRequest() *Search2BadRequest {
	return &Search2BadRequest{}
}

/*
Search2BadRequest describes a response with status code 400, with default header values.

Request body error
*/
type Search2BadRequest struct {
}

// IsSuccess returns true when this search2 bad request response has a 2xx status code
func (o *Search2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search2 bad request response has a 3xx status code
func (o *Search2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search2 bad request response has a 4xx status code
func (o *Search2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search2 bad request response has a 5xx status code
func (o *Search2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search2 bad request response a status code equal to that given
func (o *Search2BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *Search2BadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2BadRequest ", 400)
}

func (o *Search2BadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2BadRequest ", 400)
}

func (o *Search2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearch2PaymentRequired creates a Search2PaymentRequired with default headers values
func NewSearch2PaymentRequired() *Search2PaymentRequired {
	return &Search2PaymentRequired{}
}

/*
Search2PaymentRequired describes a response with status code 402, with default header values.

Insufficient balance for commercial query
*/
type Search2PaymentRequired struct {
}

// IsSuccess returns true when this search2 payment required response has a 2xx status code
func (o *Search2PaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search2 payment required response has a 3xx status code
func (o *Search2PaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search2 payment required response has a 4xx status code
func (o *Search2PaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this search2 payment required response has a 5xx status code
func (o *Search2PaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this search2 payment required response a status code equal to that given
func (o *Search2PaymentRequired) IsCode(code int) bool {
	return code == 402
}

func (o *Search2PaymentRequired) Error() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2PaymentRequired ", 402)
}

func (o *Search2PaymentRequired) String() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2PaymentRequired ", 402)
}

func (o *Search2PaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearch2Gone creates a Search2Gone with default headers values
func NewSearch2Gone() *Search2Gone {
	return &Search2Gone{}
}

/*
Search2Gone describes a response with status code 410, with default header values.

Query Failed
*/
type Search2Gone struct {
}

// IsSuccess returns true when this search2 gone response has a 2xx status code
func (o *Search2Gone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search2 gone response has a 3xx status code
func (o *Search2Gone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search2 gone response has a 4xx status code
func (o *Search2Gone) IsClientError() bool {
	return true
}

// IsServerError returns true when this search2 gone response has a 5xx status code
func (o *Search2Gone) IsServerError() bool {
	return false
}

// IsCode returns true when this search2 gone response a status code equal to that given
func (o *Search2Gone) IsCode(code int) bool {
	return code == 410
}

func (o *Search2Gone) Error() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2Gone ", 410)
}

func (o *Search2Gone) String() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2Gone ", 410)
}

func (o *Search2Gone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearch2TooManyRequests creates a Search2TooManyRequests with default headers values
func NewSearch2TooManyRequests() *Search2TooManyRequests {
	return &Search2TooManyRequests{}
}

/*
Search2TooManyRequests describes a response with status code 429, with default header values.

too many queries today.
*/
type Search2TooManyRequests struct {
}

// IsSuccess returns true when this search2 too many requests response has a 2xx status code
func (o *Search2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search2 too many requests response has a 3xx status code
func (o *Search2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search2 too many requests response has a 4xx status code
func (o *Search2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this search2 too many requests response has a 5xx status code
func (o *Search2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this search2 too many requests response a status code equal to that given
func (o *Search2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *Search2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2TooManyRequests ", 429)
}

func (o *Search2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/network/search][%d] search2TooManyRequests ", 429)
}

func (o *Search2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
