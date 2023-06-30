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

// WifiReader is a Reader for the Wifi structure.
type WifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWifiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWifiUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewWifiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewWifiNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewWifiTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewWifiInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWifiOK creates a WifiOK with default headers values
func NewWifiOK() *WifiOK {
	return &WifiOK{}
}

/*
WifiOK describes a response with status code 200, with default header values.

Request succeeded
*/
type WifiOK struct {
	Payload *models.WiFiDetail
}

// IsSuccess returns true when this wifi o k response has a 2xx status code
func (o *WifiOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wifi o k response has a 3xx status code
func (o *WifiOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wifi o k response has a 4xx status code
func (o *WifiOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wifi o k response has a 5xx status code
func (o *WifiOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wifi o k response a status code equal to that given
func (o *WifiOK) IsCode(code int) bool {
	return code == 200
}

func (o *WifiOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiOK  %+v", 200, o.Payload)
}

func (o *WifiOK) String() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiOK  %+v", 200, o.Payload)
}

func (o *WifiOK) GetPayload() *models.WiFiDetail {
	return o.Payload
}

func (o *WifiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WiFiDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWifiUnauthorized creates a WifiUnauthorized with default headers values
func NewWifiUnauthorized() *WifiUnauthorized {
	return &WifiUnauthorized{}
}

/*
WifiUnauthorized describes a response with status code 401, with default header values.

Not Authorized
*/
type WifiUnauthorized struct {
}

// IsSuccess returns true when this wifi unauthorized response has a 2xx status code
func (o *WifiUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this wifi unauthorized response has a 3xx status code
func (o *WifiUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wifi unauthorized response has a 4xx status code
func (o *WifiUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this wifi unauthorized response has a 5xx status code
func (o *WifiUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this wifi unauthorized response a status code equal to that given
func (o *WifiUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *WifiUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiUnauthorized ", 401)
}

func (o *WifiUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiUnauthorized ", 401)
}

func (o *WifiUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWifiForbidden creates a WifiForbidden with default headers values
func NewWifiForbidden() *WifiForbidden {
	return &WifiForbidden{}
}

/*
WifiForbidden describes a response with status code 403, with default header values.

Commercial request made for non-commercial resource
*/
type WifiForbidden struct {
}

// IsSuccess returns true when this wifi forbidden response has a 2xx status code
func (o *WifiForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this wifi forbidden response has a 3xx status code
func (o *WifiForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wifi forbidden response has a 4xx status code
func (o *WifiForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this wifi forbidden response has a 5xx status code
func (o *WifiForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this wifi forbidden response a status code equal to that given
func (o *WifiForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *WifiForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiForbidden ", 403)
}

func (o *WifiForbidden) String() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiForbidden ", 403)
}

func (o *WifiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWifiNotFound creates a WifiNotFound with default headers values
func NewWifiNotFound() *WifiNotFound {
	return &WifiNotFound{}
}

/*
WifiNotFound describes a response with status code 404, with default header values.

No matching network ID found.
*/
type WifiNotFound struct {
}

// IsSuccess returns true when this wifi not found response has a 2xx status code
func (o *WifiNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this wifi not found response has a 3xx status code
func (o *WifiNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wifi not found response has a 4xx status code
func (o *WifiNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this wifi not found response has a 5xx status code
func (o *WifiNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this wifi not found response a status code equal to that given
func (o *WifiNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *WifiNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiNotFound ", 404)
}

func (o *WifiNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiNotFound ", 404)
}

func (o *WifiNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWifiTooManyRequests creates a WifiTooManyRequests with default headers values
func NewWifiTooManyRequests() *WifiTooManyRequests {
	return &WifiTooManyRequests{}
}

/*
WifiTooManyRequests describes a response with status code 429, with default header values.

too many DETAIL queries today.
*/
type WifiTooManyRequests struct {
}

// IsSuccess returns true when this wifi too many requests response has a 2xx status code
func (o *WifiTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this wifi too many requests response has a 3xx status code
func (o *WifiTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wifi too many requests response has a 4xx status code
func (o *WifiTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this wifi too many requests response has a 5xx status code
func (o *WifiTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this wifi too many requests response a status code equal to that given
func (o *WifiTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *WifiTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiTooManyRequests ", 429)
}

func (o *WifiTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiTooManyRequests ", 429)
}

func (o *WifiTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWifiInternalServerError creates a WifiInternalServerError with default headers values
func NewWifiInternalServerError() *WifiInternalServerError {
	return &WifiInternalServerError{}
}

/*
WifiInternalServerError describes a response with status code 500, with default header values.

Failed to retrieve due to error.
*/
type WifiInternalServerError struct {
}

// IsSuccess returns true when this wifi internal server error response has a 2xx status code
func (o *WifiInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this wifi internal server error response has a 3xx status code
func (o *WifiInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wifi internal server error response has a 4xx status code
func (o *WifiInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this wifi internal server error response has a 5xx status code
func (o *WifiInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this wifi internal server error response a status code equal to that given
func (o *WifiInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *WifiInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiInternalServerError ", 500)
}

func (o *WifiInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v3/detail/wifi/{wifiNetworkId}][%d] wifiInternalServerError ", 500)
}

func (o *WifiInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}