package grpc

import (
	"context"
	"fmt"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/common/random"
	"github.com/pastelnetwork/gonode/walletnode/node"
	"google.golang.org/grpc"
)

const (
	logPrefix = "client"
)

type client struct{}

// Connect implements node.Client.Connect()
func (client *client) Connect(ctx context.Context, address string) (node.Connection, error) {
	id, _ := random.String(8, random.Base62Chars)
	ctx = context.WithValue(ctx, log.PrefixKey, fmt.Sprintf("%s-%s", logPrefix, id))

	grpcConn, err := grpc.DialContext(ctx, address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.WithContext(ctx).WithError(err).WithField("address", address).Warn("fail to dial")
		return nil, errors.New(err)
	}
	log.WithContext(ctx).Debugf("Connected to %s", address)

	conn := newClientConn(id, grpcConn)
	go func() {
		<-conn.Done()
		log.WithContext(ctx).Debugf("Disconnected %s", grpcConn.Target())
	}()
	return conn, nil
}

// NewClient returns a new client instance.
func NewClient() node.Client {
	return &client{}
}
