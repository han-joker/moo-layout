package moo

type server struct {
	// server kind
	Kind string
}

// 服务单例池
var serverPool = map[string]*server{}

// 单例工厂，标识于 Kind
func Server(kind string) *server {
	instance, exists := serverPool[kind]
	if !exists {
		instance = &server{
			Kind: kind,
		}
		serverPool[kind] = instance
	}
	return instance
}

func (s *server) Start() {
	switch s.Kind {
	case "websocket":
		fallthrough
	case "ws":
		s.StartWebSocket()
	}
}
