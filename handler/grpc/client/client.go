package client

import (
	"context"
	"log"
	"time"

	pt "github.com/bns-engineering/platformbanking-card/handler/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

//Options for client
type Options struct {
	Address string
}

//Client client object
type Client struct {
	CardClient pt.CardHandlerClient
	opts       *Options
}

//GetClient - get connection client to param grpc
func GetClient(o *Options) (IClient, error) {
	var (
		conn        *grpc.ClientConn
		dialOptions []grpc.DialOption
	)

	//backoffBase - delay time for reconnect
	backoffBase := backoff.DefaultConfig
	backoffBase.MaxDelay = 10 * time.Second

	//implement Backoff for reconnect grpc
	dialOptions = append(dialOptions, grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoffBase,
	}))

	//implement WithInsecure for disables transport security for this ClientConn.
	dialOptions = append(dialOptions, grpc.WithInsecure())

	//get connection to grpc
	conn, err := grpc.Dial(o.Address, dialOptions...)
	if err != nil {
		log.Println("error dial up grpc: ", o.Address, err)
		if conn != nil {
			conn.Close()
		}

		return nil, err
	}

	client := &Client{}
	client.CardClient = pt.NewCardHandlerClient(conn)

	return client, nil
}

//Ping - return ping pong
func (c *Client) Ping(ctx context.Context, message string) (*pt.PingRp, error) {
	rq := pt.PingRq{
		Message: message,
	}
	return c.CardClient.Ping(ctx, &rq)
}
