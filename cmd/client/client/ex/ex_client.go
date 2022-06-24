// Code generated by go-swagger; DO NOT EDIT.

package ex

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ex API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ex API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCar(params *CreateCarParams, opts ...ClientOption) (*CreateCarOK, error)

	ListCar(params *ListCarParams, opts ...ClientOption) (*ListCarOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCar creates car

  Create Car
*/
func (a *Client) CreateCar(params *CreateCarParams, opts ...ClientOption) (*CreateCarOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCar",
		Method:             "POST",
		PathPattern:        "/car",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateCar: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCar lists car

  List car
*/
func (a *Client) ListCar(params *ListCarParams, opts ...ClientOption) (*ListCarOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListCar",
		Method:             "GET",
		PathPattern:        "/car",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCarOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCar: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
