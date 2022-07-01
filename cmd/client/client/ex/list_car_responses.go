// Code generated by go-swagger; DO NOT EDIT.

package ex

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"kunlun-qilian/server-example/cmd/client/models"
)

// ListCarReader is a Reader for the ListCar structure.
type ListCarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCarOK creates a ListCarOK with default headers values
func NewListCarOK() *ListCarOK {
	return &ListCarOK{}
}

/* ListCarOK describes a response with status code 200, with default header values.

OK
*/
type ListCarOK struct {
	Payload []*models.ModelTExample
}

func (o *ListCarOK) Error() string {
	return fmt.Sprintf("[GET /car][%d] listCarOK  %+v", 200, o.Payload)
}
func (o *ListCarOK) GetPayload() []*models.ModelTExample {
	return o.Payload
}

func (o *ListCarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
