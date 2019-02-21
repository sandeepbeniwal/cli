// Code generated by go-swagger; DO NOT EDIT.

package fns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	modelsv2 "github.com/fnproject/fn_go/modelsv2"
)

// ListFnsReader is a Reader for the ListFns structure.
type ListFnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListFnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListFnsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFnsOK creates a ListFnsOK with default headers values
func NewListFnsOK() *ListFnsOK {
	return &ListFnsOK{}
}

/*ListFnsOK handles this case with default header values.

List of Functions.
*/
type ListFnsOK struct {
	Payload *modelsv2.FnList
}

func (o *ListFnsOK) Error() string {
	return fmt.Sprintf("[GET /fns][%d] listFnsOK  %+v", 200, o.Payload)
}

func (o *ListFnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.FnList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFnsDefault creates a ListFnsDefault with default headers values
func NewListFnsDefault(code int) *ListFnsDefault {
	return &ListFnsDefault{
		_statusCode: code,
	}
}

/*ListFnsDefault handles this case with default header values.

Error
*/
type ListFnsDefault struct {
	_statusCode int

	Payload *modelsv2.Error
}

// Code gets the status code for the list fns default response
func (o *ListFnsDefault) Code() int {
	return o._statusCode
}

func (o *ListFnsDefault) Error() string {
	return fmt.Sprintf("[GET /fns][%d] ListFns default  %+v", o._statusCode, o.Payload)
}

func (o *ListFnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
