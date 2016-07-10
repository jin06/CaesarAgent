package proxy

import (
	"caesarproxy/config"
	sty "caesarproxy/proxy/strategy"
	log "github.com/Sirupsen/logrus"

	"os"
)

const (
	//distribute strategy P-PollStrategy
	//health check H-Host P-Port U-URL
	RunStyle_Normal    = 0x11 //distribute strategy-poll strategy. health check base host.
	RunStyle_PP        = 0x12 //pollstragy and port
	RunStyle_PU        = 0x13 //pollstragy and url
	RunStyle_Poll_Code = 0x1
	RunStyle_Host_Code = 0x1
	RunStyle_Port_Code = 0x2
	RunStyle_URL_Code  = 0x3
)

var RunStyle = 0x00

func Run() {
	var server *Server
	switch config.C.Proxy.Strategy {
	case "polling":
		server = newServer(config.C.Proxy.Port, sty.NewPollingStrategy())
	default:
		log.Errorln("Unknown distribute strategy:", config.C.Proxy.Strategy)
		os.Exit(0)
	}
	switch config.C.Health.Type {
	case "host":
	case "port":
	case "url":
	default:
	}
	log.Infoln("Start proxy server in port:", config.C.Proxy.Port)
	server.Strategy.Init()
	log.Infoln(server.Server.ListenAndServe())
}
