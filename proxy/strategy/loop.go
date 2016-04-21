package strategy

import (
	"caesarproxy/config"
	"caesarproxy/endpoint"
	"container/ring"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

//LoopStrategy order to distribute redirect request
//each endpoint get same request
type LoopStrategy struct {
	Endpoints *ring.Ring
}

func NewLoopStrategy() *LoopStrategy {
	sty := &LoopStrategy{}
	return sty
}
func (ps *LoopStrategy) Init() {
	numOfEndpoints := len(config.C.Endpoints)
	ps.Endpoints = ring.New(numOfEndpoints)
	log.Debugln("endpoints number:", numOfEndpoints)
	for _, v := range config.C.Endpoints {
		endpoint := endpoint.NewEndPoint(v.Host, v.Port, v.Weight)
		ps.Endpoints.Value = endpoint
		ps.Endpoints = ps.Endpoints.Next()
	}
}

func (ps *LoopStrategy) Redirect(w http.ResponseWriter, r *http.Request) {
	log.Infoln("client:", r.Host, "url:", r.URL)
	if ps.Endpoints == nil {
		log.Warnln("null endpoints.")
		w.WriteHeader(500)
		return
	}
	if endpoint, ok := ps.Endpoints.Value.(endpoint.EndPoint); ok {
		redirect(w, endpoint.Host, endpoint.Port, 307)
		ps.Endpoints = ps.Endpoints.Next()
	}
	return
}
