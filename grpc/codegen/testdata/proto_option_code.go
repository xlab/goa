package testdata

const ServiceWithProtoOptionCode = `
syntax = "proto3";

package service_with_proto_option;

option go_package = "/service_with_proto_optionpb";

// Service is the ServiceWithProtoOption service interface.
service ServiceWithProtoOption {
	option deprecated = false;
	// ServiceWithProtoOptionMethod implements ServiceWithProtoOptionMethod.
	rpc ServiceWithProtoOptionMethod (ServiceWithProtoOptionMethodRequest) returns (ServiceWithProtoOptionMethodResponse) {
		option deprecated = true;
	}
}

message ServiceWithProtoOptionMethodRequest {
}

message ServiceWithProtoOptionMethodResponse {
}
`
