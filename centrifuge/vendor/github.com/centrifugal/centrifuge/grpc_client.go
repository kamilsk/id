package centrifuge

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/centrifugal/centrifuge/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RegisterGRPCServerClient ...
func RegisterGRPCServerClient(n *Node, server *grpc.Server, config GRPCClientServiceConfig) error {
	proto.RegisterCentrifugeServer(server, newGRPCClientService(n, config))
	return nil
}

// GRPCClientServiceConfig for GRPC client Service.
type GRPCClientServiceConfig struct{}

// GRPCClientService can work with client GRPC connections.
type grpcClientService struct {
	config GRPCClientServiceConfig
	node   *Node
}

// newGRPCClientService creates new Service.
func newGRPCClientService(n *Node, c GRPCClientServiceConfig) *grpcClientService {
	return &grpcClientService{
		config: c,
		node:   n,
	}
}

const replyBufferSize = 64

// Communicate is a bidirectional stream reading Command and
// sending Reply to client.
func (s *grpcClientService) Communicate(stream proto.Centrifuge_CommunicateServer) error {

	replies := make(chan *proto.Reply, replyBufferSize)
	transport := newGRPCTransport(stream, replies)

	c, err := newClient(stream.Context(), s.node, transport)
	if err != nil {
		s.node.logger.log(newLogEntry(LogLevelError, "error creating client", map[string]interface{}{"transport": transportGRPC}))
		return err
	}
	defer c.close(DisconnectNormal)

	s.node.logger.log(newLogEntry(LogLevelDebug, "GRPC connection established", map[string]interface{}{"client": c.ID()}))
	defer func(started time.Time) {
		s.node.logger.log(newLogEntry(LogLevelDebug, "GRPC connection completed", map[string]interface{}{"client": c.ID(), "time": time.Since(started)}))
	}(time.Now())

	go func() {
		for {
			cmd, err := stream.Recv()
			if err == io.EOF {
				c.close(DisconnectNormal)
				return
			}
			if err != nil {
				c.close(DisconnectNormal)
				return
			}
			rep, disconnect := c.handle(cmd)
			if disconnect != nil {
				s.node.logger.log(newLogEntry(LogLevelInfo, "disconnect after handling command", map[string]interface{}{"command": fmt.Sprintf("%v", cmd), "client": c.ID(), "user": c.UserID(), "reason": disconnect.Reason}))
				c.close(disconnect)
				return
			}
			if rep != nil {
				if rep.Error != nil {
					s.node.logger.log(newLogEntry(LogLevelInfo, "error in reply", map[string]interface{}{"reply": fmt.Sprintf("%v", rep), "client": c.ID(), "user": c.UserID(), "error": rep.Error.Error()}))
				}
				err = transport.Send(newPreparedReply(rep, proto.EncodingProtobuf))
				if err != nil {
					if s.node.logger.enabled(LogLevelDebug) {
						s.node.logger.log(newLogEntry(LogLevelDebug, "disconnect after sending reply", map[string]interface{}{"reply": fmt.Sprintf("%v", rep), "client": c.ID(), "user": c.UserID()}))
					}
					c.close(&Disconnect{Reason: "error sending message", Reconnect: true})
					return
				}
			}
		}
	}()

	for reply := range replies {
		if err := stream.Send(reply); err != nil {
			return err
		}
	}

	return nil
}

const (
	transportGRPC = "grpc"
)

// grpcTransport represents wrapper over stream to work with it
// from outside in abstract way.
type grpcTransport struct {
	mu      sync.Mutex
	closed  bool
	stream  proto.Centrifuge_CommunicateServer
	replies chan *proto.Reply
}

func newGRPCTransport(stream proto.Centrifuge_CommunicateServer, replies chan *proto.Reply) *grpcTransport {
	return &grpcTransport{
		stream:  stream,
		replies: replies,
	}
}

func (t *grpcTransport) Name() string {
	return transportGRPC
}

func (t *grpcTransport) Encoding() proto.Encoding {
	return proto.EncodingProtobuf
}

func (t *grpcTransport) Send(reply *preparedReply) error {
	select {
	case t.replies <- reply.Reply:
	default:
		go t.Close(DisconnectSlow)
		return io.EOF
	}
	return nil
}

func (t *grpcTransport) Close(disconnect *Disconnect) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return nil
	}
	t.closed = true
	disconnectJSON, err := json.Marshal(disconnect)
	if err != nil {
		return err
	}
	t.stream.SetTrailer(metadata.Pairs("disconnect", string(disconnectJSON)))
	close(t.replies)
	return nil
}
