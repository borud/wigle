// Code generated by go-swagger; DO NOT EDIT.

package stats_group_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/borud/wigle/models"
)

// GroupMembersReader is a Reader for the GroupMembers structure.
type GroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupMembersOK creates a GroupMembersOK with default headers values
func NewGroupMembersOK() *GroupMembersOK {
	return &GroupMembersOK{}
}

/*
GroupMembersOK describes a response with status code 200, with default header values.

successful operation
*/
type GroupMembersOK struct {
	Payload *models.GroupMemberResponse
}

// IsSuccess returns true when this group members o k response has a 2xx status code
func (o *GroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group members o k response has a 3xx status code
func (o *GroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group members o k response has a 4xx status code
func (o *GroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group members o k response has a 5xx status code
func (o *GroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group members o k response a status code equal to that given
func (o *GroupMembersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/group/groupMembers][%d] groupMembersOK  %+v", 200, o.Payload)
}

func (o *GroupMembersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/group/groupMembers][%d] groupMembersOK  %+v", 200, o.Payload)
}

func (o *GroupMembersOK) GetPayload() *models.GroupMemberResponse {
	return o.Payload
}

func (o *GroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupMemberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupMembersBadRequest creates a GroupMembersBadRequest with default headers values
func NewGroupMembersBadRequest() *GroupMembersBadRequest {
	return &GroupMembersBadRequest{}
}

/*
GroupMembersBadRequest describes a response with status code 400, with default header values.

Unable to load group
*/
type GroupMembersBadRequest struct {
}

// IsSuccess returns true when this group members bad request response has a 2xx status code
func (o *GroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this group members bad request response has a 3xx status code
func (o *GroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group members bad request response has a 4xx status code
func (o *GroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this group members bad request response has a 5xx status code
func (o *GroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this group members bad request response a status code equal to that given
func (o *GroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/group/groupMembers][%d] groupMembersBadRequest ", 400)
}

func (o *GroupMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/group/groupMembers][%d] groupMembersBadRequest ", 400)
}

func (o *GroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupMembersNotFound creates a GroupMembersNotFound with default headers values
func NewGroupMembersNotFound() *GroupMembersNotFound {
	return &GroupMembersNotFound{}
}

/*
GroupMembersNotFound describes a response with status code 404, with default header values.

Group not found
*/
type GroupMembersNotFound struct {
}

// IsSuccess returns true when this group members not found response has a 2xx status code
func (o *GroupMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this group members not found response has a 3xx status code
func (o *GroupMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group members not found response has a 4xx status code
func (o *GroupMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this group members not found response has a 5xx status code
func (o *GroupMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this group members not found response a status code equal to that given
func (o *GroupMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GroupMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/group/groupMembers][%d] groupMembersNotFound ", 404)
}

func (o *GroupMembersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/group/groupMembers][%d] groupMembersNotFound ", 404)
}

func (o *GroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}