package proxy

import (
	"net/http"
	//log "github.com/Sirupsen/logrus"
	"caesarproxy/config"
	sty "caesarproxy/proxy/strategy"
	"time"
)

var (
	readTimeOut  = config.C.Proxy.Timeout.Read
	writeTimeOut = config.C.Proxy.Timeout.Write
)

type Server struct {
	Server   *http.Server
	Strategy sty.Strategy
}

func newServer(port string, sty sty.Strategy) *Server {
	sm := http.NewServeMux()
	sm.HandleFunc("/", sty.Redirect)
	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        sm,
		ReadTimeout:    readTimeOut * time.Second,
		WriteTimeout:   writeTimeOut * time.Second,
		MaxHeaderBytes: 1 << 20, //1M
	}
	server := &Server{
		Server:   srv,
		Strategy: sty,
	}
	return server
}
