package grpc_core

import (
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/transport/grpc/grpc_domain"
)

type server struct {
	grpc_domain.UnimplementedReporterServer
	core *core.Core
}

func (s *server) Core() *core.Core { return s.core }

func New(app *core.Core) *server {
	return &server{core: app}
}
