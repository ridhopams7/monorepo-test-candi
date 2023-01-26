package warehoouse

import (
	"net/url"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

type warehoouseGRPCImpl struct {
	host    string
	authKey string
	conn    *grpc.ClientConn
}

// NewWarehoouseServiceGRPC constructor
func NewWarehoouseServiceGRPC(host string, authKey string) Warehoouse {

	if u, _ := url.Parse(host); u.Host != "" {
		host = u.Host
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  50 * time.Millisecond,
			Multiplier: 5,
			MaxDelay:   50 * time.Millisecond,
		},
		MinConnectTimeout: 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}

	return &warehoouseGRPCImpl{
		host:    host,
		authKey: authKey,
		conn:    conn,
	}
}
