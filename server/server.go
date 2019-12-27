package server

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
	SceneName          string
	streamID           uint32
	SceneInstanceChans map[string]chan bool
	SceneInstanceInfos map[string]map[InfoType]string
}
