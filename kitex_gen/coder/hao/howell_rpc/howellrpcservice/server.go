// Code generated by Kitex v0.8.0. DO NOT EDIT.
package howellrpcservice

import (
	server "github.com/cloudwego/kitex/server"
	howell_rpc "howell/howell_api/kitex_gen/coder/hao/howell_rpc"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler howell_rpc.HowellRpcService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}