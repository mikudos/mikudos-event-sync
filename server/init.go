package server

// Handler Server instance
var Handler Server

func init() {
	Handler = Server{SceneName: "", SceneInstanceChans: make(map[string]map[uint32]chan bool), SceneInstanceInfos: make(map[string]map[InfoType]string)}
}
