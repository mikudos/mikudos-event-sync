package server

import (
	"context"
	"fmt"

	pb "github.com/mikudos/mikudos_event_sync/proto/event_sync"
)

// FindSceneInstanceList FindSceneInstanceList
func (s *Server) FindSceneInstanceList(ctx context.Context, req *pb.FindRequest) (*pb.FindEventSyncResult, error) {
	list := []string{}
	for instanceID := range s.SceneInstanceInfos {
		if s.SceneInstanceInfos[instanceID][InstanceState] == req.GetState().String() {
			list = append(list, instanceID)
		}
	}

	res := &pb.FindEventSyncResult{InstanceIdArr: list}
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
		instanceID := resp.GetInstanceId()
		switch resp.GetEventType() {
		case pb.EventType_JOIN:
			if s.SceneInstanceChans[instanceID] == nil {
				// 创建单独局
				s.SceneInstanceChans[instanceID] = make(map[uint32]chan bool)
				// 对应局状态的初始化
				s.SceneInstanceInfos[instanceID] = map[InfoType]string{InstanceState: pb.InstanceState_WAITING.String(), InstanceLevel: "1", InstanceEntry: "charged"}
			}
			s.SceneInstanceChans[instanceID][s.streamID] = make(chan bool)
			if len(s.SceneInstanceChans[instanceID]) == 4 {
				s.SceneInstanceInfos[instanceID][InstanceState] = pb.InstanceState_READY.String()
			}
			break
		case pb.EventType_LEAVE:
			delete(s.SceneInstanceChans[instanceID], s.streamID)
			if len(s.SceneInstanceChans[instanceID]) == 0 {
				delete(s.SceneInstanceChans, instanceID)
			}
			break
		case pb.EventType_NORMAL: // handle event
			break
		}
	}
	return err
}
