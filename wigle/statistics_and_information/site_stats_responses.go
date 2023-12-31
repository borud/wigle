// Code generated by go-swagger; DO NOT EDIT.

package statistics_and_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SiteStatsReader is a Reader for the SiteStats structure.
type SiteStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SiteStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSiteStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSiteStatsOK creates a SiteStatsOK with default headers values
func NewSiteStatsOK() *SiteStatsOK {
	return &SiteStatsOK{}
}

/*
SiteStatsOK describes a response with status code 200, with default header values.

successful operation
*/
type SiteStatsOK struct {
	Payload map[string]interface{}
}

// IsSuccess returns true when this site stats o k response has a 2xx status code
func (o *SiteStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this site stats o k response has a 3xx status code
func (o *SiteStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this site stats o k response has a 4xx status code
func (o *SiteStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this site stats o k response has a 5xx status code
func (o *SiteStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this site stats o k response a status code equal to that given
func (o *SiteStatsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SiteStatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/stats/site][%d] siteStatsOK  %+v", 200, o.Payload)
}

func (o *SiteStatsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/stats/site][%d] siteStatsOK  %+v", 200, o.Payload)
}

func (o *SiteStatsOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *SiteStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
