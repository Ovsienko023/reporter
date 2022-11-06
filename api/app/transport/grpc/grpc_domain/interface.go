package grpc_domain

import "context"

type Server struct {
	UnimplementedReporterServer
}

type ServerInterface interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
}
