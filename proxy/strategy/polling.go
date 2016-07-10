package strategy

import (
	"caesarproxy/config"
	"caesarproxy/server"
	"container/ring"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

//LoopStrategy order to distribute redirect request
//each endpoint get same request
type PollingStrategy struct {
	Servers *ring.Ring
}

func NewPollingStrategy() *PollingStrategy {
	sty := &PollingStrategy{}
	return sty
}
func (ps *PollingStrategy) Init() {
	serverNum := len(config.C.Servers)
	ps.Servers = ring.New(serverNum)
	log.Debugln("Servers number:", serverNum)
	for _, v := range config.C.Servers {
		server := server.NewServer(v.Host, v.Port, v.Weight)
		ps.Servers.Value = server
		ps.Servers = ps.Servers.Next()
	}
}

func (ps *PollingStrategy) Redirect(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Client:", r.Host, "url:", r.URL)
	if ps.Servers == nil {
		log.Warnln("Null servers.")
		w.WriteHeader(500)
		return
	}
	if server, ok := ps.Servers.Value.(server.Server); ok {
		redirect(w, server.Host, server.Port, 307)
		ps.Servers = ps.Servers.Next()
	}
	return
}
