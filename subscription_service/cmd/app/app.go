package main

import (
	"net"
	"os"
	"subscription_service/internal/handlers"
	"subscription_service/proto"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ApplicationServer struct {
	db          *gorm.DB
	logger      *zerolog.Logger
	grpcHandler *handlers.GRPCHandler
}

func NewApplicationServer(db *gorm.DB, logger *zerolog.Logger, grpcHandler *handlers.GRPCHandler) *ApplicationServer {
	return &ApplicationServer{db: db, logger: logger, grpcHandler: grpcHandler}
}

func (a *ApplicationServer) StartHttpServer() {

}

func (a *ApplicationServer) StartGRPCServer() {
	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		a.logger.Error().Err(err).Msgf("Failed to listen on grpc port %s", os.Getenv("GRPC_PORT"))
		panic(err)
	}

	grpcServer := grpc.NewServer()
	// Register the gRPC handler
	proto.RegisterSubscriptionServiceServer(grpcServer, a.grpcHandler)

	if err := grpcServer.Serve(lis); err != nil {
		a.logger.Error().Err(err).Msg("Failed to serve grpc server")
		panic(err)
	}
}
