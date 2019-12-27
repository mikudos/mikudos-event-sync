package server

import (
	"context"
	"fmt"

	pb "github.com/mikudos/mikudos_event_sync/proto/event_sync"
)

// FindSceneInstanceList FindSceneInstanceList
func (s *Server) FindSceneInstanceList(ctx context.Context, req *pb.FindRequest) (*pb.FindEventSyncResult, error) {
	switch req.GetState() {
	case pb.InstanceState_WAITING:
		break
	case pb.InstanceState_READY:
		break
	case pb.InstanceState_BEGIN:
		break
	case pb.InstanceState_DELAY:
		break
	case pb.InstanceState_FINISH:
		break
	}

	res := &pb.FindEventSyncResult{}
	return res, nil
}

// EventSyncSceneInstance EventSyncSceneInstance
func (s *Server) EventSyncSceneInstance(stream pb.EventSyncService_EventSyncSceneInstanceServer) (err error) {
	s.streamID++
	go func() {
		defer fmt.Printf("GateStream break\n")
		for {
			select {
			case <-stream.Context().Done():
				return
			}
		}
	}()

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		instanceId := resp.GetInstanceId()
		switch resp.GetEventType() {
		case pb.EventType_JOIN:
			if s.SceneInstanceChans[instanceId] == nil {
				s.SceneInstanceChans[instanceId] = make(map[uint32]chan bool)
			}
			s.SceneInstanceChans[instanceId][s.streamID] = make(chan bool)
			break
		case pb.EventType_LEAVE:
			delete(s.SceneInstanceChans[instanceId], s.streamID)
			if len(s.SceneInstanceChans[instanceId]) == 0 {
				delete(s.SceneInstanceChans, instanceId)
			}
			break
		case pb.EventType_NORMAL: // handle event
			break
		}
	}
	return err
}
