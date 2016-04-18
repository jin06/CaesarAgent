package proxy

import (
	log "github.com/Sirupsen/logrus"
	sty "caesarproxy/proxy/strategy"

)

func Run(){
	server := newServer("9304",sty.NewPollStrategy())
	//
	log.Infoln(*server)
	log.Infoln(server.Server.ListenAndServe())
}