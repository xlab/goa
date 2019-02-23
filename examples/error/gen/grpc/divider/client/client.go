// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// divider gRPC client
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package client

import (
	"context"

	goa "goa.design/goa"
	dividerpb "goa.design/goa/examples/error/gen/grpc/divider/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli dividerpb.DividerClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the divider service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: dividerpb.NewDividerClient(cc),
		opts:    opts,
	}
}

// IntegerDivide calls the "IntegerDivide" function in dividerpb.DividerClient
// interface.
func (c *Client) IntegerDivide() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildIntegerDivideFunc(c.grpccli, c.opts...),
			EncodeIntegerDivideRequest,
			DecodeIntegerDivideResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goagrpc.DecodeError(err)
		}
		return res, nil
	}
}

// Divide calls the "Divide" function in dividerpb.DividerClient interface.
func (c *Client) Divide() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDivideFunc(c.grpccli, c.opts...),
			EncodeDivideRequest,
			DecodeDivideResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goagrpc.DecodeError(err)
		}
		return res, nil
	}
}