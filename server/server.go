package server

import pb "github.com/mikudos/mikudos_event_sync/proto/event_sync"

// InfoType InfoType
type InfoType int32

const (
	// InstanceState scene state
	InstanceState InfoType = 0
	// InstanceLevel InstanceLevel
	InstanceLevel InfoType = 1
	// InstanceEntry InstanceEntry
	InstanceEntry InfoType = 2
)

// Server Server
type Server struct {
	pb.EventSyncServiceServer
	SceneName          string
	streamID           uint32
	SceneInstanceChans map[string]map[uint32]chan bool
	SceneInstanceInfos map[string]map[InfoType]string
}
