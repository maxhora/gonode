package grpc

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/proto"
	pb "github.com/pastelnetwork/gonode/proto/walletnode"
	"github.com/pastelnetwork/gonode/walletnode/node"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	uploadImageBufferSize = 32 * 1024
)

type registerArtowrk struct {
	conn   *clientConn
	client pb.RegisterArtowrkClient

	connID string
}

func (service *registerArtowrk) ConnID() string {
	return service.connID
}

func (service *registerArtowrk) healthCheck(ctx context.Context) error {
	ctx = service.contextWithLogPrefix(ctx)
	ctx = service.contextWithMDConnID(ctx)

	stream, err := service.client.HealthCheck(ctx)
	if err != nil {
		return errors.New("failed to open HealthCheck stream")
	}

	go func() {
		defer service.conn.Close()

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.WithContext(ctx).Debug("Stream closed by peer")
				}

				switch status.Code(err) {
				case codes.Canceled, codes.Unavailable:
					log.WithContext(ctx).WithError(err).Debug("Stream closed")
				default:
					log.WithContext(ctx).WithError(err).Error("Stream closed")
				}
				return
			}
		}
	}()

	return nil
}

// Handshake implements node.RegisterArtowrk.Handshake()
func (service *registerArtowrk) Handshake(ctx context.Context, IsPrimary bool) error {
	ctx = service.contextWithLogPrefix(ctx)

	req := &pb.HandshakeRequest{
		IsPrimary: IsPrimary,
	}
	log.WithContext(ctx).WithField("req", req).Debugf("Handshake request")

	resp, err := service.client.Handshake(ctx, req)
	if err != nil {
		return errors.New("failed to reqeust Handshake")
	}
	log.WithContext(ctx).WithField("resp", resp).Debugf("Handshake response")

	service.connID = resp.ConnID
	return service.healthCheck(ctx)
}

// AcceptedNodes implements node.RegisterArtowrk.AcceptedNodes()
func (service *registerArtowrk) AcceptedNodes(ctx context.Context) (pastelIDs []string, err error) {
	ctx = service.contextWithLogPrefix(ctx)
	ctx = service.contextWithMDConnID(ctx)

	req := &pb.AcceptedNodesRequest{}
	log.WithContext(ctx).WithField("req", req).Debugf("AcceptedNodes request")

	resp, err := service.client.AcceptedNodes(ctx, req)
	if err != nil {
		return nil, errors.New("failed to request to accepted secondary nodes")
	}
	log.WithContext(ctx).WithField("resp", resp).Debugf("AcceptedNodes response")

	var ids []string
	for _, peer := range resp.Peers {
		ids = append(ids, peer.NodeID)
	}
	return ids, nil
}

// ConnectTo implements node.RegisterArtowrk.ConnectTo()
func (service *registerArtowrk) ConnectTo(ctx context.Context, nodeID, connID string) error {
	ctx = service.contextWithLogPrefix(ctx)
	ctx = service.contextWithMDConnID(ctx)

	req := &pb.ConnectToRequest{
		NodeID: nodeID,
		ConnID: connID,
	}
	log.WithContext(ctx).WithField("req", req).Debugf("ConnectTo request")

	resp, err := service.client.ConnectTo(ctx, req)
	if err != nil {
		return errors.New("failed to request to connect to primary node")
	}
	log.WithContext(ctx).WithField("resp", resp).Debugf("ConnectTo response")

	return nil
}

// SendImage implements node.RegisterArtowrk.SendImage()
func (service *registerArtowrk) SendImage(ctx context.Context, filename string) error {
	ctx = service.contextWithLogPrefix(ctx)
	ctx = service.contextWithMDConnID(ctx)

	stream, err := service.client.SendImage(ctx)
	if err != nil {
		return errors.New("failed to open stream")
	}
	defer stream.CloseSend()

	file, err := os.Open(filename)
	if err != nil {
		return errors.Errorf("failed to open file %q: %w", filename, err)
	}
	defer file.Close()

	log.WithContext(ctx).WithField("filename", filename).Debugf("SendImage uploading")
	buffer := make([]byte, uploadImageBufferSize)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		req := &pb.SendImageRequest{
			Payload: buffer[:n],
		}
		if err := stream.Send(req); err != nil {
			return errors.New("failed to send image data")
		}
	}
	log.WithContext(ctx).Debugf("SendImage uploaded")

	return nil
}

func (service *registerArtowrk) contextWithMDConnID(ctx context.Context) context.Context {
	md := metadata.Pairs(proto.MetadataKeyConnID, service.connID)
	return metadata.NewOutgoingContext(ctx, md)
}

func (service *registerArtowrk) contextWithLogPrefix(ctx context.Context) context.Context {
	return log.ContextWithPrefix(ctx, fmt.Sprintf("%s-%s", logPrefix, service.conn.id))
}

func newRegisterArtowrk(conn *clientConn) node.RegisterArtowrk {
	return &registerArtowrk{
		conn:   conn,
		client: pb.NewRegisterArtowrkClient(conn),
	}
}
