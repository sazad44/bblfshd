package protocol

import (
	xcontext "golang.org/x/net/context"
)

type protocolServiceServer struct {
}

func NewProtocolServiceServer() *protocolServiceServer {
	return &protocolServiceServer{}
}
func (s *protocolServiceServer) DriverInstanceStates(ctx xcontext.Context, in *DriverInstanceStatesRequest) (result *DriverInstanceStatesResponse, err error) {
	result = new(DriverInstanceStatesResponse)
	result = DriverInstanceStates()
	return
}
func (s *protocolServiceServer) DriverPoolStates(ctx xcontext.Context, in *DriverPoolStatesRequest) (result *DriverPoolStatesResponse, err error) {
	result = new(DriverPoolStatesResponse)
	result = DriverPoolStates()
	return
}
