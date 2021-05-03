package artworkregister

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/common/random"
	"github.com/pastelnetwork/gonode/supernode/node"
	"github.com/pastelnetwork/gonode/supernode/services/artworkregister/state"
)

const (
	connectToNodeTimeout = time.Second * 1
)

// Task is the task of registering new artwork.
type Task struct {
	*Service

	ID     string
	ConnID string
	State  *state.State

	acceptMu  sync.Mutex
	connectMu sync.Mutex

	nodes  node.SuperNodes
	doneCh chan struct{}
}

// Run starts the task
func (task *Task) Run(ctx context.Context) error {
	return nil
}

// Cancel stops the task, which causes all connections associated with that task to be closed.
func (task *Task) Cancel() {
	select {
	case <-task.Done():
		return
	default:
		close(task.doneCh)
	}
}

// Done returns a channel when the task is canceled.
func (task *Task) Done() <-chan struct{} {
	return task.doneCh
}

func (task *Task) context(ctx context.Context) context.Context {
	return context.WithValue(ctx, log.PrefixKey, fmt.Sprintf("%s-%s", logPrefix, task.ID))
}

// Handshake is handshake wallet to supernode
func (task *Task) Handshake(ctx context.Context, connID string, isPrimary bool) error {
	ctx = task.context(ctx)

	if err := task.requiredStatus(state.StatusTaskStarted); err != nil {
		return err
	}
	task.ConnID = connID

	if isPrimary {
		log.WithContext(ctx).Debugf("Acts as primary node")
		task.State.Update(ctx, state.NewStatus(state.StatusHandshakePrimaryNode))
		return nil
	}

	log.WithContext(ctx).Debugf("Acts as secondary node")
	task.State.Update(ctx, state.NewStatus(state.StatusHandshakeSecondaryNode))
	return nil
}

// AcceptedNodes waits for connection supernodes, as soon as there is the required amount returns them.
func (task *Task) AcceptedNodes(ctx context.Context) (node.SuperNodes, error) {
	ctx = task.context(ctx)

	if err := task.requiredStatus(state.StatusHandshakePrimaryNode); err != nil {
		return nil, err
	}
	log.WithContext(ctx).Debugf("Waiting for supernodes to connect")

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-task.State.Updated():
		if err := task.requiredStatus(state.StatusAcceptedNodes); err != nil {
			return nil, err
		}
		return task.nodes, nil
	}
}

// HandshakeNode accepts secondary node
func (task *Task) HandshakeNode(ctx context.Context, nodeKey string) error {
	ctx = task.context(ctx)

	task.acceptMu.Lock()
	defer task.acceptMu.Unlock()

	if err := task.requiredStatus(state.StatusHandshakePrimaryNode); err != nil {
		return err
	}

	if node := task.nodes.FindByKey(nodeKey); node != nil {
		return errors.Errorf("node %q is already registered", nodeKey)
	}

	node, err := task.findNode(ctx, nodeKey)
	if err != nil {
		return err
	}
	task.nodes.Add(node)

	log.WithContext(ctx).WithField("nodeKey", nodeKey).Debugf("Accept secondary node")

	if len(task.nodes) >= task.config.NumberConnectedNodes {
		task.State.Update(ctx, state.NewStatus(state.StatusAcceptedNodes))
	}
	return nil
}

// ConnectTo connects to primary node
func (task *Task) ConnectTo(ctx context.Context, nodeKey string) error {
	ctx = task.context(ctx)

	task.connectMu.Lock()
	defer task.connectMu.Unlock()

	if err := task.requiredStatus(state.StatusHandshakeSecondaryNode); err != nil {
		return err
	}

	node, err := task.findNode(ctx, nodeKey)
	if err != nil {
		return err
	}

	connCtx, connCancel := context.WithTimeout(ctx, connectToNodeTimeout)
	defer connCancel()

	conn, err := task.nodeClient.Connect(connCtx, node.Address)
	if err != nil {
		return err
	}

	go func() {
		select {
		case <-ctx.Done():
			conn.Close()
		case <-conn.Done():
			task.Cancel()
		}
	}()

	stream, err := conn.RegisterArtowrk(ctx)
	if err != nil {
		return err
	}

	if err := stream.Handshake(ctx, task.ConnID, task.myNode.Key); err != nil {
		return err
	}

	task.State.Update(ctx, state.NewStatus(state.StatusConnectedToNode))
	return nil
}

func (task *Task) findNode(ctx context.Context, nodeKey string) (*node.SuperNode, error) {
	masterNodes, err := task.pastelClient.TopMasterNodes(ctx)
	if err != nil {
		return nil, err
	}

	for _, masterNode := range masterNodes {
		if masterNode.ExtKey != nodeKey {
			continue
		}
		node := &node.SuperNode{
			Address: masterNode.ExtAddress,
			Key:     masterNode.ExtKey,
			Fee:     masterNode.Fee,
		}
		return node, nil
	}

	return nil, errors.Errorf("node %q not found", nodeKey)
}

func (task *Task) requiredStatus(statusType state.StatusType) error {
	latest := task.State.Latest()
	if latest == nil {
		return errors.New("not found latest status")
	}

	if latest.Type != statusType {
		return errors.Errorf("wrong order, current task status %q, ", latest.Type)
	}
	return nil
}

// NewTask returns a new Task instance.
func NewTask(service *Service) *Task {
	taskID, _ := random.String(8, random.Base62Chars)

	return &Task{
		Service: service,
		ID:      taskID,
		State:   state.New(state.NewStatus(state.StatusTaskStarted)),
		doneCh:  make(chan struct{}),
	}
}
