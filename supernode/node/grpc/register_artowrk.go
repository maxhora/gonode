package grpc

import (
	"context"
	"fmt"
	"io"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	pb "github.com/pastelnetwork/gonode/proto/supernode"
	"github.com/pastelnetwork/gonode/supernode/node"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type registerArtowrk struct {
	conn *clientConn
	pb.SuperNode_RegisterArtowrkClient

	isClosed bool

	recvCh chan *pb.RegisterArtworkReply
	errCh  chan error
}

func (stream *registerArtowrk) Handshake(ctx context.Context, connID, nodeKey string) error {
	ctx = context.WithValue(ctx, log.PrefixKey, fmt.Sprintf("%s-%s", logClientPrefix, stream.conn.id))

	req := &pb.RegisterArtworkRequest{
		Requests: &pb.RegisterArtworkRequest_Handshake{
			Handshake: &pb.RegisterArtworkRequest_HandshakeRequest{
				ConnID:  connID,
				NodeKey: nodeKey,
			},
		},
	}

	res, err := stream.sendRecv(ctx, req)
	if err != nil {
		return err
	}

	resp := res.GetHandshake()
	if resp == nil {
		return errors.Errorf("wrong response, %q", res.String())
	}
	if err := resp.Error; err.Status == pb.RegisterArtworkReply_Error_ERR {
		return errors.New(err.ErrMsg)
	}
	return nil
}

func (stream *registerArtowrk) sendRecv(ctx context.Context, req *pb.RegisterArtworkRequest) (*pb.RegisterArtworkReply, error) {
	if err := stream.send(ctx, req); err != nil {
		return nil, err
	}

	resp, err := stream.recv(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (stream *registerArtowrk) send(ctx context.Context, req *pb.RegisterArtworkRequest) error {
	ctx = context.WithValue(ctx, log.PrefixKey, fmt.Sprintf("%s-%s", logClientPrefix, stream.conn.id))

	if stream.isClosed {
		return errors.New("stream closed")
	}

	log.WithContext(ctx).WithField("req", req.String()).Debugf("Sending")
	if err := stream.SendMsg(req); err != nil {
		switch status.Code(err) {
		case codes.Canceled:
			log.WithContext(ctx).WithError(err).Debugf("Sending canceled")
		default:
			log.WithContext(ctx).WithError(err).Errorf("Sending")
		}
		return err
	}

	return nil
}

func (stream *registerArtowrk) recv(ctx context.Context) (*pb.RegisterArtworkReply, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	case resp := <-stream.recvCh:
		return resp, nil
	case err := <-stream.errCh:
		return nil, err
	}
}

func (stream *registerArtowrk) start(ctx context.Context) error {
	ctx = context.WithValue(ctx, log.PrefixKey, fmt.Sprintf("%s-%s", logClientPrefix, stream.conn.id))

	go func() {
		defer func() {
			stream.isClosed = true
			stream.conn.Close()
		}()

		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.WithContext(ctx).Debug("Stream closed by peer")
					return
				}
				switch status.Code(err) {
				case codes.Canceled, codes.Unavailable:
					log.WithContext(ctx).WithError(err).Debug("Stream closed")
				default:
					log.WithContext(ctx).WithError(err).Warn("Stream")
				}

				stream.errCh <- errors.New(err)
				break
			}
			log.WithContext(ctx).WithField("resp", resp.String()).Debug("Receiving")

			stream.recvCh <- resp
		}
	}()

	return nil
}

func newRegisterArtowrk(conn *clientConn, client pb.SuperNode_RegisterArtowrkClient) node.RegisterArtowrk {
	return &registerArtowrk{
		conn:                            conn,
		SuperNode_RegisterArtowrkClient: client,

		recvCh: make(chan *pb.RegisterArtworkReply),
		errCh:  make(chan error),
	}
}