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

// CommentReader is a Reader for the Comment structure.
type CommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCommentOK creates a CommentOK with default headers values
func NewCommentOK() *CommentOK {
	return &CommentOK{}
}

/*
CommentOK describes a response with status code 200, with default header values.

successful operation
*/
type CommentOK struct {
	Payload *models.NetCommentResponse
}

// IsSuccess returns true when this comment o k response has a 2xx status code
func (o *CommentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this comment o k response has a 3xx status code
func (o *CommentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this comment o k response has a 4xx status code
func (o *CommentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this comment o k response has a 5xx status code
func (o *CommentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this comment o k response a status code equal to that given
func (o *CommentOK) IsCode(code int) bool {
	return code == 200
}

func (o *CommentOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/network/comment][%d] commentOK  %+v", 200, o.Payload)
}

func (o *CommentOK) String() string {
	return fmt.Sprintf("[POST /api/v2/network/comment][%d] commentOK  %+v", 200, o.Payload)
}

func (o *CommentOK) GetPayload() *models.NetCommentResponse {
	return o.Payload
}

func (o *CommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetCommentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
