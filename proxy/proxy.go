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
	//switch config.C.Proxy.Strategy {
	//case "poll" :
	//	RunStyle = RunStyle_Poll_Code
	//default:
	//	log.Errorln("unknow distribute strategy:",config.C.Proxy.Strategy)
	//	os.Exit(0)
	//}
	//RunStyle = RunStyle << 4
	//switch config.C.Health.Type {
	//case "host":
	//	RunStyle = RunStyle + RunStyle_Host_Code
	//case "port":
	//	RunStyle = RunStyle + RunStyle_Port_Code
	//case "url":
	//	RunStyle = RunStyle + RunStyle_URL_Code
	//default:
	//	log.Errorln("unknow health check type:",config.C.Health.Type)
	//	os.Exit(0)
	//}
	var server *Server
	switch config.C.Proxy.Strategy {
	case "loop":
		server = newServer(config.C.Proxy.Port, sty.NewLoopStrategy())
	default:
		log.Errorln("unknown distribute strategy:", config.C.Proxy.Strategy)
		os.Exit(0)
	}
	switch config.C.Health.Type {
	case "host":
	case "port":
	case "url":
	default:
	}
	log.Infoln("start proxy server in port:", config.C.Proxy.Port)
	server.Strategy.Init()
	log.Infoln(server.Server.ListenAndServe())
}
