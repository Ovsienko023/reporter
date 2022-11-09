package grpc

import (
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/transport/grpc/grpc_core"
	domain "github.com/Ovsienko023/reporter/app/transport/grpc/grpc_domain"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func NewWebServer(grpcServer *grpc.Server) *grpcweb.WrappedGrpcServer {

	srv := grpcweb.WrapServer(grpcServer,
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(true),
	)
	return srv
}

func NewServer(core *core.Core) *grpc.Server {
	c := grpc_core.New(core)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				//grpcZap.UnaryServerInterceptor(core.GetLogger()),
				grpcRecovery.UnaryServerInterceptor(),
			),
		),
	)

	domain.RegisterReporterServer(grpcServer, c)
	return grpcServer
}

//func NewGateway(addr string, ctx context.Context) (*runtime.ServeMux, error) {
//	var (
//		grpcGateway = runtime.NewServeMux(
//			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
//				Marshaler: &runtime.JSONPb{
//					MarshalOptions: protojson.MarshalOptions{
//						Multiline:       false,
//						Indent:          "",
//						AllowPartial:    false,
//						UseProtoNames:   true,
//						UseEnumNumbers:  false,
//						EmitUnpopulated: true,
//						Resolver:        nil,
//					},
//				},
//			}),
//			runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
//				md := metadata.Pairs("RemoteAddr", request.RemoteAddr)
//				ctx = metadata.NewIncomingContext(ctx, md)
//				return md
//			}),
//
//			runtime.WithMarshalerOption("multipart/form-data", marshallers.GetFormData()),
//		)
//		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
//	)
//
//	if err := domain.RegisterManagerHandlerFromEndpoint(ctx, grpcGateway, addr, opts); err != nil {
//		return nil, err
//	}
//	domain.Register
//	return grpcGateway, nil
//}

//func NewGatewayWithWebsocketProxy(addr string) (http.Handler, error) {
//	var ctx = context.Background()
//
//	grpcGateway, err := NewGateway(addr, ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return wsproxy.WebsocketProxy(
//		grpcGateway,
//		wsproxy.WithForwardedHeaders(func(header string) bool {
//			return true
//		}),
//	), nil
//}
