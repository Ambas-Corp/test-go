package client

import (
	"context"

	pt "github.com/bns-engineering/platformbanking-card/handler/grpc/proto"
)

type (
	//IClient - interface for client grpc
	IClient interface {
		Ping(ctx context.Context, message string) (*pt.PingRp, error)
	}
)
