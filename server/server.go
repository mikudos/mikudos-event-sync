package server

// Server Server
type Server struct {
	SceneName          string
	streamID           int
	SceneInstanceChans map[string]chan bool
	SceneInstanceInfos map[string]map[string]string
}
