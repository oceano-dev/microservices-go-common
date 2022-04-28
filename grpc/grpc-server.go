package proto

import (
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_otel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/oceano-dev/microservices-go-common/config"
	"github.com/oceano-dev/microservices-go-common/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcServer struct {
	config              *config.Config
	managerCertificates *security.ManagerCertificates
}

func NewGrpcServer(
	config *config.Config,
	managerCertificates *security.ManagerCertificates,
) *GrpcServer {
	return &GrpcServer{
		config:              config,
		managerCertificates: managerCertificates,
	}
}

func (s *GrpcServer) CreateGrpcServer() (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(s.config.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(s.config.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(s.config.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(s.config.Grpc.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_otel.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpcrecovery.UnaryServerInterceptor(),
		),
		),
	)

	return grpcServer, nil
}