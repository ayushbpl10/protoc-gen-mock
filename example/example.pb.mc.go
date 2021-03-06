// Code generated by protoc-gen-defaults. DO NOT EDIT.

package example

import (
	"context"
	"github.com/google/gofuzz"
)

var fuzzer = fuzz.New()

// MockEchoServiceServer is the mock implementation of the EchoServiceServer. Use this to create mock services that
// return random data. Useful in UI Testing.
type MockEchoServiceServer struct{}

// Echo is mock implementation of the method Echo
func (MockEchoServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	var res EchoResponse
	fuzzer.Fuzz(&res)
	return &res, nil
}
