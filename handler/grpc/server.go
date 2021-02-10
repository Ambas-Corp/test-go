package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"github.com/bns-engineering/platformbanking-card/common/logging"
	pt "github.com/bns-engineering/platformbanking-card/handler/grpc/proto"
	"github.com/bns-engineering/platformbanking-presenter/common/lookup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server structure
type Server struct {
	GrpcServer *grpc.Server
	DBConn     *sql.DB
}

// New - creates new grpc server
func New() *Server {
	server := Server{}

	s := grpc.NewServer()
	pt.RegisterCardHandlerServer(s, &server)

	reflection.Register(s)

	server.GrpcServer = s
	return &server
}

// Start - starts grpc server
func (s *Server) Start(port string) {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(port))
		if err != nil {
			strlog := fmt.Sprintf("failed to listen: %v", err)
			logging.InfoLn(strlog)
		}

		strlog := fmt.Sprintf("GRPC Started. Listening on port: %s", port)
		logging.InfoLn(strlog)

		if err := s.GrpcServer.Serve(lis); err != nil {
			strlog := fmt.Sprintf("failed to serve: %s", err)
			logging.InfoLn(strlog)
		}
	}()
}

// Ping - get ping pong
func (s *Server) Ping(ctx context.Context, in *pt.PingRq) (reply *pt.PingRp, err error) {
	reply = &pt.PingRp{}
	reply.ResponseCode = lookup.SUCCESS
	return
}
